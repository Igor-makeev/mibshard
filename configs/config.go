package configs

import (
	"flag"

	"github.com/sirupsen/logrus"
)

type Config struct {
	PostgresAddres string
	ServerPort     string
}

func NewConfig() *Config {
	var cfg Config

	flag.StringVar(&cfg.PostgresAddres, "pa", "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable", "server addres to listen on")
	flag.StringVar(&cfg.ServerPort, "sp", "8080", "shortener base URL")

	flag.Parse()

	logrus.Printf("config: RedisAddr=%v", cfg.PostgresAddres)
	logrus.Printf("config: ServerPort=%v", cfg.ServerPort)

	return &cfg
}
