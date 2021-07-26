package main

import (
	"os"

	"yunion.io/x/log"

	"github.com/zexi/ncmbox/pkg/app"
)

func main() {
	log.SetLogLevelByString(log.Logger(), "info")

	app, err := app.NewApp()
	if err != nil {
		log.Errorf("New app error: %v", err)
		os.Exit(1)
	}
	if err := app.Start(); err != nil {
		log.Errorf("Start app error: %v", err)
		os.Exit(1)
	}
}
