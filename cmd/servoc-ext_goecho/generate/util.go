package generate

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/iancoleman/strcase"
)

type serverNames struct {
	Constructor       string
	HttpCompat        string
	RegisterFunc      string
	RegisterGroupFunc string
}

func newServerNames(svc *ast.Service) serverNames {
	return serverNames{
		Constructor:       "New" + strcase.ToCamel(svc.Name) + "HttpServer",
		HttpCompat:        strcase.ToLowerCamel(svc.Name) + "HttpServer",
		RegisterFunc:      "Register" + strcase.ToCamel(svc.Name) + "RPCs",
		RegisterGroupFunc: "Register" + strcase.ToCamel(svc.Name) + "RPCsGroup",
	}
}

type clientNames struct {
	DelegatingConstructor string
	Constructor           string
	Implementation        string
}

func newClientNames(svc *ast.Service) clientNames {
	return clientNames{
		DelegatingConstructor: "NewDelegating" + strcase.ToCamel(svc.Name) + "HttpClient",
		Constructor:           "New" + strcase.ToCamel(svc.Name) + "HttpClient",
		Implementation:        strcase.ToLowerCamel(svc.Name) + "HttpClient",
	}
}

type testClientNames struct {
	Constructor    string
	Implementation string
}

func newTestClientNames(svc *ast.Service) testClientNames {
	return testClientNames{
		Constructor:    "New" + strcase.ToCamel(svc.Name) + "TestHttpClient",
		Implementation: strcase.ToLowerCamel(svc.Name) + "HttpTestClient",
	}
}

type rpcNames struct {
	MethodName string
	HttpPath   string
}

func newRpcNames(svc *ast.Service, rpc *ast.Rpc) rpcNames {
	return rpcNames{
		MethodName: strcase.ToCamel(rpc.Name),
		HttpPath:   fmt.Sprintf("%s/%s", strcase.ToKebab(svc.Name), strcase.ToKebab(rpc.Name)),
	}
}

func newHeaderComment(message string, args ...any) *jen.Statement {
	comment := fmt.Sprintf("// %s //", fmt.Sprintf(message, args...))
	aboveAndBelow := strings.Repeat("/", len(comment))

	return jen.Comment(aboveAndBelow).Line().
		Comment(comment).Line().
		Comment(aboveAndBelow).Line()
}
