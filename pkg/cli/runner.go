package cli

import (
	"context"

	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, _ *slogutil.Logger, env *urfave.Env) error {
	return urfave.Command(env, &cli.Command{ //nolint:wrapcheck
		Name:  "ghomfc",
		Usage: "GitHub Organization Members' Followers Counter",
		Commands: []*cli.Command{
			(&runCommand{}).command(),
		},
	}).Run(ctx, env.Args)
}
