package main

import (
	"github.com/suzuki-shunsuke/ghomfc/pkg/cli"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
)

var version = ""

func main() {
	urfave.Main("ghomfc", version, cli.Run)
}
