package run

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/suzuki-shunsuke/ghomfc/pkg/github"
)

type Param struct {
	Org    string
	Format string
}

type GitHub interface {
	ListOrgMembers(ctx context.Context, org string) ([]*github.Member, error)
}

func Run(ctx context.Context, stdout io.Writer, gh GitHub, param *Param) error {
	// Find and read a configuration file.
	// List org members.
	// Sort org members by the number of followers.
	// Output the result.
	if param.Org == "" {
		return errors.New("the GitHub Organization name is required")
	}
	members, err := gh.ListOrgMembers(ctx, param.Org)
	if err != nil {
		return fmt.Errorf("list the GitHub Organization members: %w", err)
	}
	sort.Slice(members, func(i, j int) bool {
		return members[i].NumOfFollowers > members[j].NumOfFollowers
	})
	if err := output(stdout, param, members); err != nil {
		return err
	}
	return nil
}

func outputJSON(stdout io.Writer, members []*github.Member) error {
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(members); err != nil {
		return fmt.Errorf("output the result: %w", err)
	}
	return nil
}

func outputTable(stdout io.Writer, members []*github.Member) error {
	fmt.Fprintln(stdout, "Rank | Login (Name) | Number of Followers")
	fmt.Fprintln(stdout, "--- | --- | ---")
	prevNumOfFollowers := -1
	prevRank := 0
	for i, m := range members {
		var rank int
		if prevNumOfFollowers == m.NumOfFollowers {
			rank = prevRank
		} else {
			rank = i + 1
			prevNumOfFollowers = m.NumOfFollowers
		}
		fmt.Fprintf(stdout, "%d | [%s (%s)](https://github.com/%s) | %d\n", rank, m.Login, m.Name, m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	return nil
}

func output(stdout io.Writer, param *Param, members []*github.Member) error {
	switch param.Format {
	case "json", "":
		return outputJSON(stdout, members)
	case "table":
		return outputTable(stdout, members)
	default:
		return errors.New("unsupported output format")
	}
}
