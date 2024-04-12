package main

import (
	"os"

	"github.com/dave/jennifer/jen"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/iancoleman/strcase"
)

func (x *GoStructPlugin) Generate(file *ast.File, options Options) error {
	if !options.Enabled {
		return nil
	}

	content, err := generateFile(file, options)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(options.File, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return content.Render(f)
}

func generateFile(file *ast.File, options Options) (*jen.File, error) {
	gofile := jen.NewFile(options.Package)
	gofile.PackageComment("Code generated by servoc (gostruct plugin). DO NOT EDIT.")

	gofile.Type().DefsFunc(func(g *jen.Group) {
		for _, alias := range file.Aliases {
			g.Id(alias.Name).Add(renderType(alias.Type))
		}
	})

	gofile.Line()

	gofile.Type().DefsFunc(func(g *jen.Group) {
		for i, message := range file.Messages {
			if i != 0 {
				g.Line()
			}
			g.Id(strcase.ToCamel(message.Name)).StructFunc(func(g *jen.Group) {
				for _, field := range message.Fields {
					g.Add(renderField(field))
				}
			})
		}
	})

	gofile.Line()

	for _, enum := range file.Enums {
		gofile.Type().Id(strcase.ToCamel(enum.Name)).String()
		gofile.Const().DefsFunc(func(g *jen.Group) {
			for _, value := range enum.Values {
				g.Id(enum.Name + "_" + strcase.ToCamel(value)).Id(strcase.ToCamel(enum.Name)).Op("=").Lit(value)
			}
		})
	}

	for _, union := range file.Unions {
		gofile.Type().Id(strcase.ToCamel(union.Name)).StructFunc(func(g *jen.Group) {
			g.Id(union.Name + "Type").String().Tag(map[string]string{"json": "@type"})
			for _, member := range union.Members {
				g.Id(strcase.ToCamel(member.Name)).
					Op("*").Id(member.Type.Name).
					Tag(map[string]string{"json": member.Name + ",omitempty"})
			}
		}).Line()
	}

	gofile.Type().DefsFunc(func(g *jen.Group) {
		for i, svc := range file.Services {
			if i != 0 {
				g.Line()
			}
			g.Id(strcase.ToCamel(svc.Name)).InterfaceFunc(func(g *jen.Group) {
				for _, rpc := range svc.Rpcs {
					if rpc.Response != "" {
						g.Id(strcase.ToCamel(rpc.Name)).
							Params(jen.Qual("context", "Context"), getMethodType(rpc.Request)).
							Params(getMethodType(rpc.Response), jen.Error())
					} else {
						g.Id(strcase.ToCamel(rpc.Name)).
							Params(jen.Qual("context", "Context"), getMethodType(rpc.Request)).
							Error()
					}
				}
			})
		}
	})

	return gofile, nil
}

func renderField(field *ast.Field) *jen.Statement {
	switch ft := field.Type.(type) {
	case ast.ScalarType:
		s := jen.Id(strcase.ToCamel(field.Name))
		if field.Optional {
			s = s.Op("*")
		}
		s.Add(renderType(ft.Name)).Tag(renderTag(field))
		return s
	case ast.ListType:
		return jen.Id(strcase.ToCamel(field.Name)).Op("[]").Add(renderType(ft.ElementType.Name)).Tag(renderTag(field))
	case ast.MapType:
		return jen.Id(strcase.ToCamel(field.Name)).Map(renderType(ft.KeyType.Name)).Add(renderType(ft.ValueType.Name)).Tag(renderTag(field))
	}
	return nil
}

func renderType(name string) *jen.Statement {
	if !ast.IsPrimitive(name) {
		return jen.Id(name)
	}
	if name == "timestamp" {
		return jen.Qual("time", "Time")
	}
	return jen.Id(name)
}

func renderTag(field *ast.Field) map[string]string {
	if field.Optional {
		return map[string]string{"json": field.Name + ",omitempty"}
	}
	return map[string]string{"json": field.Name}
}

func getMethodType(name string) *jen.Statement {
	if ast.IsPrimitive(name) {
		return jen.Id(name)
	}
	return jen.Op("*").Id(name)
}
