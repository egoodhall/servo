package parser

import (
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/egoodhall/servo/internal/parser/parsegen"
	"github.com/egoodhall/servo/pkg/ast"
)

func Files(in ...*os.File) ([]*ast.File, error) {
	files := make([]*ast.File, len(in))
	var err error
	for i, rd := range in {
		files[i], err = File(rd.Name(), rd)
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

func File(name string, in io.Reader) (*ast.File, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}

	lexer := new(parsegen.Lexer)
	lexer.Init(string(data))

	tc := &tokenCollector{
		input:  string(data),
		tokens: make([]token, 0),
	}

	parser := new(parsegen.Parser)
	parser.Init(tc.HandleError, tc.Collect)

	if err := parser.Parse(lexer); err != nil {
		return nil, err
	}

	if tc.errors {
		return nil, errors.New("syntax error")
	}

	if err := validateTokens(name, tc.tokens); err != nil {
		return nil, err
	}

	f, err := gather(tc.tokens)
	if err != nil {
		return f, err
	}

	f.Name = name
	return f, validateAst(f)
}

type token struct {
	typ   parsegen.NodeType
	value string
	line  int
	col   int
}

func (t token) String() string {
	return fmt.Sprintf("%s{%s}", t.typ, t.value)
}

type tokenCollector struct {
	input  string
	tokens []token
	errors bool
}

func (tc *tokenCollector) HandleError(err parsegen.SyntaxError) bool {
	fmt.Printf("Syntax error at line %d on '%s'\n", err.Line, tc.input[err.Offset:err.Endoffset])
	tc.errors = true
	return true
}

func (tc *tokenCollector) Collect(typ parsegen.NodeType, frm, to int) {
	var line, col int
	lineStart := strings.LastIndex(tc.input[0:frm], "\n")

	if lineStart < 0 {
		line = 1
		col = frm
	} else {
		line = strings.Count(tc.input[0:frm], "\n") + 1
		col = frm - lineStart
	}

	tc.tokens = append(tc.tokens, token{
		typ:   typ,
		value: tc.input[frm:to],
		line:  line,
		col:   col,
	})
}

func gather(tokens []token) (*ast.File, error) {
	file := new(ast.File)

	tokenGroups := partition(tokens,
		parsegen.MessageName,
		parsegen.UnionName,
		parsegen.ServiceName,
		parsegen.OptionName,
		parsegen.EnumName,
		parsegen.AliasName,
	)

	for _, tkns := range tokenGroups {
		switch tkns[0].typ {
		case parsegen.OptionName:
			opt, err := gatherOption(tkns[0], tkns[1:])
			if err != nil {
				return nil, err
			}
			file.Options = append(file.Options, opt)
		case parsegen.ServiceName:
			svc, err := gatherService(tkns[0], tkns[1:])
			if err != nil {
				return nil, err
			}
			file.Services = append(file.Services, svc)
		case parsegen.MessageName:
			msg, err := gatherMessage(tkns[0], tkns[1:])
			if err != nil {
				return nil, err
			}
			file.Messages = append(file.Messages, msg)
		case parsegen.UnionName:
			uni, err := gatherUnion(tkns[0], tkns[1:])
			if err != nil {
				return nil, err
			}
			file.Unions = append(file.Unions, uni)
		case parsegen.EnumName:
			file.Enums = append(file.Enums, &ast.Enum{
				Name:   tkns[0].value,
				Values: values(tkns[1:]),
			})
		case parsegen.AliasName:
			file.Aliases = append(file.Aliases, &ast.Alias{
				Name: tkns[0].value,
				Type: tkns[1].value,
			})
		default:
			return nil, fmt.Errorf("unexpected %s token: '%s'", tkns[0].typ, tkns[0].value)
		}
	}
	return file, nil
}

func gatherOption(name token, tokens []token) (*ast.Option[any], error) {
	if len(tokens) != 1 {
		return nil, fmt.Errorf("token mismatch: need 1 value")
	}

	token := tokens[0]
	var value any
	switch token.typ {
	case parsegen.OptionString:
		value = strings.Trim(token.value, `"`)
	case parsegen.OptionBool:
		value = token.value == "true"
	case parsegen.OptionInt:
		v, err := strconv.Atoi(token.value)
		if err != nil {
			return nil, err
		}
		value = v
	case parsegen.OptionFloat:
		if v, err := strconv.ParseFloat(token.value, 32); errors.Is(err, strconv.ErrRange) {
			if v, err := strconv.ParseFloat(token.value, 64); err != nil {
				return nil, err
			} else {
				value = v
			}
		} else if err != nil {
			return nil, err
		} else {
			value = v
		}
	default:
		return nil, fmt.Errorf("unexpected option value token: %s", token.typ)
	}

	return &ast.Option[any]{
		Name:  name.value,
		Value: value,
	}, nil
}

func gatherService(name token, tokens []token) (*ast.Service, error) {
	svc := ast.Service{
		Name: name.value,
	}

	for _, method := range partition(tokens, parsegen.RpcName) {
		name := method[0]
		switch name.typ {
		case parsegen.RpcName:
			rpc, err := gatherRpc(name, method[1:])
			if err != nil {
				return nil, fmt.Errorf("method %s: %w", name.value, err)
			}
			if svc.Rpcs == nil {
				svc.Rpcs = []*ast.Rpc{rpc}
			} else {
				svc.Rpcs = append(svc.Rpcs, rpc)
			}
		default:
			return nil, fmt.Errorf("unexpected %s token: '%s'", name.typ, name.value)
		}
	}

	return &svc, nil
}

func gatherRpc(name token, tokens []token) (*ast.Rpc, error) {
	if len(tokens) < 1 || len(tokens) > 2 {
		return nil, fmt.Errorf("token mismatch: need 1 request and optionally 1 response")
	}
	var req, res token

	req = tokens[0]
	if req.typ != parsegen.RpcRequest {
		return nil, fmt.Errorf("unexpected token for rpc request: %s", req.typ)
	}

	if len(tokens) != 2 {
		return &ast.Rpc{
			Name:    name.value,
			Request: req.value,
		}, nil
	}

	res = tokens[1]
	if res.typ != parsegen.RpcResponse {
		return nil, fmt.Errorf("unexpected token for rpc response: %s", res.typ)
	}

	return &ast.Rpc{
		Name:     name.value,
		Request:  req.value,
		Response: res.value,
	}, nil
}

func gatherMessage(name token, tokens []token) (*ast.Message, error) {
	msg := ast.Message{
		Name:   name.value,
		Fields: make([]*ast.Field, 0),
	}

	for _, field := range partition(tokens, parsegen.FieldName) {
		switch field[0].typ {
		case parsegen.FieldName:
			fld, err := gatherField(field[0], field[1:])
			if err != nil {
				return nil, fmt.Errorf("field %s: %w", field[0].value, err)
			}
			msg.Fields = append(msg.Fields, fld)
		default:
			return nil, fmt.Errorf("unexpected %s token: '%s'", field[0].typ, field[0].value)
		}
	}

	return &msg, nil
}

func gatherUnion(name token, tokens []token) (*ast.Union, error) {
	uni := ast.Union{
		Name:    name.value,
		Members: make([]*ast.Member, 0),
	}

	for _, field := range partition(tokens, parsegen.FieldName) {
		switch field[0].typ {
		case parsegen.FieldName:
			mem, err := gatherMember(field[0], field[1:])
			if err != nil {
				return nil, fmt.Errorf("field %s: %w", field[0].value, err)
			}
			uni.Members = append(uni.Members, mem)
		default:
			return nil, fmt.Errorf("unexpected %s token: '%s'", field[0].typ, field[0].value)
		}
	}

	return &uni, nil
}

func gatherMember(name token, tokens []token) (*ast.Member, error) {
	switch tokens[0].typ {
	case parsegen.ScalarType:
		return &ast.Member{
			Name: name.value,
			Type: ast.ScalarType{
				Name: tokens[0].value,
			},
		}, nil
	default:
		return nil, fmt.Errorf("unexpected field type: %s", tokens[0].typ.String())
	}
}

func gatherField(name token, tokens []token) (*ast.Field, error) {
	typ, err := gatherFieldType(tokens)
	if err != nil {
		return nil, err
	}

	return &ast.Field{
		Name:     name.value,
		Type:     typ,
		Optional: isOptional(tokens),
	}, nil
}

func gatherFieldType(tokens []token) (ast.Type, error) {
	switch tokens[0].typ {
	case parsegen.ScalarType:
		return ast.ScalarType{
			Name: tokens[0].value,
		}, nil
	case parsegen.MapKeyType:
		return ast.MapType{
			KeyType:   &ast.ScalarType{Name: tokens[0].value},
			ValueType: &ast.ScalarType{Name: tokens[1].value},
		}, nil
	case parsegen.ListElement:
		return ast.ListType{
			ElementType: &ast.ScalarType{Name: tokens[0].value},
		}, nil
	default:
		return nil, fmt.Errorf("unexpected field type: %s", tokens[0].typ.String())
	}
}

func isOptional(tokens []token) bool {
	tkn := tokens[len(tokens)-1]
	return tkn.typ == parsegen.FieldMod && strings.Contains(tkn.value, "?")
}

func values(in []token) []string {
	vals := make([]string, len(in))
	for i, tok := range in {
		vals[i] = tok.value
	}
	return vals
}

func partition(in []token, delimeter parsegen.NodeType, delimeters ...parsegen.NodeType) [][]token {
	partitions := make([][]token, 0)
	start := 0
	for i, t := range in {
		if i != start && (delimeter == t.typ || slices.Contains(delimeters, t.typ)) {
			partitions = append(partitions, in[start:i])
			start = i
		}
	}
	if start < len(in)-1 {
		partitions = append(partitions, in[start:])
	}
	return partitions
}
