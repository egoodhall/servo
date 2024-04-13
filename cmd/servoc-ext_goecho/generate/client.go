package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/iancoleman/strcase"
)

func Client(gofile *jen.File, svc *ast.Service) {
	gofile.Add(newHeaderComment("%s HTTP client", svc.Name))
	gofile.Add(generateClientConstructor(svc)).Line()
	gofile.Add(generateDelegatingClientConstructor(svc)).Line()
	gofile.Add(generateClientType(svc)).Line()

	for _, rpc := range svc.Rpcs {
		if rpc.Response != "" {
			gofile.Add(generateClientMethodWithResponse(svc, rpc)).Line()
		} else {
			gofile.Add(generateClientMethodWithoutResponse(svc, rpc)).Line()
		}
	}
}

func generateClientConstructor(svc *ast.Service) *jen.Statement {
	names := newClientNames(svc)

	return jen.Func().Id(names.Constructor).
		Params(jen.Id("baseUrl").String()).
		List(jen.Id(svc.Name)).
		Block(
			jen.Return(jen.Id(names.DelegatingConstructor).Params(
				jen.Id("baseUrl"),
				jen.New(jen.Qual(pkgHttp, "Client"))),
			),
		)
}

func generateDelegatingClientConstructor(svc *ast.Service) *jen.Statement {
	names := newClientNames(svc)

	return jen.Func().Id(names.DelegatingConstructor).
		Params(jen.Id("baseUrl").String(), jen.Id("delegate").Op("*").Qual(pkgHttp, "Client")).
		List(jen.Id(svc.Name)).
		Block(
			jen.Return(jen.Op("&").Id(names.Implementation).Values(jen.Id("baseUrl"), jen.Id("delegate"))),
		)
}

func generateClientType(svc *ast.Service) *jen.Statement {
	names := newClientNames(svc)

	return jen.Var().Id("_").Id(svc.Name).Op("=").New(jen.Id(names.Implementation)).
		Line().
		Type().Id(names.Implementation).Struct(
		jen.Id("baseUrl").String(),
		jen.Id("delegate").Op("*").Qual(pkgHttp, "Client"),
	)
}

func generateClientMethodWithResponse(svc *ast.Service, rpc *ast.Rpc) *jen.Statement {
	names := newClientNames(svc)

	return jen.Func().Params(jen.Id("client").Op("*").Id(names.Implementation)).
		Id(strcase.ToCamel(rpc.Name)).
		Params(jen.Id("ctx").Qual(pkgContext, "Context"), jen.Id("request").Op("*").Id(rpc.Request)).
		Params(jen.Op("*").Id(rpc.Response), jen.Error()).
		Block(
			jen.List(jen.Id("u"), jen.Err()).Op(":=").Qual(pkgUrl, "JoinPath").Params(
				jen.Id("client").Dot("baseUrl"),
				jen.Lit(fmt.Sprintf("/%s/%s", strcase.ToKebab(svc.Name), strcase.ToKebab(rpc.Name))),
			),
			jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()),
			),
			jen.Line(),
			jen.Id("body").Op(":=").New(jen.Qual(pkgBytes, "Buffer")),
			jen.If(
				jen.Err().Op(":=").Qual(pkgJson, "NewEncoder").Params(jen.Id("body")).Dot("Encode").Params(jen.Id("request")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Nil(), jen.Err()),
			),
			jen.Line(),
			jen.List(jen.Id("req"), jen.Err()).Op(":=").Qual(pkgHttp, "NewRequestWithContext").Params(
				jen.Id("ctx"),
				jen.Qual(pkgHttp, "MethodPost"),
				jen.Id("u"),
				jen.Id("body"),
			),
			jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()),
			),
			jen.Id("req").Dot("Header").Dot("Set").Call(jen.Lit("Content-Type"), jen.Lit("application/json")),
			jen.Line(),
			jen.List(jen.Id("res"), jen.Err()).Op(":=").Id("client").Dot("delegate").Dot("Do").Call(jen.Id("req")),
			jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()),
			),
			jen.Defer().Id("res").Dot("Body").Dot("Close").Call(),
			jen.Line(),
			jen.If(jen.Id("res").Dot("StatusCode").Op("!=").Qual(pkgHttp, "StatusOK")).Block(
				jen.Return(jen.Nil(), jen.Qual(pkgErrors, "New").Call(
					jen.Lit("unexpected status code ").Op("+").Qual(pkgStrconv, "Itoa").Call(jen.Id("res").Dot("StatusCode")),
				)),
			),
			jen.Line(),
			jen.Id("response").Op(":=").New(jen.Id(rpc.Response)),
			jen.Return(
				jen.Id("response"),
				jen.Qual(pkgJson, "NewDecoder").Params(jen.Id("res").Dot("Body")).Dot("Decode").Params(jen.Id("response")),
			),
		)
}

func generateClientMethodWithoutResponse(svc *ast.Service, rpc *ast.Rpc) *jen.Statement {
	names := newClientNames(svc)

	return jen.Func().Params(jen.Id("client").Op("*").Id(names.Implementation)).
		Id(strcase.ToCamel(rpc.Name)).
		Params(jen.Id("ctx").Qual(pkgContext, "Context"), jen.Id("request").Op("*").Id(rpc.Request)).
		Error().
		Block(
			jen.List(jen.Id("u"), jen.Err()).Op(":=").Qual(pkgUrl, "JoinPath").Params(
				jen.Id("client").Dot("baseUrl"),
				jen.Lit(fmt.Sprintf("/%s/%s", strcase.ToKebab(svc.Name), strcase.ToKebab(rpc.Name))),
			),
			jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Err()),
			),
			jen.Line(),
			jen.Id("body").Op(":=").New(jen.Qual(pkgBytes, "Buffer")),
			jen.If(
				jen.Err().Op(":=").Qual(pkgJson, "NewEncoder").Params(jen.Id("body")).Dot("Encode").Params(jen.Id("request")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			),
			jen.Line(),
			jen.List(jen.Id("req"), jen.Err()).Op(":=").Qual(pkgHttp, "NewRequestWithContext").Params(
				jen.Id("ctx"),
				jen.Qual(pkgHttp, "MethodPost"),
				jen.Id("u"),
				jen.Id("body"),
			),
			jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Err()),
			),
			jen.Id("req").Dot("Header").Dot("Set").Call(jen.Lit("Content-Type"), jen.Lit("application/json")),
			jen.Line(),
			jen.List(jen.Id("res"), jen.Err()).Op(":=").Id("client").Dot("delegate").Dot("Do").Call(jen.Id("req")),
			jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Err()),
			),
			jen.Defer().Id("res").Dot("Body").Dot("Close").Call(),
			jen.Line(),
			jen.If(jen.Id("res").Dot("StatusCode").Op("!=").Qual(pkgHttp, "StatusNoContent")).Block(
				jen.Return(jen.Qual(pkgErrors, "New").Call(
					jen.Lit("unexpected status code ").Op("+").Qual(pkgStrconv, "Itoa").Call(jen.Id("res").Dot("StatusCode")),
				)),
			),
			jen.Return(jen.Nil()),
		)
}
