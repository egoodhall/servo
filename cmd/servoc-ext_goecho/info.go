package main

import (
	"github.com/egoodhall/servo/cmd/servoc-ext_goecho/options"
	"github.com/egoodhall/servo/pkg/ipc"
	"github.com/egoodhall/servo/pkg/plugin"
)

func (x *GoNrpcPlugin) Info() (*ipc.InfoResponse, error) {
	options, err := plugin.ReadOptionsDescriptor[options.Options]()
	if err != nil {
		return nil, err
	}

	return &ipc.InfoResponse{
		Name:    "Servoc Go Echo Plugin",
		Version: "v0.0.0",
		Options: options,
	}, nil
}
