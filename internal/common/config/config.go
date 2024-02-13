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
	Name string `env:"DB_NAME" default:"todo"`
	User string `env:"DB_USER" default:"citcho"`
	Pass string `env:"DB_PASS" default:"Secretp@ssw0rd"`
	Port int    `env:"DB_PORT" default:"3306"`
	Host string `env:"DB_HOST" default:"db"`
}

type Server struct {
	TodoHost string `env:"TODO_HOST"`
	TodoEnv  string `env:"TODO_ENV"`
	AppPort  int    `env:"TODO_PORT"`
}

func NewConfig() *Config {
	once.Do(func() {
		if err := env.Parse(&config); err != nil {
			panic(err)
		}
	})

	return &config
}
