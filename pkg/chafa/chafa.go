package chafa

import (
	"bytes"
	"os/exec"
	"strconv"
)

func WithSize(width, height int) string {
	if height == 0 {
		return "--size=" + strconv.Itoa(width)
	}
	return "--size=" + strconv.Itoa(width) + "x" + strconv.Itoa(height)
}

func WithSymbols(symbols string) string {
	return "--symbols=" + symbols
}

func Execute(bin string, image string, args ...string) (string, error) {
	args = append(args, image)
	cmd := exec.Command(bin, args...)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return stdout.String(), nil
}
