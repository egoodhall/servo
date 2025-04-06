package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/egoodhall/servo/cliutil"
	"github.com/egoodhall/servo/option"
	"github.com/egoodhall/servo/plugin"
)

type optionsCmd struct {
}

func (oc *optionsCmd) Run() error {
	ctx, cancel := cliutil.NewSignalCtx()
	defer cancel()

	options, err := gatherOptions(ctx)
	if err != nil {
		return err
	}

	var optlen int
	for _, option := range options {
		optlen = max(optlen, len(option.Name))
	}

	optTmpl := fmt.Sprintf("%%-%ds%%s%%s\n", optlen+4)
	for _, option := range options.SortedByName() {
		defstr := ""
		if option.Value != nil {
			defstr = fmt.Sprintf(" [default: %v]", option.Value)
		}
		fmt.Printf(optTmpl, option.Name, option.Description, defstr)
	}

	return nil
}

func gatherOptions(ctx context.Context) (option.Set, error) {
	opts, err := plugin.CollectAll(ctx, func(name string, client plugin.Client, results chan<- *option.Option) error {
		info, err := client.Info()
		if err != nil {
			return fmt.Errorf("info: %w", err)
		}
		for _, opt := range info.Options {
			opt.Name = option.QualifiedName(plugin.TrimmedName(name), opt.Name)
			results <- opt
		}
		log.Printf("%s - %+v", name, info)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return option.NewSet(opts), nil
}
