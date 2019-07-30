package test

import (
	"syrup-plum"
	"testing"
)

func Test_InitOptionWithConfigFile(t *testing.T) {
	opt, err := syrup_plum.InitOptionWithConfigFile("../example/config.ini")
	if err != nil {
		t.Fatal(err)
	}

	if opt.SavePath == "" {
		t.Fatal("save_path error")
	}

}

func Test_InitOption(t *testing.T) {
	opt, err := syrup_plum.InitOption("../example/db/")
	if err != nil {
		t.Fatal(err)
	}

	if opt.SavePath == "" {
		t.Fatal("save_path error")
	}
}
