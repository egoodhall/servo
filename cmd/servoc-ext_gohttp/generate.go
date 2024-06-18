package main

import (
	"os"

	"github.com/egoodhall/servo/cmd/servoc-ext_gohttp/generate"
	"github.com/egoodhall/servo/cmd/servoc-ext_gohttp/options"
	"github.com/egoodhall/servo/pkg/ast"
	"github.com/hashicorp/go-retryablehttp"
)

func (x *GoHttpPlugin) Generate(file *ast.File, options options.Options) error {
	retryablehttp.NewClient()
	if !options.Enabled {
		return nil
	}

	content, err := generate.ServicesAndClients(file, options)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(options.File, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return content.Render(f)
}
