package cli

import (
	"fmt"
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/internal/plugin"
	"github.com/egoodhall/servo/pkg/ipc"
)

type optionsCmd struct {
}

func (oc *optionsCmd) Run() error {
	ctx, cancel := cliutil.NewSignalCtx()
	defer cancel()

	return plugin.RunAll(ctx, func(conn io.ReadWriteCloser) error {
		request := new(ipc.InfoRequest)
		response := new(ipc.GenerateResponse)

		client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
		if err := client.Call("ServocPlugin.Info", request, response); err != nil {
			return fmt.Errorf("info request: %w", err)
		}
		return client.Close()
	})
}
