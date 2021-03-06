package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"Diplom/internal/handler"
	"Diplom/internal/server"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	srv := new(server.Server)
	go func() {
		if err := srv.Run(handler.InitRoutes()); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()
	logrus.Print("Diplom Project Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Diplom Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
