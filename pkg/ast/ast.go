package ast

type File struct {
	Options  []Option  `json:"options"`
	Messages []Message `json:"types"`
	Enums    []Enum    `json:"enums"`
	Services []Service `json:"services"`
}

type Coords [2]int

type Option struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Message struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}

type Type struct {
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
}

func (t Type) String() string {
	if t.Optional {
		return t.Type + "?"
	}
	return t.Type
}

type Primitive string

func (p Primitive) String() string {
	return string(p)
}

const (
	String  Primitive = "string"
	Int32   Primitive = "int32"
	Int64   Primitive = "int64"
	Float32 Primitive = "float32"
	Float64 Primitive = "float64"
)

type Service struct {
	Name string `json:"name"`
	Rpcs []Rpc  `json:"rpcs"`
	Pubs []Pub  `json:"pubs"`
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
