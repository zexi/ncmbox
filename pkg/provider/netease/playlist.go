package netease

import (
	"context"

	"github.com/pkg/errors"
	"github.com/winterssy/sreq"
)

func (a *API) ListUserPlaylist(ctx context.Context, userId string) (*PlaylistsResponse, error) {
	params := map[string]interface{}{
		"uid":    userId,
		"limit":  1000,
		"offset": 0,
	}
	resp := new(PlaylistsResponse)
	err := a.Request(sreq.MethodPost, WEAPI("/user/playlist"),
		sreq.WithForm(weapi(params)),
		sreq.WithContext(ctx),
	).JSON(resp)
	if err != nil {
		return nil, err
	}
	if err := CheckResponseError(resp.CommonResponse); err != nil {
		return nil, errors.Wrap(err, "get playlist")
	}

	return resp, nil
}

func (a *API) GetPlaylist(ctx context.Context, playlistId int) (*PlaylistDetailResponse, error) {
	params := map[string]interface{}{
		"id":     playlistId,
		"total":  true,
		"limit":  1000,
		"n":      100000,
		"offset": 0,
	}

	resp := new(PlaylistDetailResponse)
	err := a.Request(sreq.MethodPost, WEAPI("/v3/playlist/detail"),
		sreq.WithForm(weapi(params)),
		sreq.WithContext(ctx),
	).JSON(resp)
	if err != nil {
		return nil, err
	}
	if err := CheckResponseError(resp.CommonResponse); err != nil {
		return nil, errors.Wrap(err, "get playlist")
	}

	return resp, nil
}
