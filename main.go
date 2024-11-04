package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/spf13/cobra"
	"github.com/valyala/fasttemplate"
)

// The Commands

var (
	rootCmd = &cobra.Command{
		Use:   "film",
		Short: "film is a static photoblog generator",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	initCmd = &cobra.Command{
		Use: "create",
		Run: create,
	}

	exportCmd = &cobra.Command{
		Use: "export",
		Run: export,
	}
)

// Picture Helpers

type PictureSize struct {
	width  int
	height int
}

func newPictureSize(width int, height int) *PictureSize {
	ps := PictureSize{width: width, height: height}
	return &ps
}

func resizePicture(img image.Image, sizes []PictureSize) ([]image.Image, error) {
	pictures := []image.Image{}
	for _, size := range sizes {
		resized := transform.Resize(img, size.width, size.height, transform.Linear)
		pictures = append(pictures, resized)

	}
	return pictures, nil
}

func savePictures(title string, img []image.Image, id int) {
	for _, resize := range img {
		dim := strconv.Itoa(resize.Bounds().Max.X) + "_" + strconv.Itoa(resize.Bounds().Max.Y)
		name := dim + strconv.Itoa(id) + title + ".jpg"
		if err := imgio.Save(name, resize, imgio.JPEGEncoder(100)); err != nil {
			fmt.Println(err)
			return
		}
	}
}

// Start of commands

func create(cmd *cobra.Command, args []string) {
	err := os.Mkdir("film", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func buildPictureSite(picture string, title string, id int) {
	// opening the image and getting the bounds to use them later
	img, err := imgio.Open(picture)
	if err != nil {
		fmt.Println(err)
		return
	}
	x := img.Bounds().Min.X
	y := img.Bounds().Min.X

	sizes := []PictureSize{
		*newPictureSize(x/2, y/2),
		*newPictureSize(x/4*3, y/4*3),
	}

	resized, err := resizePicture(img, sizes)
	if err != nil {
		fmt.Println(err)
		return
	}

	savePictures(strings.Split(picture, ".")[0], resized, id)

	template := "<html><body><h1>{{title}}</h1></body></html>"
	t := fasttemplate.New(template, "{{", "}}")
	s := t.ExecuteString(map[string]interface{}{
		"title": title,
	})
	fmt.Println(s)
}

func export(cmd *cobra.Command, args []string) {
	buildPictureSite("tine.jpg", "Tine", 1)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(initCmd, exportCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
