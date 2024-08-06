package run

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/suzuki-shunsuke/ghomfc/pkg/github"
)

func createMembers(size int) []*github.Member {
	members := make([]*github.Member, size)
	for i := range size {
		a := strconv.Itoa(i)
		members[i] = &github.Member{
			Login:          "test" + a,
			Name:           "Test Test" + a,
			NumOfFollowers: i,
		}
	}
	return members
}

func BenchmarkOutputTable(b *testing.B) {
	data := []struct {
		name string
		size int
	}{
		{
			name: "10",
			size: 10,
		},
		{
			name: "1000",
			size: 1000,
		},
		{
			name: "10000",
			size: 10000,
		},
	}
	funcs := []struct {
		name string
		fn   func(stdout io.Writer, members []*github.Member) error
	}{
		{
			name: "outputTable(EveryPrint)",
			fn:   outputTable,
		},
		{
			name: "outputTable2(strings.BuilderGlow42)",
			fn:   outputTable2,
		},
		{
			name: "outputTable2(strings.BuilderGlow50)",
			fn:   outputTable7,
		},
		{
			name: "outputTable3(bytes.Buffer)",
			fn:   outputTable3,
		},
		{
			name: "outputTable4(strings.Join)",
			fn:   outputTable4,
		},
		{
			name: "outputTable5(strings.Join)",
			fn:   outputTable5,
		},
		{
			name: "outputTable6(string+)",
			fn:   outputTable6,
		},
	}
	f, err := os.Create(filepath.Join(b.TempDir(), "stdout"))
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	for _, d := range data {
		for _, fn := range funcs {
			b.Run(d.name+"/"+fn.name, func(b *testing.B) {
				members := createMembers(d.size)
				b.ResetTimer()
				for range b.N {
					fn.fn(f, members)
				}
			})
		}
	}
}
