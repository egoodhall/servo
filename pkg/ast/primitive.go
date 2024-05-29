package ast

import (
	"slices"
)

var Primitives = []string{"string", "bool", "int32", "int64", "float32", "float64", "byte", "timestamp"}

func IsPrimitive(name string) bool {
	return slices.Contains(Primitives, name)
}
