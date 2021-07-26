package model

type Song interface {
	Base
}

type SongURL interface {
	GetURL() string
}
