package plugin

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/egoodhall/servo/ast"
	"github.com/egoodhall/servo/option"
)

var (
	intPattern   = regexp.MustCompile("-?[0-9]+")
	floatPattern = regexp.MustCompile(`-?[0-9]+\.[0-9]+`)
	boolPattern  = regexp.MustCompile("true|false")
)

func ParseBoolOption(value string) bool {
	return value == "true"
}

func ParseIntOption(value string) (int, error) {
	i, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func ParseFloatOption(value string) (float64, error) {
	if v, err := strconv.ParseFloat(value, 64); err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

func ParseStringOption(value string) any {
	if boolPattern.MatchString(value) {
		return ParseBoolOption(value)
	}

	if intPattern.MatchString(value) {
		if v, err := ParseIntOption(value); err != nil {
			return v
		} else {
			return value
		}
	}

	if floatPattern.MatchString(value) {
		if v, err := ParseFloatOption(value); err != nil {
			return v
		} else {
			return value
		}
	}
	return value
}

func ReadOptionsDescriptor[T any]() ([]*option.Option, error) {
	otype := reflect.TypeOf(*new(T))

	var err error
	options := make([]*option.Option, otype.NumField())
	for i := 0; i < otype.NumField(); i++ {
		ofield := otype.Field(i)

		options[i], err = option.From(otype.Field(i))
		if err != nil {
			return nil, fmt.Errorf("error parsing option %s: %w", ofield.Name, err)
		}
	}

	return options, nil
}

func parseOptions[O any](prefix string, definitions []*option.Option, options []*ast.Option, globals []*ast.Option, into *O) error {
	prefix = strings.TrimSuffix(prefix, ".")
	optionsMap := make(map[string]any)
	// Start with default values
	intoMap(optionsMap, definitions,
		func(def *option.Option) string { return def.Name },
		func(def *option.Option) (any, bool) { return def.Value, def.Value != nil },
	)
	// Then add file-specific overrides
	intoMap(optionsMap, options,
		func(opt *ast.Option) string { return strings.TrimPrefix(opt.Name, prefix+".") },
		func(opt *ast.Option) (any, bool) { return opt.Value, opt.Value != nil },
	)
	// Then add CLI-defined values
	intoMap(optionsMap, globals,
		func(opt *ast.Option) string { return strings.TrimPrefix(opt.Name, prefix+".") },
		func(opt *ast.Option) (any, bool) { return opt.Value, opt.Value != nil },
	)

	j, err := json.Marshal(optionsMap)
	if err != nil {
		return err
	}

	return json.Unmarshal(j, into)
}

func intoMap[T, V any, K comparable](m map[K]V, items []T, keyer func(T) K, valuer func(T) (V, bool)) {
	for _, t := range items {
		if v, ok := valuer(t); ok {
			m[keyer(t)] = v
		}
	}
}
