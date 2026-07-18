// Command isnow tests, derives, explains, schedules, and serves isnow
// date/time patterns. See https://tsvsheet.github.io/docs.isnow.go/.
package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/tsvsheet/isnow.go/internal/app"
	"github.com/tsvsheet/isnow.go/internal/command"
)

var (
	osExit = os.Exit
	args   = os.Args
)

// version is the application version, set via ldflags: -X main.version=1.0.0.
var version = "dev"

func main() { osExit(run(args)) }

// run builds the command tree with the real environment and maps the result to
// an exit code; it is a thin, testable shim (no os.Exit inside).
func run(args []string) int {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	env := &app.Env{
		Now:   time.Now,
		Out:   os.Stdout,
		Err:   os.Stderr,
		Sleep: app.RealSleep,
		Spawn: app.RealSpawn,
	}
	root := command.Root(env)
	root.Version = version
	return command.Report(env.Err, root.Run(ctx, args))
}
