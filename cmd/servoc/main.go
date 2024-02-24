package main

import (
	"github.com/alecthomas/kong"
	"github.com/egoodhall/servo/cmd/servoc/cli"
)

func main() {
	ktx := kong.Parse(new(cli.Cli))
	ktx.FatalIfErrorf(ktx.Run())
}
