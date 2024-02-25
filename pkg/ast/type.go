package ast

//go:generate polyjson -type Type -package . -file type.gen.go

type Type interface {
	iType()
}

func (ScalarType) iType() {}
func (MapType) iType()    {}
func (ListType) iType()   {}
