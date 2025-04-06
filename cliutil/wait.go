package cliutil

import (
	"context"
	"os"
	"os/signal"
)

func NewSignalCtx() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt)
}
