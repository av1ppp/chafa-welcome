package sysinfo

import (
	"os"
	"os/exec"
	"strings"
)

func readFile(name string) (string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func execute(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func findAndTrimLinePrefix(data string, prefix string) (string, bool) {
	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, prefix) {
			return strings.Trim(strings.TrimPrefix(line, prefix), "\""), true
		}
	}
	return "", false
}

func bToMB(b uint64) uint64 {
	return b / 1024 / 1024
}
