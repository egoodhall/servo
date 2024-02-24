package cli

import (
	"fmt"
	"strings"

	"github.com/egoodhall/servo/internal/cliutil"
	"github.com/egoodhall/servo/internal/plugin"
)

type optionsCmd struct {
}

func (oc *optionsCmd) Run() error {
	ctx, cancel := cliutil.NewSignalCtx()
	defer cancel()

	return plugin.RunAll(ctx, func(name string, client plugin.Client) error {
		info, err := client.Info()
		if err != nil {
			return fmt.Errorf("info request: %w", err)
		}

		for _, option := range info.Options {
			optDesc := new(strings.Builder)
			optDesc.WriteString(plugin.Name(name))
			optDesc.WriteString(".")
			optDesc.WriteString(option.Name)
			optDesc.WriteString("\t")
			if option.Default != "" {
				optDesc.WriteString("(def: '")
				optDesc.WriteString(option.Name)
				optDesc.WriteString("')\t")
			}
			optDesc.WriteString(option.Description)
			fmt.Println(optDesc.String())
		}

		return nil
	})
}
