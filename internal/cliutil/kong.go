package cliutil

import "github.com/alecthomas/kong"

type Runner interface {
	Run() error
}

func Run(r Runner) {
	ctx := kong.Parse(r)
	ctx.FatalIfErrorf(ctx.Run())
}

func RunFunc(fn func() error) {
	Run(&noArgsRunner{fn})
}

type noArgsRunner struct {
	fn func() error
}

func (nar *noArgsRunner) Run() error {
	return nar.fn()
}
