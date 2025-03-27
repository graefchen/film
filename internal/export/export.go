package export

import (
	"film/internal/config"
	"fmt"
)

func Export() {
	// config
	config := config.FilmConfig()

	// reading the config
	outDir := config.GetSiteData().OutputDir
	fmt.Println(outDir)
	fmt.Println(config)
}
