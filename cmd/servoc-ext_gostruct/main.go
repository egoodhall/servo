package main

import (
	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/pkg/plugin"
)

func main() {
	cliutil.RunFunc(func() error {
		return plugin.ServeRequest(new(GoJsonPlugin))
	})
}

type GoJsonPlugin struct{}
