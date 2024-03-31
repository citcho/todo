package config

import (
	"sync"

	"github.com/caarlos0/env/v10"
)

var (
	once   sync.Once
	config Config
)

type Config struct {
	Server
	DB
}

type DB struct {
	DBName string `env:"DB_NAME" default:"todo_db"`
	DBUser string `env:"DB_USER" default:"citcho"`
	DBPass string `env:"DB_PASS" default:"Secretp@ssw0rd"`
	DBPort int    `env:"DB_PORT" default:"3306"`
	DBHost string `env:"DB_HOST" default:"db"`
}

type Server struct {
	ClientHost string `env:"CLIENT_HOST"`
	ClientPort int    `env:"CLIENT_PORT"`
	AppPort    int    `env:"TODO_PORT"`
	AppEnv     string `env:"TODO_ENV" default:"dev"`
}

func NewConfig() *Config {
	once.Do(func() {
		if err := env.Parse(&config); err != nil {
			panic(err)
		}
	})

	return &config
}
