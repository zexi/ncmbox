package ui

import (
	"context"

	"github.com/rivo/tview"
	"yunion.io/x/log"

	"github.com/zexi/ncmbox/pkg/client"
	"github.com/zexi/ncmbox/pkg/config"
	"github.com/zexi/ncmbox/pkg/player"
	"github.com/zexi/ncmbox/pkg/provider/netease"
)

func ValidateLoginPhone(textToCheck string, lastChar rune) bool {
	log.Infof("Text to check: %q, lastChar: %q", textToCheck, lastChar)
	cfg := config.EnsureGetConfig()
	log.Infof("cfg: %v", cfg)
	return true
}

type Playlists struct {
	cli    *client.Client
	player *player.Player

	app           *tview.Application
	ui            *tview.List
	songsUI       *tview.List
	playlistItems []*playlistItem
}

func NewPlaylists(
	cli *client.Client,
	player *player.Player,
	app *tview.Application,
	ui *tview.List,
	songsUI *tview.List,
) *Playlists {
	controller := &Playlists{
		cli:    cli,
		player: player,
		// UI parts
		app:     app,
		ui:      ui,
		songsUI: songsUI,
	}
	controller.Refresh()
	return controller
}

func (p *Playlists) Refresh() {
	list, err := p.cli.ListUserPlaylist(context.TODO())
	if err != nil {
		log.Errorf("get playlist: %v", err)
	}
	p.ui.Clear()
	items := make([]*playlistItem, 0)
	for _, l := range list {
		tmp := l
		item := newPlaylistItem(p.cli, p.player, p.app, &tmp, p.songsUI)
		items = append(items, item)
		p.ui.AddItem(l.Name, "", 0, item.selected)
	}
	p.playlistItems = items
	// When the user navigates to a playlist, show its songs
	p.ui.SetChangedFunc(func(idx int, name string, secText string, shortcut rune) {
		p.playlistItems[idx].refresh()
	})
}

type playlistItem struct {
	cli     *client.Client
	player  *player.Player
	app     *tview.Application
	data    *netease.Playlist
	songsUI *tview.List
}

func newPlaylistItem(
	cli *client.Client,
	player *player.Player,
	app *tview.Application,
	data *netease.Playlist,
	songsUI *tview.List) *playlistItem {
	return &playlistItem{
		cli:     cli,
		player:  player,
		app:     app,
		data:    data,
		songsUI: songsUI,
	}
}

func (item *playlistItem) refresh() {
	NewSongsList(item.cli, item.player, item.data, item.songsUI)
}

func (item *playlistItem) selected() {
	item.refresh()
	item.app.SetFocus(item.songsUI)
}

type SongsList struct {
	cli          *client.Client
	player       *player.Player
	ui           *tview.List
	playlistData *netease.Playlist
}

func NewSongsList(cli *client.Client, player *player.Player, playlistData *netease.Playlist, ui *tview.List) *SongsList {
	controller := &SongsList{
		cli:          cli,
		player:       player,
		ui:           ui,
		playlistData: playlistData,
	}
	controller.Refresh()
	return controller
}

func (l *SongsList) Refresh() {
	details, err := l.cli.GetUserPlaylist(context.TODO(), l.playlistData.Id)
	if err != nil {
		log.Errorf("get playlist %v: %v", l.playlistData.Id, err)
		return
	}
	l.ui.Clear()
	for _, song := range details.Tracks {
		tmp := song
		l.ui.AddItem(song.Name, "", 0, NewSongManager(l.cli, l.player, tmp).Play)
	}
}

type SongManager struct {
	cli    *client.Client
	player *player.Player
	data   *netease.Song
}

func NewSongManager(cli *client.Client, player *player.Player, data *netease.Song) *SongManager {
	man := &SongManager{
		cli:    cli,
		player: player,
		data:   data,
	}
	return man
}

func (m *SongManager) Play() {
	urls, err := m.cli.GetSongsURL(context.TODO(), 320, m.data.Id)
	if err != nil {
		log.Errorf("get song url: %v", err)
		return
	}
	go func() {
		m.player.PlaySong(&player.Song{Url: urls[0].URL})
	}()
}
