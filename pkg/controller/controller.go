package controller

import (
	"context"
	"strconv"

	"github.com/pkg/errors"

	"github.com/zexi/ncmbox/pkg/client"
	"github.com/zexi/ncmbox/pkg/controller/player"
	"github.com/zexi/ncmbox/pkg/model"
)

type Controller interface {
	Login(ctx context.Context) error
	ListUserPlaylist(ctx context.Context) ([]model.Playlist, error)
	GetUserPlaylist(ctx context.Context, id string) (model.PlaylistDetail, error)

	GetSongController() SongController
}

type controller struct {
	cli            client.Client
	songController SongController
}

func NewController(cli client.Client) (Controller, error) {
	songCtrl, err := newSongController(cli)
	if err != nil {
		return nil, errors.Wrap(err, "new song controller")
	}
	c := &controller{
		cli:            cli,
		songController: songCtrl,
	}
	return c, nil
}

func (c controller) Login(ctx context.Context) error {
	return c.cli.Login(ctx)
}

func (c controller) ListUserPlaylist(ctx context.Context) ([]model.Playlist, error) {
	return c.cli.ListUserPlaylist(ctx)
}

func (c controller) GetUserPlaylist(ctx context.Context, id string) (model.PlaylistDetail, error) {
	idNum, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return c.cli.GetUserPlaylist(ctx, idNum)
}

func (c controller) GetSongController() SongController {
	return c.songController
}

type SongController interface {
	Play(song model.Song) error

	SetFinishCallback(func())
	SetPauseCallback(func())
}

type songController struct {
	cli    client.Client
	player player.Player

	currentSong model.SongURL

	finishCallback func()
	pauseCallback  func()
}

func newSongController(cli client.Client) (SongController, error) {
	ctrl := &songController{
		cli: cli,
	}
	player, err := player.NewPlayer(ctrl.onFinish)
	if err != nil {
		return nil, errors.Wrap(err, "init player")
	}
	ctrl.player = player
	return ctrl, nil
}

func (c *songController) Play(song model.Song) error {
	urls, err := c.cli.GetSongsURL(context.TODO(), 320, song.GetId())
	if err != nil {
		return errors.Wrap(err, "get song url")
	}
	c.currentSong = urls[0]
	return c.playSongURL(c.currentSong)
}

func (c *songController) playSongURL(url model.SongURL) error {
	return c.player.PlaySong(&player.Song{
		Url: url.GetURL(),
	})
}

func (c *songController) SetFinishCallback(f func()) {
	c.finishCallback = f
}

func (c *songController) SetPauseCallback(f func()) {
	c.pauseCallback = f
}

func (c *songController) onFinish() {
	if c.finishCallback == nil {
		c.playSongURL(c.currentSong)
		return
	}
	c.finishCallback()
}

func (c *songController) onPause() {
	if c.pauseCallback == nil {
		return
	}
	c.pauseCallback()
}
