package cmd

import (
	"film/internal/export"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use: "export",
	Run: func(cmd *cobra.Command, args []string) {
		export.Export()
	},
}
