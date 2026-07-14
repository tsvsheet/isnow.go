package command

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/uplang/isnow.go/internal/app"
	"github.com/uplang/isnow.go/internal/domain"
)

// canonCommand prints the canonical form of an isnow.
func canonCommand(env *app.Env) *cli.Command {
	return &cli.Command{
		Name:      "canon",
		Usage:     "print the canonical form",
		ArgsUsage: "<isnow>",
		Action: func(_ context.Context, c *cli.Command) error {
			src, err := firstArg(c)
			if err != nil {
				return err
			}
			canonical, err := domain.Canon(src)
			if err != nil {
				return err
			}
			fmt.Fprintln(env.Out, canonical)
			return nil
		},
	}
}

// explainCommand prints the canonical form and English description.
func explainCommand(env *app.Env) *cli.Command {
	return &cli.Command{
		Name:      "explain",
		Usage:     "print the canonical form and an English description",
		ArgsUsage: "<isnow>",
		Action: func(_ context.Context, c *cli.Command) error {
			src, err := firstArg(c)
			if err != nil {
				return err
			}
			v, err := domain.Describe(src)
			if err != nil {
				return err
			}
			fmt.Fprintf(env.Out, "%s\n%s\n", v.Canonical, v.Explain)
			return nil
		},
	}
}
