package songlist

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"yunion.io/x/log"

	"github.com/zexi/ncmbox/pkg/controller"
	"github.com/zexi/ncmbox/pkg/model"
	"github.com/zexi/ncmbox/pkg/ui"
)

type songList struct {
	*tview.List

	mainUI ui.MainUI

	ctrl controller.Controller

	playlist model.Playlist
	songs    []model.Song
}

func NewSongList(mainUI ui.MainUI) ui.SongList {
	songsUI := tview.NewList().ShowSecondaryText(false)
	songsUI.SetBorder(true).SetTitle("Songs")

	sl := &songList{
		List:   songsUI,
		mainUI: mainUI,
		ctrl:   mainUI.GetController(),
	}

	ui.SetDefaultShortcuts(sl)

	songsUI.SetDoneFunc(func() {
		mainUI.GetApp().SetFocus(mainUI.GetPlaylist())
	})

	sl.SetSelectedBackgroundColor(tcell.ColorYellowGreen)

	return sl
}

func (ui *songList) refresh() {
	ui.Clear()
	for idx, song := range ui.songs {
		key := fmt.Sprintf("[%d] %s", idx, song.GetName())
		ui.AddItem(key, "", 0, ui.onSelected(song))
	}
}

func (ui *songList) onSelected(song model.Song) func() {
	return func() {
		if err := ui.ctrl.GetSongController().Play(song); err != nil {
			log.Errorf("Play song %s error: %v", song.GetName(), err)
		}
	}
}

func (ui songList) GetCurrentPlaylist() model.Playlist {
	return ui.playlist
}

func (ui *songList) SetSongs(playlist model.Playlist, songs []model.Song) {
	ui.playlist = playlist
	ui.songs = songs
	ui.refresh()
}
