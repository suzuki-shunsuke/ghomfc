package run

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

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

func loginAndName(m *github.Member) string {
	if m.Name == "" || m.Name == m.Login {
		return m.Login
	}
	return fmt.Sprintf("%s (%s)", m.Login, m.Name)
}

func outputTable(stdout io.Writer, members []*github.Member) {
	builder := &strings.Builder{}
	builder.Grow(56 + len(members)*50) //nolint:mnd
	fmt.Fprintln(builder, "Rank | Login (Name) | Number of Followers")
	fmt.Fprintln(builder, "--- | --- | ---")
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
		fmt.Fprintf(builder, "%d | [%s](https://github.com/%s) | %d\n", rank, loginAndName(m), m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	fmt.Fprint(stdout, builder.String())
}

func output(stdout io.Writer, param *Param, members []*github.Member) error {
	switch param.Format {
	case "json", "":
		return outputJSON(stdout, members)
	case "table":
		outputTable(stdout, members)
		return nil
	default:
		return errors.New("unsupported output format")
	}
}
