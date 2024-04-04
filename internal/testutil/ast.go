package testutil

import (
	"reflect"
	"strings"
	"testing"

	"github.com/egoodhall/servo/internal/parser"
	"github.com/egoodhall/servo/internal/textutil"
	"github.com/egoodhall/servo/pkg/ast"
)

func MustParse(t *testing.T, content string) *ast.File {
	t.Helper()
	file, err := parser.File("file.srvo", strings.NewReader(textutil.TrimDedent(content)))
	if err != nil {
		t.Fatalf("parse file: %s", err)
	}
	return file
}

func HasMessage(t *testing.T, actual *ast.File, expected ast.Message) {
	t.Helper()
	for _, actualType := range actual.Messages {
		if reflect.DeepEqual(actualType, expected) {
			return
		}
	}
	t.Fatalf("expected message was not found in parsed file:\nExpected: %+v\nActual:   %+v", expected, actual.Messages)
}

func HasOnlyMessage(t *testing.T, actual *ast.File, expected ast.Message) {
	t.Helper()
	if len(actual.Messages) != 1 {
		t.Fatalf("expected exactly 1 parsed message, but found %d", len(actual.Messages))
	}
	HasMessage(t, actual, expected)
}

func HasOption[T any](t *testing.T, actual *ast.File, expected ast.Option[T]) {
	t.Helper()
	for _, actualOption := range actual.Options {
		if reflect.DeepEqual(actualOption, expected) {
			return
		}
	}
	t.Fatalf("expected option was not found in parsed file:\nExpected: %+v\nActual:   %+v", expected, actual.Options)
}

func HasOnlyOption[T any](t *testing.T, actual *ast.File, expected ast.Option[T]) {
	t.Helper()
	if len(actual.Options) != 1 {
		t.Fatalf("expected exactly 1 parsed option, but found %d", len(actual.Options))
	}
	HasOption(t, actual, expected)
}

func HasService(t *testing.T, actual *ast.File, expected ast.Service) {
	t.Helper()
	for _, actualService := range actual.Services {
		if reflect.DeepEqual(actualService, expected) {
			return
		}
	}
	t.Fatalf("expected service was not found in parsed file:\nExpected: %+v\nActual:   %+v", expected, actual.Services)
}

func HasOnlyService(t *testing.T, actual *ast.File, expected ast.Service) {
	t.Helper()
	if len(actual.Services) != 1 {
		t.Fatalf("expected exactly 1 parsed service, but found %d", len(actual.Services))
	}
	HasService(t, actual, expected)
}
