package mainui

import (
	"context"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"yunion.io/x/log"

	"github.com/zexi/ncmbox/pkg/controller"
	"github.com/zexi/ncmbox/pkg/ui"
	"github.com/zexi/ncmbox/pkg/ui/playlist"
	"github.com/zexi/ncmbox/pkg/ui/songlist"
	"github.com/zexi/ncmbox/pkg/version"
)

type mainUI struct {
	app *tview.Application

	ctrl controller.Controller

	playlistUI ui.Playlist
	songlistUI ui.SongList
}

func NewUI(ctrl controller.Controller) ui.MainUI {
	ui := &mainUI{
		app:  tview.NewApplication(),
		ctrl: ctrl,
	}
	if err := ui.ctrl.Login(context.TODO()); err != nil {
		log.Fatalf("login error: %v", err)
	}

	ui.init()
	return ui
}

func (ui *mainUI) GetController() controller.Controller {
	return ui.ctrl
}

func (ui *mainUI) getApp() *tview.Application {
	return ui.app
}

func (ui *mainUI) SetFocus(item tview.Primitive) {
	ui.getApp().SetFocus(item)
}

func (ui *mainUI) QueueEvent(es ...tcell.Event) {
	for _, e := range es {
		ui.getApp().QueueEvent(e)
	}
}

func (ui *mainUI) GetPlaylist() ui.Playlist {
	return ui.playlistUI
}

func (ui *mainUI) GetSongList() ui.SongList {
	return ui.songlistUI
}

func (ui *mainUI) init() {
	pages := tview.NewPages()

	// songs UI
	songsUI := songlist.NewSongList(ui)

	// playlists UI
	playlistUI := playlist.NewPlaylist(ui, songsUI)
	ui.playlistUI = playlistUI

	// Create the layout
	flex := tview.NewFlex().
		AddItem(playlistUI, 0, 1, true).
		AddItem(songsUI, 0, 1, false)
	pages.AddPage("*finder*", flex, true, true)
	ui.getApp().SetRoot(pages, true).SetFocus(pages)
}

func (ui *mainUI) Run() error {
	if err := ui.playlistUI.Refresh(context.TODO()); err != nil {
		return err
	}

	return ui.getApp().Run()
}

/*
 * func (ui *mainUI) LoginCellphonePage() tview.Primitive {
 *     form := tview.NewForm().
 *         AddInputField("手机号: ", "", 9, ValidateLoginPhone, nil).
 *         AddPasswordField("密码: ", "", 0, '*', nil)
 *     form.SetBorder(true).SetTitle("请输入登录信息")
 *     return form
 * }
 */

type VersionView struct {
	*tview.TextView
}

func NewVersionView() *VersionView {
	return &VersionView{
		TextView: tview.NewTextView().SetText(version.GetJsonString()),
	}
}
