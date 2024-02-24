package main

import (
	"fmt"

	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/pkg/ipc"
	"github.com/egoodhall/servo/pkg/plugin"
)

func main() {
	cliutil.RunFunc(func() error {
		return plugin.ServeRequest(new(GoJsonPlugin))
	})
}

type GoJsonPlugin struct{}

func (x *GoJsonPlugin) Info(req *ipc.InfoRequest, res *ipc.InfoResponse) error {
	*res = ipc.InfoResponse{
		Name:    "Servoc Go JSON Plugin",
		Version: "v0.0.0",
		Options: []ipc.Option{
			{
				Name:        "enabled",
				Default:     "false",
				Description: "Whether the plugin should run",
			},
			{
				Name:        "package",
				Description: "The go package name",
			},
		},
	}
	return nil
}

func (x *GoJsonPlugin) Generate(req *ipc.GenerateRequest, res *ipc.GenerateResponse) error {
	fmt.Printf("%+v\n", req)
	return nil
}
