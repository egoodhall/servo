package cli

type Cli struct {
	Plugins  []string    `name:"plugin" short:"p" help:"Plugins that should be included during generation"`
	Generate generateCmd `cmd:"" name:"generate"`
	Options  optionsCmd  `cmd:"" name:"options"`
}
