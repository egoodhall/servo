package plugin

import (
	"github.com/egoodhall/servo/pkg/ipc"
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Client interface {
	Info() (*ipc.InfoResponse, error)
	Generate(request *ipc.GenerateRequest) (*ipc.GenerateResponse, error)
}

func newClient(conn io.ReadWriteCloser) *client {
	return &client{jsonrpc.NewClient(conn)}
}

type client struct {
	c *rpc.Client
}

func (c client) Info() (*ipc.InfoResponse, error) {
	res := new(ipc.InfoResponse)
	return res, c.c.Call("ServocPlugin.Info", new(ipc.InfoRequest), res)
}

func (c client) Generate(request *ipc.GenerateRequest) (*ipc.GenerateResponse, error) {
	res := new(ipc.GenerateResponse)
	return res, c.c.Call("ServocPlugin.Generate", request, res)
}

func (c client) close() error {
	return c.c.Close()
}
