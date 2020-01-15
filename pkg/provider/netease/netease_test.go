package netease

import (
	"context"
	"encoding/json"
	"testing"
)

var (
	client *API
	ctx    context.Context
)

func setup() {
	client = New(nil)
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
}

func TestAPI_GetPlaylist(t *testing.T) {
	playlist, err := client.GetPlaylist(ctx, "156934569")
	if err != nil {
		t.Fatal(err)
	}
	str, _ := json.MarshalIndent(playlist, "", "  ")
	t.Logf("%s", string(str))
}
