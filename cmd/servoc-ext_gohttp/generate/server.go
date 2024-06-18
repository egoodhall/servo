package generate

import (
	"github.com/dave/jennifer/jen"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/iancoleman/strcase"
)

func Server(gofile *jen.File, svc *ast.Service) {
	gofile.Add(newHeaderComment("%s HTTP server", svc.Name))
	gofile.Add(generateServerInterface(svc)).Line()
	gofile.Add(generateServerConstructor(svc)).Line()
	gofile.Add(generateServerRegisterFunc(svc)).Line()
	gofile.Add(generateServerRegisterGroupFunc(svc)).Line()
	gofile.Add(generateServerImpl(svc)).Line()

	for _, rpc := range svc.Rpcs {
		if rpc.Response != "" {
			gofile.Add(generateServerRpcHandler(svc, rpc)).Line()
		} else {
			gofile.Add(generateServerPubHandler(svc, rpc)).Line()
		}
	}
}

func generateServerInterface(svc *ast.Service) *jen.Statement {
	names := newServerNames(svc)
	return jen.Type().Id(names.HttpEndpoints).InterfaceFunc(func(g *jen.Group) {
		for _, rpc := range svc.Rpcs {
			rpcNames := newRpcNames(svc, rpc)
			method := jen.Id(rpcNames.MethodName).Params(
				jen.Qual(pkgEcho, "Context"),
				jen.Op("*").Id(rpc.Request),
			)
			if rpc.Response == "" {
				method.Error()
			} else {
				method.Params(jen.Op("*").Id(rpc.Response), jen.Error())
			}
			g.Add(method)
		}
	})
}

func generateServerConstructor(svc *ast.Service) *jen.Statement {
	names := newServerNames(svc)

	return jen.Func().Id(names.Constructor).Params(
		jen.Id("svc").Id(names.HttpEndpoints),
	).Op("*").Qual(pkgEcho, "Echo").Block(
		jen.Id("srv").Op(":=").Qual(pkgEcho, "New").Call(),
		jen.Id(names.RegisterFunc).Call(jen.Id("svc"), jen.Id("srv")),
		jen.Return(jen.Id("srv")),
	)
}

func generateServerRegisterFunc(svc *ast.Service) *jen.Statement {
	names := newServerNames(svc)

	return jen.Func().Id(names.RegisterFunc).Params(
		jen.Id("svc").Id(names.HttpEndpoints),
		jen.Id("srv").Op("*").Qual(pkgEcho, "Echo"),
	).Block(
		jen.Id(names.RegisterGroupFunc).Call(jen.Id("svc"), jen.Id("srv").Dot("Group").Call(jen.Lit("/"))),
	)
}

func generateServerRegisterGroupFunc(svc *ast.Service) *jen.Statement {
	names := newServerNames(svc)

	return jen.Func().Id(names.RegisterGroupFunc).Params(
		jen.Id("svc").Id(names.HttpEndpoints),
		jen.Id("srv").Op("*").Qual(pkgEcho, "Group"),
	).BlockFunc(func(g *jen.Group) {
		g.Id("compat").Op(":=").Op("&").Id(names.HttpAdapter).Values(jen.Id("svc"))
		for _, rpc := range svc.Rpcs {
			rNames := newRpcNames(svc, rpc)

			g.Id("srv").Dot("POST").Call(jen.Lit(rNames.HttpPath), jen.Id("compat").Dot(rNames.MethodName))
		}
	})
}

func generateServerImpl(svc *ast.Service) *jen.Statement {
	names := newServerNames(svc)

	return jen.Type().Id(names.HttpAdapter).Struct(
		jen.Id("svc").Id(names.HttpEndpoints),
	)
}

func generateServerRpcHandler(svc *ast.Service, rpc *ast.Rpc) *jen.Statement {
	sNames := newServerNames(svc)
	rNames := newRpcNames(svc, rpc)

	return jen.Commentf("HTTP compatibility wrapper for %s.%s.", svc.Name, rpc.Name).Line().
		Func().Parens(jen.Id("s").Op("*").Id(sNames.HttpAdapter)).Id(rNames.MethodName).Params(jen.Id("c").Qual(pkgEcho, "Context")).Error().Block(
		jen.Id("req").Op(":=").New(jen.Id(rpc.Request)),
		jen.If(
			jen.Err().Op(":=").Id("c").Dot("Bind").Params(jen.Id("req")),
			jen.Err().Op("!=").Nil(),
		).Block(
			jen.Return(jen.Err()),
		),
		jen.List(jen.Id("res"), jen.Err()).Op(":=").Id("s").Dot("svc").Dot(strcase.ToCamel(rpc.Name)).Params(
			jen.Id("c"),
			jen.Id("req"),
		),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Qual(pkgEcho, "NewHTTPError").Params(
				jen.Qual(pkgHttp, "StatusBadRequest"),
				jen.Err(),
			)),
		),
		jen.Return(jen.Id("c").Dot("JSON").Params(
			jen.Qual(pkgHttp, "StatusOK"),
			jen.Id("res"),
		)),
	)
}

func generateServerPubHandler(svc *ast.Service, rpc *ast.Rpc) *jen.Statement {
	sNames := newServerNames(svc)
	rNames := newRpcNames(svc, rpc)

	return jen.Commentf("HTTP compatibility wrapper for %s.%s.", svc.Name, rpc.Name).Line().
		Func().Parens(jen.Id("s").Op("*").Id(sNames.HttpAdapter)).Id(rNames.MethodName).Params(jen.Id("c").Qual(pkgEcho, "Context")).Error().
		Block(
			jen.Id("req").Op(":=").New(jen.Id(rpc.Request)),
			jen.If(
				jen.Err().Op(":=").Id("c").Dot("Bind").Params(jen.Id("req")),
				jen.Err().Op("!=").Nil(),
			).Block(
				jen.Return(jen.Err()),
			),
			jen.Err().Op(":=").Id("s").Dot("svc").Dot(strcase.ToCamel(rpc.Name)).Params(
				jen.Id("c"),
				jen.Id("req"),
			),
			jen.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Qual(pkgEcho, "NewHTTPError").Params(
					jen.Qual(pkgHttp, "StatusBadRequest"),
					jen.Err(),
				)),
			),
			jen.Return(jen.Id("c").Dot("NoContent").Params(
				jen.Qual(pkgHttp, "StatusNoContent"),
			)),
		)
}
