package model

type Playlist interface {
	Base
}

type PlaylistDetail interface {
	Base

	GetSongs() []Song
}
