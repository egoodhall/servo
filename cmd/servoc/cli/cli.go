package cli

type Cli struct {
	Generate generateCmd `cmd:"" name:"generate"`
	Options  optionsCmd  `cmd:"" name:"options"`
}
