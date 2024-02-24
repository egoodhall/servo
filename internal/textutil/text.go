package textutil

import (
	"regexp"
	"strings"
)

func TrimDedent(in string) string {
	return Dedent(TrimBlankLines(in))
}

var blankLine = regexp.MustCompile("^[\t ]*$")

func TrimBlankLines(in string) string {
	lines := strings.Split(in, "\n")
	var i, j int
	for i = 0; blankLine.MatchString(lines[i]); i++ {
	}
	for j = len(lines) - 1; blankLine.MatchString(lines[j]); j-- {
	}
	return strings.Join(lines[i:j+1], "\n")
}

var leadingWhitespace = regexp.MustCompile("(?m)^( +|\t+)")

func Dedent(in string) string {
	var char byte
	var amount int
	for _, match := range leadingWhitespace.FindAllString(in, -1) {
		if char == 0 {
			char = match[0]
		} else if char != match[0] {
			return in
		}

		if amount == 0 || amount > len(match) {
			amount = len(match)
		}
	}

	return leadingWhitespace.ReplaceAllStringFunc(in, func(s string) string {
		return s[:len(s)-amount]
	})
}
