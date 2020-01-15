package netease

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/winterssy/sreq"
)

// 批量获取歌曲播放地址，br: 比特率，128/192/320/999
func (a *API) GetSongsURL(ctx context.Context, br int, songIds ...int) (*SongURLResponse, error) {
	var tmpBr int
	switch br {
	case 128, 192, 320:
		tmpBr = br
	default:
		tmpBr = 999
	}
	enc, _ := json.Marshal(songIds)
	data := map[string]interface{}{
		"br":  tmpBr * 1000,
		"ids": string(enc),
	}

	resp := new(SongURLResponse)
	err := a.Request(sreq.MethodPost, WEAPI("/song/enhance/player/url"),
		sreq.WithForm(weapi(data)),
		sreq.WithContext(ctx)).JSON(resp)
	if err != nil {
		return nil, err
	}
	if err := CheckResponseError(resp.CommonResponse); err != nil {
		return nil, errors.Wrap(err, "get songs url")
	}
	return resp, nil
}
