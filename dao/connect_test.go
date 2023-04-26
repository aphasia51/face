package dao

import (
	"testing"
)

func TestConnectMySQl(t *testing.T) {
	got := ConnectMySQL()
	var want error

	if got != want {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
