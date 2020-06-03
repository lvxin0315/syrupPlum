package syrupPlum

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, errors.New(fmt.Sprintf("%s is not exist", path))
	}
	return false, err
}

func SetDebug(debug bool) {
	Debug = debug
}

func SPError(a ...interface{}) {
	if Debug {
		fmt.Println(a)
	}
}
