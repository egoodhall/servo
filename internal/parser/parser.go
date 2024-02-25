package parser

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/egoodhall/servo/internal/parser/parsegen"
	"github.com/egoodhall/servo/internal/validate"
	"github.com/egoodhall/servo/pkg/ast"
)

func Files[T io.Reader](in ...T) ([]ast.File, error) {
	files := make([]ast.File, len(in))
	var err error
	for i, rd := range in {
		files[i], err = File(rd)
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

func File(in io.Reader) (ast.File, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return ast.File{}, err
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
		return ast.File{}, err
	}

	if tc.errors {
		return ast.File{}, errors.New("syntax error")
	}

	f, err := gather(tc.tokens)
	if err != nil {
		return f, err
	}

	return f, validate.File(f)
}

type token struct {
	typ   parsegen.NodeType
	value string
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
	tc.tokens = append(tc.tokens, token{
		typ:   typ,
		value: tc.input[frm:to],
	})
}

func gather(tokens []token) (ast.File, error) {
	file := ast.File{
		Options:  make([]ast.Option, 0),
		Enums:    make([]ast.Enum, 0),
		Messages: make([]ast.Message, 0),
		Services: make([]ast.Service, 0),
	}

	tokenGroups := partition(tokens, func(typ parsegen.NodeType) bool {
		return typ == parsegen.MessageName ||
			typ == parsegen.ServiceName ||
			typ == parsegen.OptionName ||
			typ == parsegen.EnumName
	})

	for _, tkns := range tokenGroups {
		switch tkns[0].typ {
		case parsegen.OptionName:
			opt, err := gatherOption(tkns[0], tkns[1:])
			if err != nil {
				return ast.File{}, err
			}
			file.Options = append(file.Options, opt)
		case parsegen.ServiceName:
			svc, err := gatherService(tkns[0], tkns[1:])
			if err != nil {
				return ast.File{}, err
			}
			file.Services = append(file.Services, svc)
		case parsegen.MessageName:
			msg, err := gatherMessage(tkns[0], tkns[1:])
			if err != nil {
				return ast.File{}, err
			}
			file.Messages = append(file.Messages, msg)
		case parsegen.EnumName:
			file.Enums = append(file.Enums, ast.Enum{
				Name:   tkns[0].value,
				Values: values(tkns[1:]),
			})
		default:
			return ast.File{}, fmt.Errorf("unexpected %s token: '%s'", tkns[0].typ, tkns[0].value)
		}
	}
	return file, nil
}

func gatherOption(name token, tokens []token) (ast.Option, error) {
	if len(tokens) != 1 {
		return ast.Option{}, fmt.Errorf("token mismatch: need 1 value")
	}
	return ast.Option{
		Name:  name.value,
		Value: strings.Trim(tokens[0].value, `"`),
	}, nil
}

func gatherService(name token, tokens []token) (ast.Service, error) {
	svc := ast.Service{
		Name: name.value,
	}

	for _, method := range partition(tokens, func(typ parsegen.NodeType) bool {
		return typ == parsegen.PubName || typ == parsegen.RpcName
	}) {
		name := method[0]
		switch name.typ {
		case parsegen.PubName:
			pub, err := gatherPub(name, method[1:])
			if err != nil {
				return ast.Service{}, fmt.Errorf("method %s: %w", name.value, err)
			}
			if svc.Pubs == nil {
				svc.Pubs = []ast.Pub{pub}
			} else {
				svc.Pubs = append(svc.Pubs, pub)
			}
		case parsegen.RpcName:
			rpc, err := gatherRpc(name, method[1:])
			if err != nil {
				return ast.Service{}, fmt.Errorf("method %s: %w", name.value, err)
			}
			if svc.Rpcs == nil {
				svc.Rpcs = []ast.Rpc{rpc}
			} else {
				svc.Rpcs = append(svc.Rpcs, rpc)
			}
		default:
			return ast.Service{}, fmt.Errorf("unexpected %s token: '%s'", name.typ, name.value)
		}
	}

	return svc, nil
}

func gatherPub(name token, tokens []token) (ast.Pub, error) {
	if len(tokens) != 1 {
		return ast.Pub{}, fmt.Errorf("token mismatch: need 1 message")
	}

	msg := tokens[0]
	if msg.typ != parsegen.PubMessage {
		return ast.Pub{}, fmt.Errorf("unexpected token for pub message: %s", msg.typ)
	}
	return ast.Pub{
		Name:    name.value,
		Message: msg.value,
	}, nil
}

func gatherRpc(name token, tokens []token) (ast.Rpc, error) {
	if len(tokens) != 2 {
		return ast.Rpc{}, fmt.Errorf("token mismatch: need 1 request and 1 response")
	}

	req := tokens[0]
	if req.typ != parsegen.RpcRequest {
		return ast.Rpc{}, fmt.Errorf("unexpected token for rpc request: %s", req.typ)
	}

	res := tokens[1]
	if res.typ != parsegen.RpcResponse {
		return ast.Rpc{}, fmt.Errorf("unexpected token for rpc response: %s", res.typ)
	}

	return ast.Rpc{
		Name:     name.value,
		Request:  req.value,
		Response: res.value,
	}, nil
}

func gatherMessage(name token, tokens []token) (ast.Message, error) {
	svc := ast.Message{
		Name:   name.value,
		Fields: make([]ast.Field, 0),
	}

	for _, field := range partition(tokens, func(typ parsegen.NodeType) bool {
		return typ == parsegen.FieldName
	}) {
		switch field[0].typ {
		case parsegen.FieldName:
			fld, err := gatherField(field[0], field[1:])
			if err != nil {
				return ast.Message{}, fmt.Errorf("field %s: %w", field[0].value, err)
			}
			svc.Fields = append(svc.Fields, fld)
		default:
			return ast.Message{}, fmt.Errorf("unexpected %s token: '%s'", field[0].typ, field[0].value)
		}
	}

	return svc, nil
}

func gatherField(name token, tokens []token) (ast.Field, error) {
	typ, err := gatherFieldType(tokens)
	if err != nil {
		return ast.Field{}, err
	}

	return ast.Field{
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
			KeyType:   ast.ScalarType{Name: tokens[0].value},
			ValueType: ast.ScalarType{Name: tokens[1].value},
		}, nil
	case parsegen.ListElement:
		return ast.ListType{
			ElementType: ast.ScalarType{Name: tokens[0].value},
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

func partition(in []token, test func(parsegen.NodeType) bool) [][]token {
	partitions := make([][]token, 0)
	start := 0
	for i, t := range in {
		if test(t.typ) && i != start {
			partitions = append(partitions, in[start:i])
			start = i
		}
	}
	if start < len(in)-1 {
		partitions = append(partitions, in[start:])
	}
	return partitions
}
