package configs

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	DBAddress  string
	ServerPort string
}

func NewConfig() *Config {
	var cfg Config
	cfg.DBAddress = os.Getenv("DB_ADDRESS")
	flag.StringVar(&cfg.ServerPort, "sp", "8081", "shard server port")

	flag.Parse()
	logrus.Printf("config: dbAddress=%v", cfg.DBAddress)
	logrus.Printf("config: ServerPort=%v", cfg.ServerPort)

	return &cfg
}
