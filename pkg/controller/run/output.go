package run

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/suzuki-shunsuke/ghomfc/pkg/github"
)

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
