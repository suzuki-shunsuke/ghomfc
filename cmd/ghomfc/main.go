package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/suzuki-shunsuke/ghomfc/pkg/cli"
	"github.com/suzuki-shunsuke/slog-error/slogerr"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
)

var (
	version = ""
	commit  = "" //nolint:gochecknoglobals
	date    = "" //nolint:gochecknoglobals
)

func main() {
	if code := core(); code != 0 {
		os.Exit(code)
	}
}

func core() int {
	logger := slogutil.New(&slogutil.InputNew{
		Name:    "ghomfc",
		Version: version,
		Out:     os.Stderr,
	})
	runner := cli.Runner{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		LDFlags: &cli.LDFlags{
			Version: version,
			Commit:  commit,
			Date:    date,
		},
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := runner.Run(ctx, os.Args...); err != nil {
		slogerr.WithError(logger, err).Error("ghomfc failed")
		return 1
	}
	return 0
}
