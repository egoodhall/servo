package option

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func QualifiedName(plugin, name string) string {
	return fmt.Sprintf("%s.%s", plugin, name)
}

type Type uint16

const (
	keyShift        = 6
	collectionShift = 12
	typeMask        = 0b111111
)

const (
	// Supported primitive types
	Bool    = Type(reflect.Bool)
	Int     = Type(reflect.Int)
	Int8    = Type(reflect.Int8)
	Int16   = Type(reflect.Int16)
	Int32   = Type(reflect.Int32)
	Int64   = Type(reflect.Int64)
	Uint    = Type(reflect.Uint)
	Uint8   = Type(reflect.Uint8)
	Uint16  = Type(reflect.Uint16)
	Uint32  = Type(reflect.Uint32)
	Uint64  = Type(reflect.Uint64)
	Float32 = Type(reflect.Float32)
	Float64 = Type(reflect.Float64)
	String  = Type(reflect.String)

	// Supported collection types
	Map  Type = 1 << collectionShift
	List Type = 2 << collectionShift
)

func (t Type) KeyType() Type {
	return (t >> keyShift) & typeMask
}

func (t Type) ValueType() Type {
	return t & typeMask
}

func (typ Type) IsMap() bool {
	return typ&Map != 0
}

func (typ Type) IsList() bool {
	return typ&List != 0
}

func (typ Type) IsCollection() bool {
	return typ.IsMap() || typ.IsList()
}

func (typ Type) Reflect() reflect.Type {
	if typ.IsMap() {
		return reflect.MapOf(typ.KeyType().Reflect(), typ.ValueType().Reflect())
	}
	if typ.IsList() {
		return reflect.SliceOf(typ.ValueType().Reflect())
	}
	switch typ.ValueType() {
	case Bool:
		return reflect.TypeFor[bool]()
	case Int:
		return reflect.TypeFor[int64]()
	case Int8:
		return reflect.TypeFor[int8]()
	case Int16:
		return reflect.TypeFor[int16]()
	case Int32:
		return reflect.TypeFor[int32]()
	case Int64:
		return reflect.TypeFor[int64]()
	case Uint:
		return reflect.TypeFor[uint64]()
	case Uint8:
		return reflect.TypeFor[uint8]()
	case Uint16:
		return reflect.TypeFor[uint16]()
	case Uint32:
		return reflect.TypeFor[uint32]()
	case Uint64:
		return reflect.TypeFor[uint64]()
	case Float32:
		return reflect.TypeFor[float32]()
	case Float64:
		return reflect.TypeFor[float64]()
	case String:
		return reflect.TypeFor[string]()
	}
	return nil
}

func MapType(key, value Type) Type {
	return Map | key<<6 | value
}

func ListType(elem Type) Type {
	return List | elem
}

func TypeFrom(typ reflect.Type) (Type, error) {
	switch typ.Kind() {
	case reflect.Bool:
		return Bool, nil
	case reflect.Int:
		return Int, nil
	case reflect.Int8:
		return Int8, nil
	case reflect.Int16:
		return Int16, nil
	case reflect.Int32:
		return Int32, nil
	case reflect.Int64:
		return Int64, nil
	case reflect.Uint:
		return Uint, nil
	case reflect.Uint8:
		return Uint8, nil
	case reflect.Uint16:
		return Uint16, nil
	case reflect.Uint32:
		return Uint32, nil
	case reflect.Uint64:
		return Uint64, nil
	case reflect.Float32:
		return Float32, nil
	case reflect.Float64:
		return Float64, nil
	case reflect.String:
		return String, nil
	case reflect.Slice, reflect.Array:
		elem, err := TypeFrom(typ.Elem())
		if err != nil {
			return 0, err
		}
		if elem.IsCollection() {
			return 0, errors.New("cannot nest map or list in slice")
		}
		return ListType(elem), nil
	case reflect.Map:
		key, keyerr := TypeFrom(typ.Key())
		val, valerr := TypeFrom(typ.Elem())
		if keyerr != nil || valerr != nil {
			return 0, errors.Join(keyerr, valerr)
		}

		errs := []error{}
		if key.IsCollection() {
			errs = append(errs, errors.New("cannot use collection as map key"))
		}
		if val.IsMap() {
			errs = append(errs, errors.New("cannot use map as map value"))
		}
		if len(errs) > 0 {
			return 0, errors.Join(errs...)
		}
		return MapType(key, val), nil
	default:
		return 0, fmt.Errorf("unsupported type: %s", typ.Kind())
	}
}

func Parse(value string, as Type) (any, error) {
	if as.IsList() {
		var err error
		pvals := make([]any, 0)
		for i, val := range strings.Split(value, ",") {
			pvals[i], err = Parse(val, as.ValueType())
			if err != nil {
				return nil, fmt.Errorf("list element %d: %w", i, err)
			}
		}
		return pvals, nil
	}

	if as.IsMap() {
		vals := reflect.MakeMap(as.Reflect())
		for _, pair := range strings.Split(value, ";") {
			k, v, found := strings.Cut(pair, "=")
			if !found {
				return nil, fmt.Errorf("invalid key-value pair: %s", pair)
			}
			key, keyerr := Parse(k, as.KeyType())
			val, valerr := Parse(v, as.ValueType())
			if keyerr != nil || valerr != nil {
				return nil, errors.Join(keyerr, valerr)
			}
			vals.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
		}
		return vals, nil
	}

	switch as.ValueType() {
	case Bool:
		return strconv.ParseBool(value)
	case Int:
		return strconv.ParseInt(value, 10, 64)
	case Int8:
		return strconv.ParseInt(value, 10, 8)
	case Int16:
		return strconv.ParseInt(value, 10, 16)
	case Int32:
		return strconv.ParseInt(value, 10, 32)
	case Int64:
		return strconv.ParseInt(value, 10, 64)
	case Uint:
		return strconv.ParseUint(value, 10, 64)
	case Uint8:
		return strconv.ParseUint(value, 10, 8)
	case Uint16:
		return strconv.ParseUint(value, 10, 16)
	case Uint32:
		return strconv.ParseUint(value, 10, 32)
	case Uint64:
		return strconv.ParseUint(value, 10, 64)
	case Float32:
		return strconv.ParseFloat(value, 32)
	case Float64:
		return strconv.ParseFloat(value, 64)
	case String:
		return value, nil
	default:
		return nil, fmt.Errorf("unsupported type: %v", as.ValueType())
	}
}

func From(field reflect.StructField) (*Option, error) {
	opt := Option{
		Name:        field.Tag.Get("name"),
		Description: field.Tag.Get("desc"),
	}

	if opt.Name == "" {
		opt.Name = field.Name
	}

	var err error
	opt.Type, err = TypeFrom(field.Type)
	if err != nil {
		return nil, err
	}

	if field.Tag.Get("default") != "" {
		opt.Value, err = Parse(field.Tag.Get("default"), opt.Type)
		if err != nil {
			return nil, err
		}
	}

	return &opt, nil
}

type Option struct {
	Name        string `json:"name"`
	Type        Type   `json:"type"`
	Description string `json:"description"`
	Value       any    `json:"value"`
}
