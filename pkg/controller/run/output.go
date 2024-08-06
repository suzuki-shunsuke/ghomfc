package run

import (
	"bytes"
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

func outputTable2(stdout io.Writer, members []*github.Member) error {
	builder := &strings.Builder{}
	builder.Grow(56 + len(members)*42)
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
		fmt.Fprintf(builder, "%d | [%s (%s)](https://github.com/%s) | %d\n", rank, m.Login, m.Name, m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	fmt.Fprint(stdout, builder.String())
	return nil
}

func outputTable7(stdout io.Writer, members []*github.Member) error {
	builder := &strings.Builder{}
	builder.Grow(56 + len(members)*50)
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
		fmt.Fprintf(builder, "%d | [%s (%s)](https://github.com/%s) | %d\n", rank, m.Login, m.Name, m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	fmt.Fprint(stdout, builder.String())
	return nil
}

func outputTable3(stdout io.Writer, members []*github.Member) error {
	buf := &bytes.Buffer{}
	fmt.Fprintln(buf, "Rank | Login (Name) | Number of Followers")
	fmt.Fprintln(buf, "--- | --- | ---")
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
		fmt.Fprintf(buf, "%d | [%s (%s)](https://github.com/%s) | %d\n", rank, m.Login, m.Name, m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	fmt.Fprint(stdout, buf.String())
	return nil
}

func outputTable4(stdout io.Writer, members []*github.Member) error {
	lines := make([]string, 2+len(members))
	lines[0] = "Rank | Login (Name) | Number of Followers"
	lines[1] = "--- | --- | ---"
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
		lines[i+2] = fmt.Sprintf("%d | [%s (%s)](https://github.com/%s) | %d", rank, m.Login, m.Name, m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	fmt.Fprintln(stdout, strings.Join(lines, "\n"))
	return nil
}

func outputTable5(stdout io.Writer, members []*github.Member) error {
	lines := make([]string, 1+len(members))
	lines[0] = "Rank | Login (Name) | Number of Followers\n--- | --- | ---\n"
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
		lines[i+1] = fmt.Sprintf("%d | [%s (%s)](https://github.com/%s) | %d\n", rank, m.Login, m.Name, m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	fmt.Fprint(stdout, strings.Join(lines, ""))
	return nil
}

func outputTable6(stdout io.Writer, members []*github.Member) error {
	text := "Rank | Login (Name) | Number of Followers\n--- | --- | ---\n"
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
		text += fmt.Sprintf("%d | [%s (%s)](https://github.com/%s) | %d\n", rank, m.Login, m.Name, m.Login, m.NumOfFollowers)
		prevRank = rank
	}
	fmt.Fprint(stdout, text)
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
