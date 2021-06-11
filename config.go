package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	environment string
}

func loadConfig() {
	// Read config environment
	newConfig()

	// Get env variable from config file
	cfg.environment = viper.GetString("environment")
}

func newConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config error: %s", err.Error()))
	}

	// Init default value
	viper.SetDefault("environment", "dev")
}
