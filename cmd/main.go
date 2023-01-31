package main

import (
	"fmt"
	"mibshard/configs"
	"mibshard/internal/handler"
	"mibshard/internal/repository"
	redis "mibshard/internal/repository/db/redis"
	"mibshard/internal/service"
	"mibshard/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := configs.NewConfig()
	redisVault := redis.NewRedisVault(cfg.RedisAddr)
	repository := repository.NewRepository(redisVault, cfg)

	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	server := new(server.Server)
	serverErrChan := server.Run(cfg.ServerPort, handler)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:
		fmt.Println("main: got terminate signal. Shutting down...", nil)

	case <-serverErrChan:
		fmt.Println("main: got server err signal. Shutting down...", nil)
	}
}
