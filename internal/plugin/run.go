package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/egoodhall/servo/pkg/ipc"
	"golang.org/x/sync/errgroup"
)

func RunAll(parent context.Context, run func(string, Client) error) error {
	grp, ctx := errgroup.WithContext(parent)
	plugins, err := Discover()
	if err != nil {
		return err
	}
	for _, path := range plugins {
		p := path
		grp.Go(func() error {
			if err := runPlugin(ctx, p, run); err != nil {
				return fmt.Errorf("%s: %w", filepath.Base(p), err)
			}
			return nil
		})
	}
	return grp.Wait()
}

func runPlugin(parent context.Context, plugin string, run func(string, Client) error) error {
	grp, ctx := errgroup.WithContext(parent)

	// Create the command for running the plugin
	cmd := exec.CommandContext(ctx, plugin)
	cmd.Stderr = os.Stderr

	// Create a "connection" to the plugin using STDIO.
	conn, err := ipc.NewClientConn(cmd)
	if err != nil {
		return fmt.Errorf("open IPC connection: %w", err)
	}

	// Run the plugin. It should wait to serve 1 request and
	// then exit automatically.
	grp.Go(cmd.Run)

	// Run whatever we're doing with the plugin.
	grp.Go(func() error {
		client := newClient(conn)
		if err := run(filepath.Base(plugin), client); err != nil {
			return err
		}
		return client.close()
	})

	return grp.Wait()
}
