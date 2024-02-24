package plugin

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	"github.com/egoodhall/servo/pkg/ipc"
)

type ServocExtension interface {
	Info(*ipc.InfoRequest, *ipc.InfoResponse) error
	Generate(*ipc.GenerateRequest, *ipc.GenerateResponse) error
}

func ServeRequest(extension ServocExtension) error {
	// stdout is used for IPC. We can make both stderr and stdout go
	// to the same place while the extension is being served, which
	// will let (most) normal printing & logging continue to work.
	stdout := os.Stdout
	os.Stdout = os.Stderr
	defer func() { os.Stdout = stdout }()

	srv, err := newServer(extension)
	if err != nil {
		return fmt.Errorf("construct server: %w", err)
	}

	return srv.ServeRequest(jsonrpc.NewServerCodec(ipc.NewConn(os.Stdin, stdout)))
}

func newServer(ext ServocExtension) (*rpc.Server, error) {
	srv := rpc.NewServer()
	if err := srv.RegisterName("ServocPlugin", ext); err != nil {
		return nil, err
	}
	return srv, nil
}
