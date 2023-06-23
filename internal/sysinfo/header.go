package sysinfo

import (
	"fmt"
	"os"
	"os/user"
)

func collectHeader() (string, error) {
	user_, err := user.Current()
	if err != nil {
		return "", err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s@%s", user_.Username, hostname), nil
}
