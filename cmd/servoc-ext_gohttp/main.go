package main

import (
	"github.com/egoodhall/servo/cmd/servoc-ext_gohttp/options"
	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/pkg/plugin"
)

func main() {
	cliutil.RunFunc(func() error {
		return plugin.ServeRequest[options.Options](new(GoHttpPlugin))
	})
}

type GoHttpPlugin struct{}
