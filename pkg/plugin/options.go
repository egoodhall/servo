package plugin

import (
	"encoding/json"
	"fmt"
	"github.com/egoodhall/servo/internal/plugin"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/egoodhall/servo/pkg/ipc"
	"reflect"
	"strings"
)

func ReadOptionsDescriptor[T any]() ([]*ipc.Option[any], error) {
	otype := reflect.TypeOf(*new(T))
	options := make([]*ipc.Option[any], otype.NumField())
	for i := 0; i < otype.NumField(); i++ {
		ofield := otype.Field(i)

		options[i] = &ipc.Option[any]{
			Name:        ofield.Tag.Get("json"),
			Description: ofield.Tag.Get("desc"),
		}

		odef, ok := ofield.Tag.Lookup("default")
		kind := ofield.Type.Kind()
		if ok {
			switch kind {
			case reflect.Bool:
				options[i].Default = plugin.ParseBoolOption(odef)
			case reflect.Int:
				if val, err := plugin.ParseIntOption(odef); err != nil {
					return nil, fmt.Errorf("error parsing '%s' as %s", odef, kind)
				} else {
					options[i].Default = val
				}
			case reflect.Float64:
				if val, err := plugin.ParseFloatOption(odef); err != nil {
					return nil, fmt.Errorf("error parsing '%s' as %s", odef, kind)
				} else {
					options[i].Default = val
				}
			case reflect.String:
				options[i].Default = odef
			default:
				return nil, fmt.Errorf("unsupported option type: %s", kind)
			}
		}
	}
	return options, nil
}

func parseOptions[O any](prefix string, definitions []*ipc.Option[any], options []*ast.Option[any], globals []*ast.Option[any], into *O) error {
	prefix = strings.TrimSuffix(prefix, ".")
	optionsMap := make(map[string]any)
	// Start with default values
	intoMap(optionsMap, definitions,
		func(def *ipc.Option[any]) string { return def.Name },
		func(def *ipc.Option[any]) (any, bool) { return def.Default, def.Default != nil },
	)
	// Then add globally defined values
	intoMap(optionsMap, globals,
		func(opt *ast.Option[any]) string { return strings.TrimPrefix(opt.Name, prefix+".") },
		func(opt *ast.Option[any]) (any, bool) { return opt.Value, opt.Value != nil },
	)
	// Then add file-specific overrides
	intoMap(optionsMap, options,
		func(opt *ast.Option[any]) string { return strings.TrimPrefix(opt.Name, prefix+".") },
		func(opt *ast.Option[any]) (any, bool) { return opt.Value, opt.Value != nil },
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
