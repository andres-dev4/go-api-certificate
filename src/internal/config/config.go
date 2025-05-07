package config

import "github.com/spf13/viper"

type Config struct {
	Port string
	// Agrega otros campos de configuración
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	cfg := &Config{
		Port: viper.GetString("PORT"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg, nil
}