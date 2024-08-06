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
	switch param.Format {
	case "json", "":
		encoder := json.NewEncoder(stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(members); err != nil {
			return fmt.Errorf("output the result: %w", err)
		}
		return nil
	case "table":
		fmt.Fprintln(stdout, "Rank | Login (Name) | Number of Followers")
		fmt.Fprintln(stdout, "--- | --- | ---")
		for i, m := range members {
			fmt.Fprintf(stdout, "%d | %s (%s) | %d\n", i+1, m.Login, m.Name, m.NumOfFollowers)
		}
		return nil
	default:
		return errors.New("unsupported output format")
	}
	return nil
}
