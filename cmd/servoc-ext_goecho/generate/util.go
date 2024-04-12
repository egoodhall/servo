package generate

import (
	"fmt"

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
		Constructor:    "NewTest" + strcase.ToCamel(svc.Name) + "HttpClient",
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
