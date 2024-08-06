package cli

import (
	"os"

	"github.com/suzuki-shunsuke/ghomfc/pkg/controller/run"
	"github.com/suzuki-shunsuke/ghomfc/pkg/github"
	"github.com/urfave/cli/v2"
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
	}
}

func (rc *runCommand) action(c *cli.Context) error {
	gh := github.New(c.Context, os.Getenv("GITHUB_TOKEN"))
	return run.Run(c.Context, os.Stdout, gh, &run.Param{ //nolint:wrapcheck
		Org: c.Args().First(),
	})
}
