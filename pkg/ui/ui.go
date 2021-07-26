package ui

import (
	"context"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/zexi/ncmbox/pkg/controller"
	"github.com/zexi/ncmbox/pkg/model"
)

type MainUI interface {
	Run() error

	GetApp() *tview.Application
	GetController() controller.Controller

	GetPlaylist() Playlist
	GetSongList() SongList
}

type Playlist interface {
	tview.Primitive

	Refresh(ctx context.Context) error
}

type SongList interface {
	tview.Primitive

	SetSongs([]model.Song)
}

type View interface {
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *tview.Box
}

func SetDefaultShortcuts(view View) {
	view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'h' {
			return tcell.NewEventKey(tcell.KeyLeft, ' ', tcell.ModNone)
		}
		if event.Rune() == 'j' {
			return tcell.NewEventKey(tcell.KeyDown, ' ', tcell.ModNone)
		}
		if event.Rune() == 'k' {
			return tcell.NewEventKey(tcell.KeyUp, ' ', tcell.ModNone)
		}
		if event.Rune() == 'l' {
			return tcell.NewEventKey(tcell.KeyRight, ' ', tcell.ModNone)
		}
		return event
	})
}
