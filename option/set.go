package option

import (
	"fmt"
	"slices"
	"strings"

	"github.com/egoodhall/servo/ast"
)

type Set map[string]*Option

func (s Set) Parse(name, value string) (any, error) {
	opt, ok := s[name]
	if !ok {
		return nil, fmt.Errorf("unknown option: %s", name)
	}
	return Parse(value, opt.Type)
}

func (s Set) SortedByName() []*Option {
	opts := make([]*Option, 0, len(s))
	for _, opt := range s {
		opts = append(opts, opt)
	}

	slices.SortFunc(opts, func(a, b *Option) int {
		return strings.Compare(a.Name, b.Name)
	})
	return opts
}

func (s Set) ToAst(values map[string]string) ([]*ast.Option, error) {
	options := make([]*ast.Option, 0, len(s))
	for _, opt := range s {
		option := &ast.Option{
			Name:  opt.Name,
			Value: opt.Value,
		}

		if value, ok := values[opt.Name]; ok {
			var err error
			if option.Value, err = s.Parse(opt.Name, value); err != nil {
				return nil, err
			}
		}
		options = append(options, option)
	}

	return options, nil
}

func NewSet(opts []*Option) Set {
	s := make(Set)
	for _, opt := range opts {
		s[opt.Name] = opt
	}
	return s
}
