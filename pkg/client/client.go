package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/winterssy/sreq"

	"github.com/zexi/ncmbox/pkg/provider/netease"
)

type Client struct {
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

func NewClient(username, password string) *Client {
	cli := &Client{
		username: username,
		password: password,
		api:      netease.New(newHttpClient()),
	}
	if number, err := strconv.Atoi(username); err == nil {
		cli.cellphone = number
	}
	return cli
}

func (c *Client) UserId() string {
	return fmt.Sprintf("%d", c.loginInfo.Account.Id)
}

func (c *Client) Login(ctx context.Context) error {
	resp, err := c.api.CellphoneLoginRaw(ctx, 86, c.cellphone, c.password)
	if err != nil {
		return err
	}
	c.loginInfo = resp
	return nil
}

func (c *Client) ListUserPlaylist(ctx context.Context) ([]netease.Playlist, error) {
	resp, err := c.api.ListUserPlaylist(ctx, c.UserId())
	if err != nil {
		return nil, err
	}
	return resp.Playlists, nil
}

func (c *Client) GetUserPlaylist(ctx context.Context, id int) (*netease.PlaylistDetail, error) {
	resp, err := c.api.GetPlaylist(ctx, id)
	if err != nil {
		return nil, err
	}
	return &resp.PlaylistDetail, nil
}

func (c *Client) GetSongsURL(ctx context.Context, br int, songIds ...int) ([]netease.SongURL, error) {
	resp, err := c.api.GetSongsURL(ctx, br, songIds...)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
