package test

import (
	"fmt"

	"github.com/spf13/viper"
)

func TesViper() *viper.Viper {
	config := viper.New()

	config.SetConfigFile("../.env")
	// config.AddConfigPath("../")

	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error env file: %w", err))
	}

	return config
}
