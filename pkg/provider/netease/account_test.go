package netease

import (
	"context"
	"testing"
)

func TestAPI_CellphoneLoginRaw(t *testing.T) {
	cli := New(nil)
	resp, err := cli.CellphoneLoginRaw(context.TODO(), 86, 12345678910, "test")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", resp)
}
