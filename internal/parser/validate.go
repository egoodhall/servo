package parser

import (
	"errors"
	"fmt"

	"github.com/egoodhall/servo/internal/parser/parsegen"
	"github.com/egoodhall/servo/pkg/ast"
)

func validateAst(file *ast.File) error {
	errs := make([]error, 0)
	for _, msg := range file.Messages {
		errs = append(errs, validateFields("message field", msg.Fields, func(f *ast.Field) string {
			return f.Name
		})...)
	}

	for _, uni := range file.Unions {
		errs = append(errs, validateFields("union member", uni.Members, func(f *ast.Member) string {
			return f.Name
		})...)
	}

	for _, msg := range file.Enums {
		errs = append(errs, validateFields("enum constant", msg.Values, func(f string) string {
			return f
		})...)
	}

	errs = append(errs, validateFields("option", file.Options, func(o *ast.Option[any]) string {
		return o.Name
	})...)

	return errors.Join(errs...)
}

func validateFields[T any](typ string, items []T, extractor func(T) string) []error {
	m := make(map[string]struct{})
	errs := make([]error, 0)

	for _, t := range items {
		s := extractor(t)
		if _, ok := m[s]; ok {
			errs = append(errs, fmt.Errorf("duplicate %s name: %s", typ, s))
		}
		m[s] = struct{}{}
	}

	return errs
}

func validateTokens(filename string, tokens []token) error {
	errs := make([]error, 0)
	defs := collectDefinitionTokens(tokens, errs)

	for _, tkn := range tokens {
		switch tkn.typ {
		case parsegen.ScalarType, parsegen.ListElement,
			parsegen.MapKeyType, parsegen.MapValueType,
			parsegen.RpcRequest, parsegen.RpcResponse,
			parsegen.AliasType:

			if _, ok := defs[tkn.value]; !ok && !ast.IsPrimitive(tkn.value) {
				errs = append(errs, fmt.Errorf("reference to undefined type '%s' %s:%d:%d", tkn.value, filename, tkn.line, tkn.col))
			}

		}
	}

	return errors.Join(errs...)
}

func collectDefinitionTokens(tokens []token, errs []error) (defs map[string]struct{}) {
	defs = make(map[string]struct{})

	for _, tkn := range tokens {
		switch tkn.typ {
		case parsegen.EnumName, parsegen.MessageName,
			parsegen.UnionName, parsegen.ServiceName,
			parsegen.AliasName:

			if _, ok := defs[tkn.value]; ok {
				errs = append(errs, fmt.Errorf("multiple type definitions for name: '%s'", tkn.value))
			}
			defs[tkn.value] = struct{}{}
		}
	}
	return defs
}
