package config

import (
	"time"
)

type Config struct {
	Port           string
	SecretKey      string
	TokenDuration  time.Duration
	LogLevel       uint
	DatabaseURI    string
	DbMaxOpenConns int
	DbMaxIdleConns int
}

func (c *Config) DbUri() string {
	return c.DatabaseURI
}

func (c *Config) DbOpenConns() int {
	return c.DbMaxOpenConns
}

func (c *Config) DbIdleConns() int {
	return c.DbMaxIdleConns
}
