package sysinfo

import (
	"os"
	"os/user"
)

func collectUsername() (string, error) {
	user_, err := user.Current()
	if err != nil {
		return "", err
	}
	return user_.Username, nil
}

func collectHostname() (string, error) {
	return os.Hostname()
}
