package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type fileConfig struct {
	v        *viper.Viper
	siteDate SiteData
	photos   []Photo
}

func NewFileConfig(file string) Config {
	v := viper.New()
	// name of config file
	v.SetConfigFile(file)
	// look for the config in the current directory
	v.AddConfigPath(".")
	// Find and read the config file
	err := v.ReadInConfig()
	// Handle errors reading the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := fileConfig{v: v}

	_ = v.UnmarshalKey("site", &config.siteDate)
	_ = v.UnmarshalKey("photos", &config.photos)

	return config
}

func (f fileConfig) GetSiteData() SiteData {
	return f.siteDate
}

func (f fileConfig) GetPhotos() []Photo {
	return f.photos
}
