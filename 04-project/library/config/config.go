package config

import "os"

type WebServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	ConnectionString string
}

type Config struct {
	WebServerConfig *WebServerConfig
	DatabaseConfig  *DatabaseConfig
}

func GetConfig() *Config {
	return &Config{
		WebServerConfig: &WebServerConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},

		DatabaseConfig: &DatabaseConfig{
			ConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		},
	}
}
