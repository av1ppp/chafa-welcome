package global

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	homeDir     = ""
	homeDirOnce = &sync.Once{}
)

const (
	ModeFile = 0664
	ModeDir  = os.ModePerm
)

func HomeDir() string {
	homeDirOnce.Do(func() {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		homeDir = filepath.Join(userHomeDir, ".chafa-welcome")

		err = os.MkdirAll(homeDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	})

	return homeDir
}
