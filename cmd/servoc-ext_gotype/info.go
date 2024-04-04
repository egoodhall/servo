package main

import (
	"github.com/egoodhall/servo/pkg/ipc"
	"github.com/egoodhall/servo/pkg/plugin"
)

func (x *GoStructPlugin) Info() (*ipc.InfoResponse, error) {
	options, err := plugin.ReadOptionsDescriptor[Options]()
	if err != nil {
		return nil, err
	}

	return &ipc.InfoResponse{
		Name:    "Servoc Go JSON Plugin",
		Version: "v0.0.0",
		Options: options,
	}, nil
}
