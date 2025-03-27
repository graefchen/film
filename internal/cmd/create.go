package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		name := askFor("What is the name of your blog? ")
		println(name)
	},
}

func askFor(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}
