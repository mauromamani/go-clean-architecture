package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	SQLite SQLiteConfig
}

type ServerConfig struct {
	Port string
}

type SQLiteConfig struct {
	FileName string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("Config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	c := &Config{}

	err := v.Unmarshal(c)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return c, nil
}