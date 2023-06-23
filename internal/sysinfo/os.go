package sysinfo

import (
	"runtime"
)

func collectOS() (string, error) {
	data, err := execute("sh", "-c", "cat /etc/*-release")
	if err != nil {
		return "", err
	}

	prettyName, found := findAndTrimLinePrefix(data, "PRETTY_NAME=")
	if found {
		return prettyName + " " + runtime.GOOS, nil
	}

	name, found := findAndTrimLinePrefix(data, "NAME=")
	if found {
		return name + " " + runtime.GOOS, nil
	}

	return runtime.GOOS, nil
}
