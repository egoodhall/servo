package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/iancoleman/strcase"
	"os"
)

func (x *GoJsonPlugin) Generate(file *ast.File, options Options) error {
	if !options.Enabled {
		return nil
	}

	fmt.Printf("%+v\n", options)

	content, err := generateFile(file, options)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(options.File, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return content.Render(f)
}

func generateFile(file *ast.File, options Options) (*jen.File, error) {
	gofile := jen.NewFile(options.Package)

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
					Params(jen.Op("*").Id(strcase.ToCamel(rpc.Response)), jen.Error())
			}
			for _, pub := range svc.Pubs {
				g.Id(strcase.ToCamel(pub.Name)).Params(jen.Op("*").Id(strcase.ToCamel(pub.Message))).Error()
			}
		})
	}

	return gofile, nil
}
