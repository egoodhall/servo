package main

import (
	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/pkg/plugin"
)

type Options struct {
	Enabled bool   `json:"enabled" default:"false" desc:"If false, the gostruct plugin will not generate code"`
	Package string `json:"package" desc:"The name of the package to use for the generated go file"`
	File    string `json:"file" default:"gotype.gen.go" desc:"The file to generate to generate code in"`
}

func main() {
	cliutil.RunFunc(func() error {
		return plugin.ServeRequest[Options](new(GoStructPlugin))
	})
}

type GoStructPlugin struct{}
