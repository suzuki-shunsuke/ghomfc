package cli

import (
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
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
	compiledDate, err := time.Parse(time.RFC3339, r.LDFlags.Date)
	if err != nil {
		compiledDate = time.Now()
	}
	app := cli.App{
		Name:                 "ghomfc",
		Usage:                "GitHub Organization Members' Followers Counter",
		Version:              r.LDFlags.Version + " (" + r.LDFlags.Commit + ")",
		Compiled:             compiledDate,
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			(&versionCommand{}).command(),
			(&runCommand{
				// logE: r.LogE,
			}).command(),
			(&completionCommand{
				logE:   r.LogE,
				stdout: r.Stdout,
			}).command(),
		},
	}

	return app.RunContext(ctx, args) //nolint:wrapcheck
}
