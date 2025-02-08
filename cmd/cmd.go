package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	root := cobra.Command{
		Use:   "run",
		Short: "Run the a script on the super computer",
		Run: func(cmd *cobra.Command, args []string) {
			script := `#!/bin/bash
			echo "Running script on supercomputer"
			# Add your script commands here
			`

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Enter your script (end with an empty line):")
			var scriptLines []string
			for {
				line, err := reader.ReadString('\n')
				if err != nil && err != io.EOF {
					fmt.Println("Error reading input:", err)
					return
				}
				line = strings.TrimSpace(line)
				if line == "" {
					break
				}
				scriptLines = append(scriptLines, line)
			}
			script = strings.Join(scriptLines, "\n")

			fmt.Println("Executing the following script:")
			fmt.Println(script)
		},
	}
	return &root
}
