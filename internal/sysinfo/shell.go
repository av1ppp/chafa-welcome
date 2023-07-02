package sysinfo

import (
	"os"
)

func collectShell() (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "-", nil
	}
	return shell, nil
}
