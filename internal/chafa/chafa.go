package chafa

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/av1ppp/chafa-welcome/internal/config"
	"github.com/av1ppp/chafa-welcome/internal/global"

	"github.com/av1ppp/chafa-welcome/pkg/chafa"
)

func Execute(conf *config.Config) (string, error) {
	hash := asHash(conf)
	cacheName := filepath.Join(global.HomeDir(), "cache", hash)

	err := os.MkdirAll(filepath.Dir(cacheName), global.ModeDir)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(cacheName)

	if err == nil {
		data, err := os.ReadFile(cacheName)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}

	if os.IsNotExist(err) {
		data, err := chafa.Execute(conf.ChafaBin, conf.Image.Source,
			chafa.WithSize(conf.Image.Size, 0),
			chafa.WithSymbols(conf.Image.Symbols))
		if err != nil {
			return "", err
		}

		err = os.WriteFile(cacheName, []byte(data), global.ModeFile)
		if err != nil {
			return "", err
		}

		return data, nil
	}

	return "", err
}

func asHash(conf *config.Config) string {
	data := strconv.Itoa(conf.Image.Size) + ";" + conf.Image.Symbols + ";" + conf.Image.Source
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
