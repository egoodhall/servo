package main

import "github.com/egoodhall/servo/pkg/ipc"

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
			{
				Name:        "directory",
				Description: "The directory to write to",
			},
		},
	}
	return nil
}
