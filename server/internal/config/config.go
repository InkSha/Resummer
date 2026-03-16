package config

import "github.com/joho/godotenv"

type Config struct {
	config map[string]string
}

var config *Config = &Config{}

func LoadConfig() {
	rawConfig, err := godotenv.Read()

	if err != nil {
		panic(err)
	}

	config.config = rawConfig
}

func Get(key string) string {
	return config.config[key]
}
