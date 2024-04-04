package cli

import (
	"fmt"

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

		var optlen int
		for _, option := range info.Options {
			optlen = max(optlen, len(plugin.OptionName(name, option.Name)))
		}
		optTmpl := fmt.Sprintf("%%-%ds%%s%%s\n", optlen+4)

		for _, option := range info.Options {
			defstr := ""
			if option.Default != nil {
				defstr = fmt.Sprintf(" [default: %v]", option.Default)
			}
			fmt.Printf(optTmpl, plugin.OptionName(name, option.Name), option.Description, defstr)
		}

		return nil
	})
}
