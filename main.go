package main

import (
	"context"
	"fmt"
	"time"

	"github.com/thanhhm/stw/service"
	"github.com/thanhhm/stw/util"
)

func main() {
	hc := util.NewHTTPC(&util.HTTPCConfig{Timeout: 3 * time.Second})
	crawlerSvc := service.NewCrawler(hc)

	ctx := context.Background()
	_, err := crawlerSvc.GetTradingView(ctx)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
}
