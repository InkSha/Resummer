package config

import (
	"fmt"
	"strings"
)

type DatabaseConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Name     string `json:"name"`
	TimeZone string `json:"timezone"`
}

type GenDSNError struct {
	msg string
}

func (e *GenDSNError) Error() string {
	return e.msg
}

var (
	databaseConfig *DatabaseConfig
)

func GetDatabaseDSN() (string, error) {
	databaseConfig = GetDatabaseConfig()

	if strings.EqualFold(databaseConfig.Type, "postgres") {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			databaseConfig.Host, databaseConfig.Port, databaseConfig.User, databaseConfig.Pass, databaseConfig.Name, databaseConfig.TimeZone), nil
	}

	return "", &GenDSNError{msg: fmt.Sprintf("[%s] unsupported database type", databaseConfig.Type)}
}

func GetDatabaseConfig() *DatabaseConfig {
	if databaseConfig == nil {
		databaseConfig = &DatabaseConfig{
			Type:     Get("DATABASE_TYPE"),
			Host:     Get("DATABASE_HOST"),
			Port:     Get("DATABASE_PORT"),
			User:     Get("DATABASE_USER"),
			Pass:     Get("DATABASE_PASS"),
			Name:     Get("DATABASE_NAME"),
			TimeZone: Get("TIMEZONE"),
		}
	}

	return databaseConfig
}
