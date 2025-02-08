package main

import (
	"os"
)

func main() {
	err := root.rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
