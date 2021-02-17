package main

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thanhhm/stw/service"
	"github.com/thanhhm/stw/util"
)

type config struct {
	Environment string
}

var cfg config

func init() {
	// Read config environment
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config error: %s", err.Error()))
	}

	// Load config
	loadConfig()

	// Init logger config
	initLogger(cfg.Environment)
}

func main() {
	hc := util.NewHTTPC(&util.HTTPCConfig{Timeout: 3 * time.Second})
	crawlerSvc := service.NewCrawler(hc)

	ctx := context.Background()
	tvData, err := crawlerSvc.GetTradingView(ctx)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("title: ", tvData.Title)
}

func loadConfig() {
	cfg.Environment = viper.GetString("environment")
}

func initLogger(envType string) {
	if cfg.Environment == "dev" {
		log.SetOutput(os.Stdout)
	} else {
		f, err := os.OpenFile(".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Errorf("Create log file error: %s", err.Error()))
		}
		log.SetOutput(f)
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.ErrorLevel)
	}
}
