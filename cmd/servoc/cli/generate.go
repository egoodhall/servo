package cli

import (
	"fmt"
	"os"

	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/internal/parser"
	"github.com/egoodhall/servo/internal/plugin"
	"github.com/egoodhall/servo/pkg/ipc"
)

type generateCmd struct {
	Options map[string]string `name:"option" short:"o"`
	Files   []*os.File        `arg:"" required:"" name:"files" type:"existingFile"`
}

func (gc *generateCmd) Run() error {
	files, err := parser.Files(gc.Files...)
	if err != nil {
		return err
	}

	ctx, cancel := cliutil.NewSignalCtx()
	defer cancel()

	return plugin.RunAll(ctx, func(name string, client plugin.Client) error {
		_, err := client.Generate(&ipc.GenerateRequest{
			Options: plugin.ToOptions(gc.Options),
			Files:   files,
		})
		if err != nil {
			return fmt.Errorf("generate: %w", err)
		}

		return nil
	})
}
