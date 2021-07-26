package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func readConfig() {
	// Look for config in the working directory
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func loadConfig() {
	loadDBConfig()
}

func Configure() {
	readConfig()
	loadConfig()
}
