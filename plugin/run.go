package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/egoodhall/servo/ipc"
	"golang.org/x/sync/errgroup"
)

func RunAll(parent context.Context, run func(string, Client) error) error {
	_, err := CollectAll(parent, func(name string, client Client, tc chan<- any) error {
		return run(name, client)
	})
	return err
}

func CollectAll[T any](parent context.Context, run func(string, Client, chan<- T) error) ([]T, error) {
	grp, ctx := errgroup.WithContext(parent)
	plugins, err := Discover()
	if err != nil {
		return nil, err
	}

	tc := make(chan T)
	for _, path := range plugins {
		p := path
		grp.Go(func() error {
			if err := runPlugin(ctx, p, tc, run); err != nil {
				return fmt.Errorf("%s: %w", filepath.Base(p), err)
			}
			return nil
		})
	}

	go func() {
		grp.Wait()
		close(tc)
	}()

	ts := make([]T, 0)
	for t := range tc {
		ts = append(ts, t)
	}
	return ts, nil
}

func runPlugin[T any](parent context.Context, plugin string, results chan<- T, run func(string, Client, chan<- T) error) error {
	grp, ctx := errgroup.WithContext(parent)

	// Create the command for running the plugin
	cmd := exec.CommandContext(ctx, plugin)
	cmd.Stderr = os.Stderr

	// Create a "connection" to the plugin using STDIN/STDOUT.
	conn, err := ipc.NewClientConn(cmd)
	if err != nil {
		return fmt.Errorf("open IPC connection: %w", err)
	}

	// Run the plugin. It should wait to serve 1 request and
	// then exit automatically.
	grp.Go(cmd.Run)

	// Run whatever we're doing with the plugin.
	client := newClient(conn)
	if err := run(filepath.Base(plugin), client, results); err != nil {
		return err
	} else if err := grp.Wait(); err != nil {
		return err
	} else {
		return client.close()
	}
}
