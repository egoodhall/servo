package ipc

import "github.com/egoodhall/servo/pkg/ast"

type GenerateRequest struct {
	// Options that should be applied to all files
	Options []*ast.Option[any]
	// AST for all parsed
	Files []*ast.File
}

type GenerateResponse struct {
}

type InfoRequest struct {
}

type InfoResponse struct {
	Name    string
	Version string
	Options []*Option[any]
}

type Option[T any] struct {
	Name        string
	Default     any
	Description string
}
