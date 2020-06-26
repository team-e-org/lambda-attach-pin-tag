package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     int
	TimeZone string
}

func ReadDBConfig() (*DBConfig, error) {
	dbPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("reading env var 'MYSQL_PORT': %w", err)
	}

	dbConfig := &DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     dbPort,
		TimeZone: os.Getenv("TZ"),
	}

	return dbConfig, nil
}

func ReadConfig() (*Config, error) {

	dbConfig, err := ReadDBConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		*dbConfig,
	}, nil
}
