package app

import (
	"github.com/zexi/ncmbox/pkg/ui"
)

type App struct {
	ui *ui.UI
}

func NewApp() *App {
	app := &App{
		ui: ui.NewUI(),
	}
	return app
}

func (app *App) Start() error {
	return app.ui.Run()
}
