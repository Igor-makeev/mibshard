package configs

import (
	"flag"

	"github.com/sirupsen/logrus"
)

type Config struct {
	RedisAddr  string
	ServerPort string
}

func NewConfig() *Config {
	var cfg Config

	flag.StringVar(&cfg.RedisAddr, "ra", "localhost:6379", "server addres to listen on")
	flag.StringVar(&cfg.ServerPort, "sp", "8080", "shortener base URL")

	flag.Parse()

	logrus.Printf("config: RedisAddr=%v", cfg.RedisAddr)
	logrus.Printf("config: ServerPort=%v", cfg.ServerPort)

	return &cfg
}
