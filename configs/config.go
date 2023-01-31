package configs

type Config struct {
	RedisAddr  string
	ServerPort string
}

func NewConfig() *Config {
	return &Config{
		RedisAddr:  "localhost:6379",
		ServerPort: "8080",
	}
}
