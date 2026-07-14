package command

import (
	"time"

	"github.com/urfave/cli/v3"

	"github.com/uplang/isnow.go/internal/app"
)

func tzFlag(dst *string) cli.Flag {
	return &cli.StringFlag{
		Name:        "tz",
		Usage:       "evaluation time zone (IANA name)",
		Sources:     cli.EnvVars("ISNOW_TZ"),
		Destination: dst,
	}
}

func instantFlag(name string, dst *string) cli.Flag {
	return &cli.StringFlag{Name: name, Usage: "instant (RFC 3339); default now", Destination: dst}
}

// resolveInstant resolves the evaluation instant from --at/--from and --tz.
func resolveInstant(env *app.Env, value, tz string) (time.Time, error) {
	loc, err := app.ParseZone(tz)
	if err != nil {
		return time.Time{}, err
	}
	if value == "" {
		return env.Now().In(loc), nil
	}
	return app.ParseInstant(value, loc)
}

// firstArg returns the isnow argument or a usage error.
func firstArg(c *cli.Command) (string, error) {
	src := c.Args().First()
	if src == "" {
		return "", usageError{msg: "an isnow argument is required"}
	}
	return src, nil
}
