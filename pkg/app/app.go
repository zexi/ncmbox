package app

import (
	"yunion.io/x/log"

	"github.com/zexi/ncmbox/pkg/client"
	"github.com/zexi/ncmbox/pkg/config"
	"github.com/zexi/ncmbox/pkg/controller"
	"github.com/zexi/ncmbox/pkg/ui"
	"github.com/zexi/ncmbox/pkg/ui/mainui"
)

type App struct {
	ui ui.MainUI
}

func NewApp() (*App, error) {
	cfg := config.EnsureGetConfig()
	log.Infof("Config: %v", cfg)

	cli := client.NewClient(cfg.GetUsername(), cfg.GetPassword())
	controller, err := controller.NewController(cli)
	if err != nil {
		return nil, err
	}

	app := &App{
		ui: mainui.NewUI(controller),
	}
	return app, nil
}

func (app *App) Start() error {
	return app.ui.Run()
}
