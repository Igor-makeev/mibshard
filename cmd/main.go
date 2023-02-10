package main

import (
	"context"
	"fmt"
	"mibshard/configs"
	"mibshard/internal/handler"
	"mibshard/internal/repository"
	postgres "mibshard/internal/repository/db/postgres"
	"mibshard/internal/service"
	"mibshard/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg := configs.NewConfig()
	ctx := context.Background()
	psclient, err := postgres.NewPostgresClient(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	psstore := postgres.NewPostgresWalletKeeper(cfg, psclient)

	repository := repository.NewRepository(psstore, cfg)

	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	server := new(server.Server)
	serverErrChan := server.Run(cfg.ServerPort, handler)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:

		fmt.Println("main: got terminate signal. Shutting down...")

		if err := server.Shutdown(); err != nil {
			fmt.Printf("main: received an error while shutting down the server: %v", err)
		}

		if err := service.Close(ctx); err != nil {
			fmt.Printf("main: an error was received while closing the service: %v", err)
		}

	case <-serverErrChan:
		fmt.Println("main: got server err signal. Shutting down...")

		if err := service.Close(ctx); err != nil {
			fmt.Printf("main: an error was received while closing the service: %v", err)
		}
	}
}
