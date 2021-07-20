package netease

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/pkg/errors"
	"github.com/winterssy/sreq"
)

// 邮箱登录
func (a *API) EmailLoginRaw(ctx context.Context, email string, password string) (*LoginResponse, error) {
	passwordHash := md5.Sum([]byte(password))
	password = hex.EncodeToString(passwordHash[:])
	data := map[string]interface{}{
		"username":      email,
		"password":      password,
		"rememberLogin": true,
	}

	resp := new(LoginResponse)
	err := a.Request(sreq.MethodPost, WEAPI("/login"),
		sreq.WithForm(weapi(data)),
		sreq.WithContext(ctx),
	).JSON(resp)
	if err != nil {
		return nil, err
	}

	if err := CheckResponseError(resp.CommonResponse); err != nil {
		return nil, errors.Wrap(err, "email login")
	}
	return resp, nil
}

// 手机登录
func (a *API) CellphoneLoginRaw(ctx context.Context, countryCode int, phone int, password string) (*LoginResponse, error) {
	passwordHash := md5.Sum([]byte(password))
	password = hex.EncodeToString(passwordHash[:])
	data := map[string]interface{}{
		"phone":         phone,
		"countrycode":   countryCode,
		"password":      password,
		"rememberLogin": true,
	}

	resp := new(LoginResponse)
	err := a.Request(sreq.MethodPost, WEAPI("/login/cellphone"),
		sreq.WithForm(weapi(data)),
		sreq.WithContext(ctx),
		sreq.WithCookies(map[string]string{
			"os": "pc",
		}),
	).JSON(resp)
	if err != nil {
		return nil, err
	}

	if err := CheckResponseError(resp.CommonResponse); err != nil {
		return nil, errors.Wrap(err, "cellphone login")
	}
	return resp, nil
}

// 刷新登录状态
func (a *API) RefreshLoginRaw(ctx context.Context) (*CommonResponse, error) {
	resp := new(CommonResponse)
	err := a.Request(sreq.MethodPost, WEAPI("/login/token/refresh"),
		sreq.WithForm(weapi(struct{}{})),
		sreq.WithContext(ctx),
	).JSON(resp)
	if err != nil {
		return nil, err
	}
	if err := CheckResponseError(*resp); err != nil {
		return nil, errors.Wrap(err, "refresh login")
	}
	return resp, nil
}

// 退出登录
func (a *API) LogoutRaw(ctx context.Context) (*CommonResponse, error) {
	resp := new(CommonResponse)
	err := a.Request(sreq.MethodPost, WEAPI("/logout"),
		sreq.WithForm(weapi(struct{}{})),
		sreq.WithContext(ctx),
	).JSON(resp)
	if err != nil {
		return nil, err
	}
	if err := CheckResponseError(*resp); err != nil {
		return nil, errors.Wrap(err, "logout")
	}
	return resp, nil
}
