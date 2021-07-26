package playlist

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rivo/tview"
	"yunion.io/x/log"

	"github.com/zexi/ncmbox/pkg/controller"
	"github.com/zexi/ncmbox/pkg/model"
	"github.com/zexi/ncmbox/pkg/ui"
)

type playlist struct {
	*tview.List

	mainUI  ui.MainUI
	ctrl    controller.Controller
	songsUI ui.SongList
}

func NewPlaylist(mainUI ui.MainUI, songsUI ui.SongList) ui.Playlist {
	plUI := tview.NewList().ShowSecondaryText(false)
	plUI.SetBorder(true).SetTitle("Playlists")

	pl := &playlist{
		List: plUI,

		mainUI:  mainUI,
		ctrl:    mainUI.GetController(),
		songsUI: songsUI,
	}

	ui.SetDefaultShortcuts(pl)

	return pl
}

func (p *playlist) Refresh(ctx context.Context) error {
	list, err := p.ctrl.ListUserPlaylist(ctx)
	if err != nil {
		return errors.Wrap(err, "get playlist")
	}

	p.Clear()

	for _, obj := range list {
		p.AddItem(obj.GetName(), "", 0, p.onSelected(ctx, obj))
	}

	return nil
}

func (p *playlist) onSelected(ctx context.Context, obj model.Playlist) func() {
	return func() {
		details, err := p.ctrl.GetUserPlaylist(ctx, obj.GetId())
		if err != nil {
			log.Errorf("Get user playlist %v: %v", obj.GetName(), err)
			return
		}

		p.songsUI.SetSongs(details.GetSongs())
		p.mainUI.GetApp().SetFocus(p.songsUI)
	}
}
