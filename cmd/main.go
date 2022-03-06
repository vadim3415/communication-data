package main

import (
	"Diplom/internal/processingData"
	web "Diplom/internal/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"Diplom/internal/handler"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//resultSMS := processingData.ResultSMS()
	//fmt.Println("\n SMS", resultSMS, "\n")

	resultVoiceCall := processingData.ResultVoiceCall()
	fmt.Println("\n VOICE", resultVoiceCall, "\n")

	resultEmail := processingData.ResultEmail()
	fmt.Println("\n Email", resultEmail, "\n")

	resultBilling := processingData.ResultBilling()
	fmt.Println("\n Billing", resultBilling, "\n")

	processingData.GetResultData()

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

	logrus.Print("Diplom Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
