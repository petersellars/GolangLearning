package config

import "github.com/spf13/viper"

type Constants struct {
	PORT  string
	Mongo struct {
		URL    string
		DBName string
	}
}

func initViper() (Constants, error) {
	viper.SetConfigName("todo.config") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")           // Search the root directory for the configuration file
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		return Constants{}, err
	}
	viper.SetDefault("PORT", "8080")

	var constants Constants
	err = viper.Unmarshal(&constants)
	return constants, err
}
