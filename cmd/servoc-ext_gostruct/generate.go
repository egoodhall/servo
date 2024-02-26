package main

import (
	"github.com/dave/jennifer/jen"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/iancoleman/strcase"
	"os"
)

type GenerateOptions struct {
	Enabled bool   `json:"enabled" default:"false" desc:"If false, the gostruct plugin will not generate code"`
	Package string `json:"package" desc:"The name of the package to use for the generated go file"`
	File    string `json:"file" default:"." desc:"The directory to generate code in"`
}

func (x *GoJsonPlugin) Generate(file *ast.File, options GenerateOptions) error {
	if !options.Enabled {
		return nil
	}

	content, err := generateFile(file, options)
	if err != nil {
		return err
	}

	if err := content.Render(os.Stderr); err != nil {
		return err
	}
	return nil
}

func generateFile(file *ast.File, options GenerateOptions) (*jen.File, error) {
	gofile := jen.NewFile("servoc")

	for _, message := range file.Messages {
		gofile.Type().Id(strcase.ToCamel(message.Name)).StructFunc(func(g *jen.Group) {
			for _, field := range message.Fields {
				switch ft := field.Type.(type) {
				case ast.ScalarType:
					g.Id(strcase.ToCamel(field.Name)).Id(ft.Name).Tag(map[string]string{"json": field.Name})
				case ast.ListType:
					g.Id(strcase.ToCamel(field.Name)).Op("[]").Id(ft.ElementType.Name).Tag(map[string]string{"json": field.Name})
				case ast.MapType:
					g.Id(strcase.ToCamel(field.Name)).Map(jen.Id(ft.KeyType.Name)).Id(ft.ValueType.Name).Tag(map[string]string{"json": field.Name})
				}
			}
		})
		gofile.Line()
	}

	for _, enum := range file.Enums {
		gofile.Type().Id(strcase.ToCamel(enum.Name)).String()
		gofile.Const().DefsFunc(func(g *jen.Group) {
			for _, value := range enum.Values {
				g.Id(strcase.ToCamel(value)).Id(strcase.ToCamel(enum.Name)).Op("=").Lit(value)
			}
		})
	}

	for _, svc := range file.Services {
		gofile.Type().Id(strcase.ToCamel(svc.Name)).InterfaceFunc(func(g *jen.Group) {
			for _, rpc := range svc.Rpcs {
				g.Id(strcase.ToCamel(rpc.Name)).
					Params(jen.Op("*").Id(strcase.ToCamel(rpc.Request))).
					Params(jen.Op("*").Id(strcase.ToCamel(rpc.Response)), jen.Err())
			}
			for _, pub := range svc.Pubs {
				g.Id(strcase.ToCamel(pub.Name)).Params(jen.Op("*").Id(strcase.ToCamel(pub.Message))).Err()
			}
		})
	}

	return gofile, nil
}
