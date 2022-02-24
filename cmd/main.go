package main

import (
	"Diplom/pkg/processingData"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	web "Diplom"
	"Diplom/pkg/handler"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//var resultSMS []model.SMSData

	resultSMS := processingData.ResultSMS()
	fmt.Println("\n", resultSMS, "\n")

	resultVoiceCall := processingData.ResultVoiceCall()
	fmt.Println("voice", resultVoiceCall)

	srv := new(web.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
			logrus.Fatalf(err.Error())
		}
	}()

	logrus.Print("Diplom Project Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("HTTP31 Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
