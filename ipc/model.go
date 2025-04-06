package ipc

import (
	"github.com/egoodhall/servo/ast"
	"github.com/egoodhall/servo/option"
)

type GenerateRequest struct {
	// Options that should be applied to all files
	Options []*ast.Option
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
	Options []*option.Option
}
