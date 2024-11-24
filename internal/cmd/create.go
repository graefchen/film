package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("What is the name of your blog? ")
		var input string
		fmt.Scanln(&input)
		fmt.Print(input)
	},
}
