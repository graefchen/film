package cmd

import (
	"film/internal/config"
	"fmt"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use: "export",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.FilmConfig()

		outDir := config.GetSiteData().OutputDir
		fmt.Println(outDir)
		fmt.Println(config)
	},
}
