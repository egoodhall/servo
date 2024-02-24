package cli

import (
	"fmt"
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/internal/parser"
	"github.com/egoodhall/servo/internal/plugin"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/egoodhall/servo/pkg/ipc"
)

type generateCmd struct {
	Plugins []string   `name:"plugin" short:"p"`
	Files   []*os.File `arg:"" required:"" name:"files" type:"existingFile"`
}

func (gc *generateCmd) Run() error {
	files, err := parser.Files(gc.Files...)
	if err != nil {
		return err
	}

	ctx, cancel := cliutil.NewSignalCtx()
	defer cancel()

	return plugin.RunAll(ctx, func(conn io.ReadWriteCloser) error {
		request := &ipc.GenerateRequest{
			Options: make([]ast.Option, 0),
			Files:   files,
		}
		response := new(ipc.GenerateResponse)

		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
		if err := client.Call("ServocPlugin.Generate", request, response); err != nil {
			return fmt.Errorf("generate request: %w", err)
		}
		return client.Close()
	})
}
