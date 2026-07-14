package app

import (
	"context"
	"os"
	"os/exec"
)

// RealSpawn runs a command, wiring its stdio to the process, honoring context.
func RealSpawn(ctx context.Context, name string, args []string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
