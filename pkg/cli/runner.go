package cli

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/vcmd"
	"github.com/urfave/cli/v3"
)

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	LDFlags *LDFlags
	LogE    *logrus.Entry
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (r *Runner) Run(ctx context.Context, args ...string) error {
	app := cli.Command{
		Name:                  "ghomfc",
		Usage:                 "GitHub Organization Members' Followers Counter",
		Version:               r.LDFlags.Version + " (" + r.LDFlags.Commit + ")",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			(&runCommand{
				// logE: r.LogE,
			}).command(),
			(&completionCommand{
				logE:   r.LogE,
				stdout: r.Stdout,
			}).command(),
			vcmd.New(&vcmd.Command{
				Name:    "ghomfc",
				Version: r.LDFlags.Version,
				SHA:     r.LDFlags.Commit,
			}),
		},
	}

	return app.Run(ctx, args) //nolint:wrapcheck
}
