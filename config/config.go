package config

import "github.com/spf13/viper"

type App struct {
	Name           string `json:"name"`
	HTTPPort       string `json:"http_port"`
	LogLevel       int    `json:"log_level"`
	LogTimeFormat  string `json:"time_format"`
	ContextTimeout int    `json:"context_timeout "`
	JWTSecret      string `json:"jwt_secret"`
}

type PostgreSQL struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type Config struct {
	App        App
	PostgreSQL PostgreSQL
}

func NewConfig() Config {
	return Config{
		App: App{
			Name:           viper.GetString("APP_NAME"),
			HTTPPort:       viper.GetString("HTTP_PORT"),
			LogLevel:       viper.GetInt("LOG_LEVEL"),
			LogTimeFormat:  viper.GetString("LOG_TIME_FORMAT"),
			ContextTimeout: viper.GetInt("CONTEXT_TIMEOUT"),
			JWTSecret:      viper.GetString("JWT_SECRET"),
		},
		PostgreSQL: PostgreSQL{
			Name:     viper.GetString("PSQL_DB_NAME"),
			Host:     viper.GetString("PSQL_DB_HOST"),
			User:     viper.GetString("PSQL_DB_USER"),
			Port:     viper.GetString("PSQL_DB_PORT"),
			Password: viper.GetString("PSQL_DB_PASS"),
		},
	}
}
