package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/winterssy/sreq"

	"github.com/zexi/ncmbox/pkg/model"
	"github.com/zexi/ncmbox/pkg/provider/netease"
)

type Client interface {
	Login(context.Context) error
	ListUserPlaylist(context.Context) ([]model.Playlist, error)
	GetUserPlaylist(ctx context.Context, id int) (model.PlaylistDetail, error)
	GetSongsURL(ctx context.Context, br int, songIds ...string) ([]model.SongURL, error)
}

type client struct {
	username  string
	cellphone int
	password  string
	api       *netease.API
	loginInfo *netease.LoginResponse
}

func newHttpClient() *sreq.Client {
	httpCli := sreq.New()
	defaultTransport := sreq.DefaultTransport()
	defaultTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
	httpCli.RawClient.Transport = defaultTransport
	return httpCli
}

func NewClient(username, password string) Client {
	cli := &client{
		username: username,
		password: password,
		api:      netease.New(newHttpClient()),
	}
	if number, err := strconv.Atoi(username); err == nil {
		cli.cellphone = number
	}
	return cli
}

func (c *client) UserId() string {
	return fmt.Sprintf("%d", c.loginInfo.Account.Id)
}

func (c *client) Login(ctx context.Context) error {
	resp, err := c.api.CellphoneLoginRaw(ctx, 86, c.cellphone, c.password)
	if err != nil {
		return err
	}
	c.loginInfo = resp
	return nil
}

func (c *client) ListUserPlaylist(ctx context.Context) ([]model.Playlist, error) {
	resp, err := c.api.ListUserPlaylist(ctx, c.UserId())
	if err != nil {
		return nil, err
	}

	ret := make([]model.Playlist, 0)
	for _, obj := range resp.Playlists {
		ret = append(ret, obj)
	}

	return ret, nil
}

func (c *client) GetUserPlaylist(ctx context.Context, id int) (model.PlaylistDetail, error) {
	resp, err := c.api.GetPlaylist(ctx, id)
	if err != nil {
		return nil, err
	}
	return &resp.PlaylistDetail, nil
}

func (c *client) GetSongsURL(ctx context.Context, br int, songIds ...string) ([]model.SongURL, error) {

	songIdNum := make([]int, 0)
	for _, id := range songIds {
		idNum, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		songIdNum = append(songIdNum, idNum)
	}

	resp, err := c.api.GetSongsURL(ctx, br, songIdNum...)
	if err != nil {
		return nil, err
	}

	ret := make([]model.SongURL, 0)
	for _, obj := range resp.Data {
		ret = append(ret, obj)
	}
	return ret, nil
}
