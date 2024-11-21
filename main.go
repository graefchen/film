package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// name of config file (without extension)
	viper.SetConfigName("film.yaml")
	// REQUIRED if the config file does not have the extension in the name
	viper.SetConfigType("yaml")
	// look for the config in the current
	viper.AddConfigPath(".")

	// Find and read the config file
	err := viper.ReadInConfig()

	// Handle errors reading the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println("Hello, World!")
}
