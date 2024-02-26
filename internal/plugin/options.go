package plugin

import (
	"github.com/egoodhall/servo/pkg/ast"
	"regexp"
	"strconv"
)

func ToOptions(m map[string]string) []*ast.Option[any] {
	opts := make([]*ast.Option[any], len(m))
	i := 0
	for k, v := range m {
		opts[i] = &ast.Option[any]{
			Name:  k,
			Value: ParseStringOption(v),
		}
		i++
	}
	return opts
}

var (
	intPattern   = regexp.MustCompile("-?[0-9]+")
	floatPattern = regexp.MustCompile("-?[0-9]+\\.[0-9]+")
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
