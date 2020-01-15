package provider

import (
	"context"
)

type API interface {
	Login(ctx context.Context, username, password string)
	ListUserPlaylist(ctx context.Context, userId string, offset int, limit int)
}
