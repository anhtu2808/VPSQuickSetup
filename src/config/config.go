package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	IP       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config, path string) error {
	viper.Set("ip", config.IP)
	viper.Set("user", config.User)
	viper.Set("password", config.Password)

	return viper.WriteConfig()
}
