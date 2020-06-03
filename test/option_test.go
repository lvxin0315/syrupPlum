package test

import (
	"github.com/lvxin0315/syrupPlum"
	"testing"
)

func Test_InitOptionWithConfigFile(t *testing.T) {
	opt, err := syrupPlum.InitOptionWithConfigFile("../example/config.ini")
	if err != nil {
		t.Fatal(err)
	}

	if opt.SavePath == "" {
		t.Fatal("save_path error")
	}

}

func Test_InitOption(t *testing.T) {
	opt, err := syrupPlum.InitOption("../example/db/")
	if err != nil {
		t.Fatal(err)
	}

	if opt.SavePath == "" {
		t.Fatal("save_path error")
	}
}
