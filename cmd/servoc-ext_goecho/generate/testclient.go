package generate

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/iancoleman/strcase"
)

func TestClient(gofile *jen.File, svc *ast.Service) {
	gofile.Add(generateTestClientConstructor(svc)).Line()
	gofile.Add(generateTestClientType(svc)).Line()

	for _, rpc := range svc.Rpcs {
		if rpc.Response != "" {
			gofile.Add(generateTestClientMethodWithResponse(svc, rpc)).Line()
		} else {
			gofile.Add(generateTestClientMethodWithoutResponse(svc, rpc)).Line()
		}
	}
}

func generateTestClientConstructor(svc *ast.Service) *jen.Statement {
	names := newTestClientNames(svc)

	return jen.Func().Id(names.Constructor).
		Params(jen.Id("svc").Id(svc.Name)).
		List(jen.Id(svc.Name)).
		Block(jen.Return(jen.Op("&").Id(names.Implementation).Values(jen.Id("svc"))))
}

func generateTestClientType(svc *ast.Service) *jen.Statement {
	names := newTestClientNames(svc)

	return jen.Var().Id("_").Id(svc.Name).Op("=").New(jen.Id(names.Implementation)).
		Line().
		Type().Id(names.Implementation).Struct(
		jen.Id("service").Id(svc.Name),
	)
}

func generateTestClientMethodWithResponse(svc *ast.Service, rpc *ast.Rpc) *jen.Statement {
	names := newTestClientNames(svc)
	sNames := newServerNames(svc)
	rNames := newRpcNames(svc, rpc)

	return jen.Func().Params(jen.Id("client").Op("*").Id(names.Implementation)).
		Id(strcase.ToCamel(rpc.Name)).
		Params(jen.Id("_").Qual(pkgContext, "Context"), jen.Id("request").Op("*").Id(rpc.Request)).
		Params(jen.Op("*").Id(rpc.Response), jen.Error()).
		Block(
			jen.Id("body").Op(":=").New(jen.Qual(pkgBytes, "Buffer")),
			jen.If(
				jen.Err().Op(":=").Qual(pkgJson, "NewEncoder").Params(jen.Id("body")).Dot("Encode").Params(jen.Id("request")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Nil(), jen.Err()),
			),
			jen.Id("req").Op(":=").Qual(pkgHttptest, "NewRequest").Params(
				jen.Qual(pkgHttp, "MethodPost"),
				jen.Lit(fmt.Sprintf("/%s/%s", strcase.ToKebab(svc.Name), strcase.ToKebab(rpc.Name))),
				jen.Id("body"),
			),
			jen.Id("req").Dot("Header").Dot("Set").Call(jen.Lit("Content-Type"), jen.Lit("application/json")),
			jen.List(jen.Id("res")).Op(":=").Qual(pkgHttptest, "NewRecorder").Call(),
			jen.Line(),
			jen.Id("ctx").Op(":=").
				Id(sNames.Constructor).Call(jen.Id("client").Dot("service")).
				Dot("NewContext").Call(jen.Id("req"), jen.Id("res")),
			jen.If(
				jen.Err().Op(":=").Parens(jen.Op("&").Id(sNames.HttpCompat).Values(jen.Id("client").Dot("service"))).Dot(rNames.MethodName).Call(jen.Id("ctx")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Nil(), jen.Err()),
			).Else().If(jen.Id("res").Dot("Code").Op("!=").Qual(pkgHttp, "StatusOk")).Block(
				jen.Return(jen.Nil(), jen.Qual(pkgErrors, "New").Call(
					jen.Lit("unexpected status code ").Op("+").Qual(pkgStrconv, "Itoa").Call(jen.Id("res").Dot("Code")),
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

func generateTestClientMethodWithoutResponse(svc *ast.Service, rpc *ast.Rpc) *jen.Statement {
	names := newTestClientNames(svc)
	sNames := newServerNames(svc)
	rNames := newRpcNames(svc, rpc)

	return jen.Func().Params(jen.Id("client").Op("*").Id(names.Implementation)).
		Id(strcase.ToCamel(rpc.Name)).
		Params(jen.Id("_").Qual(pkgContext, "Context"), jen.Id("request").Op("*").Id(rpc.Request)).
		Error().
		Block(
			jen.Id("body").Op(":=").New(jen.Qual(pkgBytes, "Buffer")),
			jen.If(
				jen.Err().Op(":=").Qual(pkgJson, "NewEncoder").Params(jen.Id("body")).Dot("Encode").Params(jen.Id("request")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			),
			jen.Id("req").Op(":=").Qual(pkgHttptest, "NewRequest").Params(
				jen.Qual(pkgHttp, "MethodPost"),
				jen.Lit(fmt.Sprintf("/%s/%s", strcase.ToKebab(svc.Name), strcase.ToKebab(rpc.Name))),
				jen.Id("body"),
			),
			jen.Id("req").Dot("Header").Dot("Set").Call(jen.Lit("Content-Type"), jen.Lit("application/json")),
			jen.List(jen.Id("res")).Op(":=").Qual(pkgHttptest, "NewRecorder").Call(),
			jen.Line(),
			jen.Id("ctx").Op(":=").
				Id(sNames.Constructor).Call(jen.Id("client").Dot("service")).
				Dot("NewContext").Call(jen.Id("req"), jen.Id("res")),
			jen.If(
				jen.Err().Op(":=").Parens(jen.Op("&").Id(sNames.HttpCompat).Values(jen.Id("client").Dot("service"))).Dot(rNames.MethodName).Call(jen.Id("ctx")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			).Else().If(jen.Id("res").Dot("Code").Op("!=").Qual(pkgHttp, "StatusNoContent")).Block(
				jen.Return(jen.Qual(pkgErrors, "New").Call(
					jen.Lit("unexpected status code ").Op("+").Qual(pkgStrconv, "Itoa").Call(jen.Id("res").Dot("Code")),
				)),
			),
			jen.Return(jen.Nil()),
		)
}
