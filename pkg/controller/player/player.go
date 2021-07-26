package player

import (
	"fmt"
	"io"
	"os/exec"

	"yunion.io/x/log"
)

type Song struct {
	Url string
}

type Player interface {
	PlaySong(*Song) error
}

type player struct {
	currentSong *Song
	// player is backend mpg123 process
	player *mpg123
}

func NewPlayer() (Player, error) {
	player := &player{
		player: newMPG123(),
	}
	if err := player.player.Start(); err != nil {
		return nil, err
	}
	return player, nil
}

func (p *player) PlaySong(song *Song) error {
	if err := p.player.Play(song.Url); err != nil {
		return err
	}
	return nil
}

type mpg123 struct {
	proc  *exec.Cmd
	stdin io.WriteCloser
}

func newMPG123() *mpg123 {
	obj := &mpg123{
		proc: exec.Command("mpg123", "-R"),
	}
	return obj
}

func (helper *mpg123) Start() error {
	proc := helper.proc
	// run mpg123 generic control interface
	// mpg123 will then read and execute commands from stdin
	procStdin, err := helper.proc.StdinPipe()
	if err != nil {
		return err
	}
	if err := proc.Start(); err != nil {
		return err
	}
	helper.stdin = procStdin
	return nil
}

func (helper *mpg123) Play(url string) error {
	cmd := fmt.Sprintf("L %s\n", url)
	log.Debugf("play %s", url)
	if _, err := helper.stdin.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
