package main

import (
	"github.com/egoodhall/servo/pkg/ipc"
	"github.com/egoodhall/servo/pkg/plugin"
)

var Options = []ipc.Option[any]{
	{
		Name:        "enabled",
		Default:     false,
		Description: "Whether the plugin should run",
	},
	{
		Name:        "package",
		Description: "The go package name",
	},
	{
		Name:        "directory",
		Description: "The directory to write to",
	},
}

func (x *GoJsonPlugin) Info() (*ipc.InfoResponse, error) {
	options, err := plugin.ReadOptionsDescriptor[GenerateOptions]()
	if err != nil {
		return nil, err
	}
	return &ipc.InfoResponse{
		Name:    "Servoc Go JSON Plugin",
		Version: "v0.0.0",
		Options: options,
	}, nil
}
