package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfig() {
	// Look for config in the working directory
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
