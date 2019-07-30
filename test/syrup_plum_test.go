package test

import (
	"syrup-plum"
	"testing"
)

func Test_NewSyrupPlum(t *testing.T) {
	option := syrup_plum.Option{
		SavePath: "../example/db/",
	}
	sp := syrup_plum.NewSyrupPlum(&option)
	if sp == nil {
		t.Fatal("sp is nil")
	}
}
