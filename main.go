package main

import (
	"fmt"
	"os"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
	"github.com/spf13/cobra"
	"github.com/valyala/fasttemplate"
)

var (
	rootCmd = &cobra.Command{
		Use:   "film",
		Short: "film is a static photoblog generator",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Film",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("film ∞")
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Init...")
		},
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a picture",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Add picture...")
			ds, err := disk.New("./film", ".lock")
			if err != nil {
				fmt.Println(err)
			}
			db, err := hare.New(ds)
			if err != nil {
				fmt.Println(err)
			}

			db.Close()
		},
	}

	removeCommand = &cobra.Command{
		Use:   "remove",
		Short: "removes a picture",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	listCommand = &cobra.Command{
		Use:   "remove",
		Short: "removes a picture",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "build the picture blog",
		Run:   build,
	}
)

func build(cmd *cobra.Command, args []string) {
	fmt.Println("Build...")
	img, err := imgio.Open("input.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	resized := transform.Resize(img, 800, 800, transform.Linear)
	if err := imgio.Save("output.png", resized, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}

	template := "<h1>{{ string }}</h1>"
	t := fasttemplate.New(template, "{{", "}}")
	s := t.ExecuteString(map[string]interface{}{
		"string": "Hello, World!",
	})
	fmt.Println(s)
}

func init() {
	rootCmd.AddCommand(versionCmd, initCmd, addCmd, buildCmd, removeCommand, listCommand)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
