package main

import (
	"os"

	"github.com/OmBiradar/supercomputer_cli/cmd/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
