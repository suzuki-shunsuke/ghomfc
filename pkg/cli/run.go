package cli

import (
	"fmt"
	"os"

	"github.com/suzuki-shunsuke/ghomfc/pkg/controller/run"
	"github.com/suzuki-shunsuke/ghomfc/pkg/github"
	"github.com/urfave/cli/v3"
)

type runCommand struct {
	// logE *logrus.Entry
}

func (rc *runCommand) command() *cli.Command {
	return &cli.Command{
		Name:      "run",
		Usage:     "List a GitHub Organization members and sort them by the number of followers",
		UsageText: "ghomfc run <GitHub Organization Name>",
		Description: `List a GitHub Organization members and sort them by the number of followers.

$ ghomfc run <GitHub Organization Name>
`,
		Action: rc.action,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Usage: "output format (json, table)",
				Value: "json",
			},
		},
	}
}

func (rc *runCommand) action(ctx context.Context, c *cli.Command) error {
	gh, err := github.New(c.Context, os.Getenv("GITHUB_TOKEN"))
	if err != nil {
		return fmt.Errorf("create a GitHub Client: %w", err)
	}
	return run.Run(c.Context, os.Stdout, gh, &run.Param{ //nolint:wrapcheck
		Org:    c.Args().First(),
		Format: c.String("format"),
	})
}
