package sysinfo

import (
	"errors"
	"os"

	"github.com/av1ppp/chafa-welcome/internal/config"
)

func collectShell(conf *config.Config) (string, error) {
	shell, exists := os.LookupEnv("SHELL")
	if !exists {
		return "", errors.New("SHELL environment not set")
	}
	return shell, nil
}
