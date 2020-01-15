package netease

import (
	"net/url"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/winterssy/sreq"
)

const (
	APIEndpoint = "https://music.163.com"
)

var (
	cookies sreq.Cookies

	defaultHeaders = sreq.Headers{
		"Origin":  APIEndpoint,
		"Referer": APIEndpoint,
	}
)

func init() {
	cookies, _ = createCookie()
}

type API struct {
	Client *sreq.Client
}

func New(client *sreq.Client) *API {
	if client == nil {
		client = sreq.New()
		// client.OnBeforeRequest(sreq.SetDefaultUserAgent())
	}
	return &API{
		Client: client,
	}
}

func (a *API) Request(method string, path string, opts ...sreq.RequestOption) *sreq.Response {
	opts = append(opts, sreq.WithHeaders(defaultHeaders))
	u, _ := url.Parse(APIEndpoint)
	u.Path = path
	url := u.String()
	// 如果已经登录，不需要额外设置cookies，cookie jar会自动管理
	_, err := a.Client.FilterCookie(url, "MUSIC_U")
	if err != nil {
		opts = append(opts, sreq.WithCookies(cookies))
	}

	return a.Client.Send(method, url, opts...)
}

func WEAPI(path ...string) string {
	weapi := "/weapi"
	paths := []string{weapi}
	paths = append(paths, path...)
	return filepath.Join(paths...)
}

func CheckResponseError(resp CommonResponse) error {
	if resp.Code != 200 {
		return errors.Errorf("code: %d, message: %s", resp.Code, resp.errorMessage())
	}
	return nil
}
