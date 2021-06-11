package main

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/thanhhm/stw/service"
	"github.com/thanhhm/stw/util"
)

var cfg config

func init() {
	// Load config
	loadConfig()

	// Init logger config
	initLogger(cfg.environment)
}

func main() {
	hc := util.NewHTTPC(&util.HTTPCConfig{Timeout: 3 * time.Second})
	crawlerSvc := service.NewCrawler(hc)

	ctx := context.Background()
	tvData, err := crawlerSvc.GetTradingView(ctx)
	if err != nil {
		log.Errorln("Get trading view title error: ", err.Error())
	}
	fmt.Println("title: ", tvData.Title)
}
