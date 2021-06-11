package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func initLogger(envType string) {
	if cfg.environment == "dev" {
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
