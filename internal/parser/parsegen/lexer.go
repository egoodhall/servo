// generated by Textmapper; DO NOT EDIT

package parsegen

import (
	"strings"
	"unicode/utf8"

	"github.com/egoodhall/servo/internal/parser/parsegen/token"
)

// Lexer states.
const (
	StateInitial             = 0
	StateInMessageDefinition = 1
	StateInServiceDefinition = 2
	StateInEnumDefinition    = 3
	StateInOption            = 4
)

// Lexer uses a generated DFA to scan through a utf-8 encoded input string. If
// the string starts with a BOM character, it gets skipped.
type Lexer struct {
	source string

	ch          rune // current character, -1 means EOI
	offset      int  // character offset
	tokenOffset int  // last token byte offset
	line        int  // current line number (1-based)
	tokenLine   int  // last token line
	scanOffset  int  // scanning offset
	value       interface{}

	State int // lexer state, modifiable
}

var bomSeq = "\xef\xbb\xbf"

// Init prepares the lexer l to tokenize source by performing the full reset
// of the internal state.
func (l *Lexer) Init(source string) {
	l.source = source

	l.ch = 0
	l.offset = 0
	l.tokenOffset = 0
	l.line = 1
	l.tokenLine = 1
	l.State = 0

	if strings.HasPrefix(source, bomSeq) {
		l.offset += len(bomSeq)
	}

	l.rewind(l.offset)
}

// Next finds and returns the next token in l.source. The source end is
// indicated by Token.EOI.
//
// The token text can be retrieved later by calling the Text() method.
func (l *Lexer) Next() token.Type {
restart:
	l.tokenLine = l.line
	l.tokenOffset = l.offset

	state := tmStateMap[l.State]
	hash := uint32(0)
	backupRule := -1
	var backupOffset int
	backupHash := hash
	for state >= 0 {
		var ch int
		if uint(l.ch) < tmRuneClassLen {
			ch = int(tmRuneClass[l.ch])
		} else if l.ch < 0 {
			state = int(tmLexerAction[state*tmNumClasses])
			if state > tmFirstRule && state < 0 {
				state = (-1 - state) * 2
				backupRule = tmBacktracking[state]
				backupOffset = l.offset
				backupHash = hash
				state = tmBacktracking[state+1]
			}
			continue
		} else {
			ch = 1
		}
		state = int(tmLexerAction[state*tmNumClasses+ch])
		if state > tmFirstRule {
			if state < 0 {
				state = (-1 - state) * 2
				backupRule = tmBacktracking[state]
				backupOffset = l.offset
				backupHash = hash
				state = tmBacktracking[state+1]
			}
			hash = hash*uint32(31) + uint32(l.ch)
			if l.ch == '\n' {
				l.line++
			}

			// Scan the next character.
			// Note: the following code is inlined to avoid performance implications.
			l.offset = l.scanOffset
			if l.offset < len(l.source) {
				r, w := rune(l.source[l.offset]), 1
				if r >= 0x80 {
					// not ASCII
					r, w = utf8.DecodeRuneInString(l.source[l.offset:])
				}
				l.scanOffset += w
				l.ch = r
			} else {
				l.ch = -1 // EOI
			}
		}
	}

	rule := tmFirstRule - state
recovered:
	switch rule {
	case 17:
		hh := hash & 7
		switch hh {
		case 5:
			if hash == 0x1b2fd && "pub" == l.source[l.tokenOffset:l.offset] {
				rule = 19
				break
			}
			if hash == 0x1b9e5 && "rpc" == l.source[l.tokenOffset:l.offset] {
				rule = 18
				break
			}
		}
	}

	tok := tmToken[rule]
	var space bool
	switch rule {
	case 0:
		if backupRule >= 0 {
			rule = backupRule
			hash = backupHash
			l.rewind(backupOffset)
		} else if l.offset == l.tokenOffset {
			l.rewind(l.scanOffset)
		}
		if rule != 0 {
			goto recovered
		}
	case 2: // 'enum': /enum/
		{
			l.State = StateInEnumDefinition
		}
	case 3: // 'message': /message/
		{
			l.State = StateInMessageDefinition
		}
	case 4: // 'service': /service/
		{
			l.State = StateInServiceDefinition
		}
	case 5: // 'option': /option/
		{
			l.State = StateInOption
		}
	case 6: // WS: /[ \t\n\r]+/
		space = true
	case 7: // EolComment: /\/\/[^\n]+\n/
		space = true
	case 8: // BlockComment: /\/\*([^*]|\*+[^*\/])*\**\*\//
		space = true
	case 25: // '}': /\}/
		{
			l.State = StateInitial
		}
	case 28: // ';': /;/
		{
			l.State = StateInitial
		}
	}
	if space {
		goto restart
	}
	return tok
}

// Pos returns the start and end positions of the last token returned by Next().
func (l *Lexer) Pos() (start, end int) {
	start = l.tokenOffset
	end = l.offset
	return
}

// Line returns the line number of the last token returned by Next() (1-based).
func (l *Lexer) Line() int {
	return l.tokenLine
}

// Text returns the substring of the input corresponding to the last token.
func (l *Lexer) Text() string {
	return l.source[l.tokenOffset:l.offset]
}

// Value returns the value associated with the last returned token.
func (l *Lexer) Value() interface{} {
	return l.value
}

// rewind can be used in lexer actions to accept a portion of a scanned token, or to include
// more text into it.
func (l *Lexer) rewind(offset int) {
	if offset < l.offset {
		l.line -= strings.Count(l.source[offset:l.offset], "\n")
	} else {
		if offset > len(l.source) {
			offset = len(l.source)
		}
		l.line += strings.Count(l.source[l.offset:offset], "\n")
	}

	// Scan the next character.
	l.scanOffset = offset
	l.offset = offset
	if l.offset < len(l.source) {
		r, w := rune(l.source[l.offset]), 1
		if r >= 0x80 {
			// not ASCII
			r, w = utf8.DecodeRuneInString(l.source[l.offset:])
		}
		l.scanOffset += w
		l.ch = r
	} else {
		l.ch = -1 // EOI
	}
}
