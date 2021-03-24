package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Build-time variables, that are set at compile time with ldflags
var (
	Version   = "unknown-version"
	GitCommit = "unknown-commit"
	BuildTime = "unknown-buildtime"
)

func init() {
	cli.VersionPrinter = printVersion
}

func printVersion(c *cli.Context) {
	fmt.Fprintf(os.Stdout, "Version: %s \nGit Commit: %s \nBuildTime: %s\n", Version, GitCommit, BuildTime)
	os.Exit(0)
}

var VersionCmd = cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "version",
	Action: func(c *cli.Context) error {
		fmt.Fprintf(os.Stdout, "Version: %s \nGit Commit: %s \nBuildTime: %s\n", Version, GitCommit, BuildTime)
		os.Exit(0)
		return nil
	},
}
