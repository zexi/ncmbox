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
	cli    client.Client
	player player.Player

	songController SongController
}

func NewController(cli client.Client) (Controller, error) {
	player, err := player.NewPlayer()
	if err != nil {
		return nil, errors.Wrap(err, "init player")
	}
	c := &controller{
		cli:            cli,
		player:         player,
		songController: newSongController(cli, player),
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
}

type songController struct {
	cli    client.Client
	player player.Player
}

func newSongController(cli client.Client, player player.Player) SongController {
	return &songController{
		cli:    cli,
		player: player,
	}
}

func (c *songController) Play(song model.Song) error {
	urls, err := c.cli.GetSongsURL(context.TODO(), 320, song.GetId())
	if err != nil {
		return errors.Wrap(err, "get song url")
	}
	return c.player.PlaySong(&player.Song{
		Url: urls[0].GetURL(),
	})
}
