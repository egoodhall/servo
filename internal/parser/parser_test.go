package parser_test

import (
	"testing"

	"github.com/egoodhall/servo/internal/testutil"
	"github.com/egoodhall/servo/pkg/ast"
)

func TestParsesTypes(t *testing.T) {
	file := testutil.MustParse(t, `
		// A request message
		message Request {
			// The message to send
			message: string;
			// The message size
			size: int32;
			// When the message was created
			createdAt: int64;
		}
	`)

	testutil.HasOnlyMessage(t, file, ast.Message{
		Name: "Request",
		Fields: []*ast.Field{
			{Name: "message", Type: ast.ScalarType{Name: "string"}},
			{Name: "size", Type: ast.ScalarType{Name: "int32"}},
			{Name: "createdAt", Type: ast.ScalarType{Name: "int64"}},
		},
	})
}

func TestParsesOptions(t *testing.T) {
	file := testutil.MustParse(t, `
		// Define the go package
		option go_package = "github.com/egoodhall/servo";
	`)

	testutil.HasOnlyOption(t, file, ast.Option[string]{
		Name:  "go_package",
		Value: "github.com/egoodhall/servo",
	})
}

func TestParsesRpcServices(t *testing.T) {
	file := testutil.MustParse(t, `
		message EchoRequest {}
		message EchoResponse {}
		service EchoService {
			rpc echo(EchoRequest): EchoResponse;
		}
	`)

	testutil.HasOnlyService(t, file, ast.Service{
		Name: "EchoService",
		Rpcs: []*ast.Rpc{
			{
				Name:     "echo",
				Request:  "EchoRequest",
				Response: "EchoResponse",
			},
		},
	})
}

func TestParsesPublisherServices(t *testing.T) {
	file := testutil.MustParse(t, `
		message Message {}
		service PublisherService {
			pub publish(Message);
		}
	`)

	testutil.HasOnlyService(t, file, ast.Service{
		Name: "PublisherService",
		Pubs: []*ast.Pub{
			{
				Name:    "publish",
				Message: "Message",
			},
		},
	})
}

func TestParsesEverything(t *testing.T) {
	file := testutil.MustParse(t, `
		option go_package = "github.com/egoodhall/servo";

		message EchoRequest {
			message: string;
		}

		message EchoResponse {
			message: string;
		}

		service EchoService {
			rpc echo(EchoRequest): EchoResponse;
		}

		message Message {
			message: string;
		}

		service PublisherService {
			pub publish(Message);
		}
	`)

	testutil.HasMessage(t, file, ast.Message{
		Name: "EchoRequest",
		Fields: []*ast.Field{
			{Name: "message", Type: ast.ScalarType{Name: "string"}},
		},
	})

	testutil.HasMessage(t, file, ast.Message{
		Name: "EchoResponse",
		Fields: []*ast.Field{
			{Name: "message", Type: ast.ScalarType{Name: "string"}},
		},
	})

	testutil.HasService(t, file, ast.Service{
		Name: "EchoService",
		Rpcs: []*ast.Rpc{
			{Name: "echo", Request: "EchoRequest", Response: "EchoResponse"},
		},
	})

	testutil.HasMessage(t, file, ast.Message{
		Name: "Message",
		Fields: []*ast.Field{
			{Name: "message", Type: ast.ScalarType{Name: "string"}},
		},
	})

	testutil.HasService(t, file, ast.Service{
		Name: "PublisherService",
		Pubs: []*ast.Pub{
			{Name: "publish", Message: "Message"},
		},
	})
}
