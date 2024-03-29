package ast

type File struct {
	Name     string         `json:"name"`
	Options  []*Option[any] `json:"options"`
	Messages []*Message     `json:"types"`
	Enums    []*Enum        `json:"enums"`
	Services []*Service     `json:"services"`
}

type Option[T any] struct {
	Name  string `json:"name"`
	Value T      `json:"value"`
}

type Message struct {
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

type Field struct {
	Name     string `json:"name"`
	Type     Type   `json:"type"`
	Optional bool   `json:"optional"`
}

type MapType struct {
	KeyType   *ScalarType `json:"key"`
	ValueType *ScalarType `json:"value"`
}

type ListType struct {
	ElementType *ScalarType `json:"element"`
}

type ScalarType struct {
	Name string `json:"name"`
}

type Service struct {
	Name string `json:"name"`
	Rpcs []*Rpc `json:"rpcs"`
	Pubs []*Pub `json:"pubs"`
}

type Rpc struct {
	Name     string `json:"name"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Pub struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type Enum struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
