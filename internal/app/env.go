package app

import (
	"context"
	"io"
	"time"
)

// Clock supplies the current instant; injected so commands are testable.
type Clock func() time.Time

// Sleeper waits for d or until ctx is cancelled, returning ctx.Err() on cancel.
type Sleeper func(ctx context.Context, d time.Duration) error

// Spawner runs an external command; injected so `run` is testable.
type Spawner func(ctx context.Context, name string, args []string) error

// Env is the injected runtime for every command: clock, output streams, and the
// wait/spawn effects. Tests supply fakes; main wires the real ones.
type Env struct {
	Now   Clock
	Out   io.Writer
	Err   io.Writer
	Sleep Sleeper
	Spawn Spawner
	Zone  *time.Location
}

// RealSleep waits with a timer honoring context cancellation.
func RealSleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
