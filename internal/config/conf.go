package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type Config struct {
	LogLevel string
	Address  string
	Port     string
}

const (
	defaultServer     = "localhost:"
	defaultServerPort = "8080"
)

func New(logLevel string) (*Config, error) {
	cfg := Config{
		LogLevel: logLevel,
	}

	flag.StringVar(&cfg.Address, "a", defaultServer, "server address [localhost]")
	flag.StringVar(&cfg.Port, "p", defaultServerPort, "server port [:8080]")

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
