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

func Execute(bin string, picture string, args ...string) (string, error) {
	args = append(args, picture)

	cmd := exec.Command(bin, args...)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return stdout.String(), nil
}
