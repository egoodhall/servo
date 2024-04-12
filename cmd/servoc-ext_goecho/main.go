package main

import (
	"github.com/egoodhall/servo/cmd/servoc-ext_goecho/options"
	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/pkg/plugin"
)

func main() {
	cliutil.RunFunc(func() error {
		return plugin.ServeRequest[options.Options](new(GoNrpcPlugin))
	})
}

type GoNrpcPlugin struct{}
