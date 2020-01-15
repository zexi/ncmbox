package ui

import (
	"context"

	"github.com/rivo/tview"
	"yunion.io/x/log"

	"github.com/zexi/ncmbox/pkg/client"
	"github.com/zexi/ncmbox/pkg/config"
	"github.com/zexi/ncmbox/pkg/player"
	"github.com/zexi/ncmbox/pkg/version"
)

type UI struct {
	*tview.Application
}

func NewUI() *UI {
	ui := &UI{
		Application: tview.NewApplication(),
	}
	cfg := config.EnsureGetConfig()
	log.Infof("Config: %v", cfg)
	cli := client.NewClient(cfg.Username, cfg.Password)
	if err := cli.Login(context.TODO()); err != nil {
		log.Fatalf("login error: %v", err)
	}
	player, err := player.NewPlayer()
	if err != nil {
		log.Fatalf("init player error: %v", err)
	}
	ui.init(cli, player)
	return ui
}

func (ui *UI) init(cli *client.Client, player *player.Player) {
	pages := tview.NewPages()

	// playlists UI
	playlistUI := tview.NewList().ShowSecondaryText(false)
	playlistUI.SetBorder(true).SetTitle("Playlists")

	// songs UI
	songsUI := tview.NewList().ShowSecondaryText(false)
	songsUI.SetBorder(true).SetTitle("Songs")

	NewPlaylists(cli, player, ui.Application, playlistUI, songsUI)

	// Create the layout
	flex := tview.NewFlex().
		AddItem(playlistUI, 0, 1, true).
		AddItem(songsUI, 0, 1, false)
	pages.AddPage("*finder*", flex, true, true)
	ui.SetRoot(pages, true).SetFocus(pages)
}

func (ui *UI) Run() error {
	return ui.Application.Run()
}

func (ui *UI) LoginCellphonePage() tview.Primitive {
	form := tview.NewForm().
		AddInputField("手机号: ", "", 9, ValidateLoginPhone, nil).
		AddPasswordField("密码: ", "", 0, '*', nil)
	form.SetBorder(true).SetTitle("请输入登录信息")
	return form
}

type VersionView struct {
	*tview.TextView
}

func NewVersionView() *VersionView {
	return &VersionView{
		TextView: tview.NewTextView().SetText(version.GetJsonString()),
	}
}
