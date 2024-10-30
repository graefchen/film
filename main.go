package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/jameycribbs/hare"
	"github.com/jameycribbs/hare/datastores/disk"
	"github.com/spf13/cobra"
	"github.com/valyala/fasttemplate"
)

// The Commands

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
		Run:   initFilm,
	}

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a picture",
		Run:   addPicture,
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

// Picture Helpers

type PictureType struct {
	OTHER int
	JPEG  int
	PNG   int
}

var PType PictureType = PictureType{
	OTHER: 0,
	JPEG:  1,
	PNG:   2,
}

// DB Helpers

func openDB() *hare.Database {
	ds, err := disk.New("./film", ".json")
	if err != nil {
		fmt.Println("Could not finde the 'film.json' file.\nDid you use 'film init'?")
		os.Exit(1)
	}
	db, err := hare.New(ds)
	if err != nil {
		fmt.Println(err)
	}
	return db

}

// Start of commands

func initFilm(cmd *cobra.Command, args []string) {
	fmt.Println("Init...")
}

func addPicture(cmd *cobra.Command, args []string) {
	db := openDB()

	db.Close()
}

func buildPictureSite(picture string, title string, id int) {
	// opening the image and getting the bounds to use them later
	img, err := imgio.Open(picture)
	fmt.Println(img.Bounds().String())
	if err != nil {
		fmt.Println(err)
		return
	}

	// refactor into an own function that can *dynamically*
	// resize the picture to make it way more resiliant
	resized := transform.Resize(img, 400, 400, transform.Linear)
	if err := imgio.Save(strconv.Itoa(id)+".jpg", resized, imgio.JPEGEncoder(100)); err != nil {
		fmt.Println(err)
		return
	}

	template := "<html><h1>{{title}}</h1></html>"
	t := fasttemplate.New(template, "{{", "}}")
	s := t.ExecuteString(map[string]interface{}{
		"title": title,
	})
	fmt.Println(s)
}

func buildArchiveSite() {

}

func build(cmd *cobra.Command, args []string) {
	db := openDB()
	fmt.Println("Build...")
	buildPictureSite("picture.jpg", "", 1)
	buildArchiveSite()
	db.Close()
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(versionCmd, initCmd, addCmd, buildCmd, removeCommand, listCommand)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
