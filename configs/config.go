package configs

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	*PostgresConfig
	ServerPort string
}

type PostgresConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	Name     string
	SslMode  string
}

func NewPostgressConfig() *PostgresConfig {
	return &PostgresConfig{
		UserName: os.Getenv("DB_USER_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSLMODE"),
	}
}

func NewConfig() *Config {
	var cfg Config
	cfg.PostgresConfig = NewPostgressConfig()
	flag.StringVar(&cfg.ServerPort, "sp", "8080", "shard server port")

	flag.Parse()
	logrus.Printf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.UserName, cfg.Name, cfg.Password, cfg.SslMode)
	logrus.Printf("config: ServerPort=%v", cfg.ServerPort)

	return &cfg
}
