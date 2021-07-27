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

	playlist    model.Playlist
	songs       []model.Song
	currentSong model.Song
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
		mainUI.SetFocus(mainUI.GetPlaylist())
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
		songCtrl := ui.ctrl.GetSongController()
		songCtrl.SetFinishCallback(ui.onFinish)

		ui.currentSong = song
		if err := songCtrl.Play(song); err != nil {
			log.Errorf("Play song %s error: %v", song.GetName(), err)
		}
	}
}

func (lui *songList) onFinish() {
	log.Debugf("====on finish called: %d", lui.GetCurrentItem()+1)
	// ui.SetCurrentItem(ui.GetCurrentItem() + 1)
	lui.mainUI.QueueEvent(
		ui.NewEventKeyDown(),
		ui.NewEventKeyEnter())
}

func (ui songList) GetCurrentPlaylist() model.Playlist {
	return ui.playlist
}

func (ui *songList) SetSongs(playlist model.Playlist, songs []model.Song) {
	ui.playlist = playlist
	ui.songs = songs
	ui.refresh()
}
