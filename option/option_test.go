package option_test

import (
	"reflect"
	"testing"

	"github.com/egoodhall/servo/option"
	"github.com/egoodhall/servo/testutil"
)

func TestTypeOf(t *testing.T) {
	tests := []struct {
		input    reflect.Type
		expected option.Type
	}{
		{reflect.TypeFor[bool](), option.Bool},
		{reflect.TypeFor[int](), option.Int},
		{reflect.TypeFor[int8](), option.Int8},
		{reflect.TypeFor[int16](), option.Int16},
		{reflect.TypeFor[int32](), option.Int32},
		{reflect.TypeFor[int64](), option.Int64},
		{reflect.TypeFor[uint](), option.Uint},
		{reflect.TypeFor[uint8](), option.Uint8},
		{reflect.TypeFor[uint16](), option.Uint16},
		{reflect.TypeFor[uint32](), option.Uint32},
		{reflect.TypeFor[uint64](), option.Uint64},
		{reflect.TypeFor[float32](), option.Float32},
		{reflect.TypeFor[float64](), option.Float64},
		{reflect.TypeFor[string](), option.String},
	}

	for _, test := range tests {
		t.Run(test.input.String(), func(t *testing.T) {
			actual, err := option.TypeFrom(test.input)
			testutil.AssertNoError(t, err)

			if actual != test.expected {
				t.Errorf("expected %v, got %v", test.expected, actual)
			}
		})
	}
}
