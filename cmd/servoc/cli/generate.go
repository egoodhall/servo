package cli

import (
	"fmt"
	"os"

	"github.com/egoodhall/servo/cliutil"
	"github.com/egoodhall/servo/ipc"
	"github.com/egoodhall/servo/parser"
	"github.com/egoodhall/servo/plugin"
)

type generateCmd struct {
	Options map[string]string `name:"option" short:"o"`
	Files   []*os.File        `arg:"" required:"" name:"files" type:"existingFile"`
}

func (gc *generateCmd) Run() error {
	ctx, cancel := cliutil.NewSignalCtx()
	defer cancel()

	options, err := gatherOptions(ctx)
	if err != nil {
		return err
	}

	files, err := parser.Files(gc.Files, options)
	if err != nil {
		return err
	}

	astOptions, err := options.ToAst(gc.Options)
	if err != nil {
		return err
	}

	return plugin.RunAll(ctx, func(name string, client plugin.Client) error {
		_, err := client.Generate(&ipc.GenerateRequest{
			Options: astOptions,
			Files:   files,
		})
		if err != nil {
			return fmt.Errorf("generate: %w", err)
		}

		return nil
	})
}
