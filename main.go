package main

import (
	"fmt"
	"os"
)

// Build-time variables, that are set at compile time with ldflags
var (
	Version   = "unknown-version"
	GitCommit = "unknown-commit"
	BuildTime = "unknown-buildtime"
)

func main() {
	fmt.Fprintf(os.Stdout, "You are using mimic with version:%s, gitcommit:%s, build:%s", Version, GitCommit, BuildTime)
}
