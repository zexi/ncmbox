package player

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"

	// "github.com/pkg/errors"
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

func NewPlayer(onFinish func()) (Player, error) {
	player := &player{
		player: newMPG123(onFinish),
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
	proc     *exec.Cmd
	stdin    io.WriteCloser
	stdout   io.ReadCloser
	onFinish func()
}

func newMPG123(onFinish func()) *mpg123 {
	obj := &mpg123{
		proc:     exec.Command("mpg123", "-R"),
		onFinish: onFinish,
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
	procStdout, err := helper.proc.StdoutPipe()
	if err != nil {
		return err
	}
	if err := proc.Start(); err != nil {
		return err
	}
	helper.stdin = procStdin
	helper.stdout = procStdout

	go func() {
		if err := helper.trackOutput(); err != nil {
			log.Errorf("trackOutput error: %v", err)
		}
		log.Errorf("==trackOutput should not exit")
	}()
	return nil
}

func (helper *mpg123) trackOutput() error {
	scanner := bufio.NewScanner(helper.stdout)
	for scanner.Scan() {
		strOut := scanner.Text()
		// log.Infof("--out: %q", strOut)
		if len(strOut) == 0 {
			continue
		}
		if len(strOut) < 2 {
			continue
		}
		if strings.HasPrefix(strOut, "@F") {

		} else if strings.HasPrefix(strOut, "@P 0") {
			// at end
			if strOut == "@P 0" {
				helper.onFinish()
			}
		}
		/*
		 * prefix := strOut[0:2]
		 * switch prefix {
		 * case "@F":
		 *     // playing, update progress
		 *     continue
		 * case "@E":
		 *     continue
		 * case "@P 0":
		 *     // at end
		 *     helper.onFinish()
		 * }
		 */
	}
	return nil
}

func (helper *mpg123) Play(url string) error {
	cmd := fmt.Sprintf("LOAD %s\n", url)
	log.Debugf("play %s", url)
	if _, err := helper.stdin.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
