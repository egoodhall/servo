package validate

import (
	"errors"
	"fmt"

	"github.com/egoodhall/servo/pkg/ast"
)

func File(file ast.File) error {
	errs := make([]error, 0)
	if err := validateUniqueDeclaredTypeNames(file); err != nil {
		errs = append(errs, err)
	}

	for _, enum := range file.Enums {
		if err := validateUnique(enum.Values); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", enum.Name, err))
		}
	}

	for _, message := range file.Messages {
		if err := validateUnique(extract(message.Fields, func(f ast.Field) string { return f.Name })); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", message.Name, err))
		}
	}

	for _, service := range file.Services {
		if err := validateUnique(
			extract(service.Rpcs, func(r ast.Rpc) string { return r.Name }),
			extract(service.Pubs, func(p ast.Pub) string { return p.Name }),
		); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", service.Name, err))
		}
	}

	if err := validateTypeRefs(file); err != nil {
		errs = append(errs, err)
	}

	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}

func validateUniqueDeclaredTypeNames(file ast.File) error {
	return validateUnique(
		extract(file.Enums, func(e ast.Enum) string { return e.Name }),
		extract(file.Messages, func(m ast.Message) string { return m.Name }),
		extract(file.Services, func(s ast.Service) string { return s.Name }),
	)
}

func validateTypeRefs(file ast.File) error {
	declaredNames := getDeclaredTypeNames(file)
	referencedNames := getReferencedTypeNames(file)
	errs := make(set[error])
	for name := range referencedNames {
		if _, ok := declaredNames[name]; !ok {
			errs[fmt.Errorf("reference to undeclared type %s", name)] = struct{}{}
		}
	}
	return errors.Join(errs.toSlice()...)
}

func getReferencedTypeNames(file ast.File) set[string] {
	names := make(set[string])
	for _, message := range file.Messages {
		for _, field := range message.Fields {
			switch t := (field.Type).(type) {
			case ast.ScalarType:
				names[t.Name] = struct{}{}
			case ast.ListType:
				names[t.ElementType.Name] = struct{}{}
			case ast.MapType:
				names[t.KeyType.Name] = struct{}{}
				names[t.ValueType.Name] = struct{}{}
			}
		}
	}
	for _, service := range file.Services {
		for _, rpc := range service.Rpcs {
			names[rpc.Request] = struct{}{}
			names[rpc.Response] = struct{}{}
		}
		for _, pub := range service.Pubs {
			names[pub.Message] = struct{}{}
		}
	}
	return names
}

func getDeclaredTypeNames(file ast.File) set[string] {
	names := buildSet(
		extract(file.Enums, func(e ast.Enum) string { return e.Name }),
		extract(file.Messages, func(m ast.Message) string { return m.Name }),
		extract(file.Services, func(s ast.Service) string { return s.Name }),
	)
	for _, p := range []string{"string", "int32", "int64", "float32", "float64"} {
		names[p] = struct{}{}
	}
	return names
}

func extract[T any, C comparable](ts []T, extract func(T) C) []C {
	cs := make([]C, len(ts))
	for i, t := range ts {
		cs[i] = extract(t)
	}
	return cs
}

type set[T comparable] map[T]struct{}

func (s set[T]) toSlice() []T {
	slice := make([]T, len(s))
	i := 0
	for t := range s {
		slice[i] = t
		i++
	}
	return slice
}

func buildSet[C comparable](slices ...[]C) set[C] {
	values := make(set[C])
	for _, slice := range slices {
		for _, value := range slice {
			values[value] = struct{}{}
		}
	}
	return values
}

func validateUnique[C comparable](slices ...[]C) error {
	seen := make(set[C])
	dupes := make(set[error])
	for _, slice := range slices {
		for i, l := range slice {
			for _, r := range slice[i+1:] {
				if l != r {
					continue
				}
				if _, ok := seen[l]; ok {
					continue
				}
				seen[l] = struct{}{}
				dupes[fmt.Errorf("duplicate value %v", l)] = struct{}{}
			}
		}
	}

	return errors.Join(dupes.toSlice()...)
}
