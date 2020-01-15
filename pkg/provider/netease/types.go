package netease

import (
	"strconv"
)

type CommonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func (c *CommonResponse) errorMessage() string {
	if c.Msg == "" {
		return strconv.Itoa(c.Code)
	}
	return c.Msg
}

type ObjectMeta struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Song struct {
	ObjectMeta
	Artists []Artist `json:"ar"`
	Album   Album    `json:"al"`
	Track   int      `json:"no"`
	Lyric   string   `json:"-"`
	URL     string   `json:"-"`
}

type Artist struct {
	ObjectMeta
	PicURL string `json:"picUrl"`
}

type Album struct {
	ObjectMeta
	PicURL string `json:"picUrl"`
}

type PlaylistCreator struct {
	Nickname string `json:"nickname"`
}

type Playlist struct {
	ObjectMeta
	PlaylistCreator PlaylistCreator `json:"creator"`
}

type PlaylistsResponse struct {
	CommonResponse
	Playlists []Playlist `json:"playlist"`
}

type PlaylistDetail struct {
	ObjectMeta
	CoverImgUrl string  `json:"coverImgUrl"`
	Tracks      []*Song `json:"tracks"`
	TrackIds    []struct {
		Id int `json:"id"`
	} `json:"trackIds"`
	TrackCount int `json:"trackCount"`
}

type PlaylistDetailResponse struct {
	CommonResponse
	PlaylistDetail PlaylistDetail `json:"playlist"`
}

// https://binaryify.github.io/NeteaseCloudMusicApi/#/?id=_1-%e6%89%8b%e6%9c%ba%e7%99%bb%e5%bd%95
type LoginResponse struct {
	CommonResponse
	LoginType int `json:"loginType"`
	Account   struct {
		Id       int    `json:"id"`
		UserName string `json:"userName"`
	} `json:"account"`
	Profile struct {
		Nickname string `json:"nickname"`
		UserType int    `json:"userType"`
		VIPType  int    `json:"vipType"`
	} `json:"profile"`
}

type SongsResponse struct {
	CommonResponse
	Songs []*Song `json:"songs"`
}

type SongURL struct {
	Code int    `json:"code"`
	Id   int    `json:"id"`
	BR   int    `json:"br"`
	URL  string `json:"url"`
}

type SongURLResponse struct {
	CommonResponse
	Data []SongURL `json:"data"`
}
