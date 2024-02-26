package plugin

import (
	"fmt"
	"github.com/egoodhall/servo/internal/plugin"
	"github.com/egoodhall/servo/pkg/ast"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/egoodhall/servo/pkg/ipc"
)

type ServocPlugin[O any] interface {
	Info() (*ipc.InfoResponse, error)
	Generate(file *ast.File, options O) error
}

type ServocExtension interface {
	Info(*ipc.InfoRequest, *ipc.InfoResponse) error
	Generate(*ipc.GenerateRequest, *ipc.GenerateResponse) error
}

func ServeRequest[O any](plugin ServocPlugin[O]) error {
	// stdout is used for IPC. We can make both stderr and stdout go
	// to the same place while the plugin is being served, which
	// will let (most) normal printing & logging continue to work.
	stdout := os.Stdout
	os.Stdout = os.Stderr
	defer func() { os.Stdout = stdout }()

	srv := rpc.NewServer()
	if err := srv.RegisterName("ServocPlugin", &pluginCompat[O]{plugin}); err != nil {
		return err
	}

	return srv.ServeRequest(jsonrpc.NewServerCodec(ipc.NewConn(os.Stdin, stdout)))
}

type pluginCompat[O any] struct {
	plugin ServocPlugin[O]
}

func (pc *pluginCompat[O]) Info(req *ipc.InfoRequest, res *ipc.InfoResponse) error {
	info, err := pc.plugin.Info()
	if err != nil {
		return err
	}
	*res = *info
	return nil
}

func (pc *pluginCompat[O]) Generate(req *ipc.GenerateRequest, res *ipc.GenerateResponse) error {
	info, err := pc.plugin.Info()
	if err != nil {
		return err
	}

	for _, file := range req.Files {
		options := new(O)
		if err := parseOptions[O](plugin.Name(os.Args[0]), info.Options, file.Options, req.Options, options); err != nil {
			return fmt.Errorf("parse options: %w", err)
		}

		if err := pc.plugin.Generate(file, *options); err != nil {
			return fmt.Errorf("generate code: %w", err)
		}
	}
	return nil
}
