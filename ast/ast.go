package ast

type File struct {
	Name     string     `json:"name"`
	Options  []*Option  `json:"options,omitempty"`
	Messages []*Message `json:"messages,omitempty"`
	Unions   []*Union   `json:"unions,omitempty"`
	Enums    []*Enum    `json:"enums,omitempty"`
	Services []*Service `json:"services,omitempty"`
	Aliases  []*Alias   `json:"aliases,omitempty"`
}

type Option struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type Alias struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Union struct {
	Name    string    `json:"name"`
	Members []*Member `json:"fields"`
}

type Member struct {
	Name string     `json:"name"`
	Type ScalarType `json:"type"`
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
}

type Rpc struct {
	Name     string `json:"name"`
	Request  string `json:"request"`
	Response string `json:"response"`
}

type Enum struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
