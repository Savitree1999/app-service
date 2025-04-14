package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Env  string
		Port string
	}
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	viper.AddConfigPath("config") // มองหาใน ./config
	viper.SetConfigName(fmt.Sprintf("config.%s", env)) // เช่น config.dev
	viper.SetConfigType("yaml")

	// Default
	viper.SetDefault("App.Env", env)
	viper.SetDefault("App.Port", "8080")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error loading config file %s: %w", viper.ConfigFileUsed(), err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}