package main

import (
	"log"
	"os"

	"github.com/zexi/ncmbox/pkg/app"
)

func main() {
	if err := app.NewApp().Start(); err != nil {
		log.Printf("Start app error: %v", err)
		os.Exit(1)
	}
}
