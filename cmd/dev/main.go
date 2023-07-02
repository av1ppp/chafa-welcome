package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/av1ppp/chafa-welcome/internal/config"
	"github.com/av1ppp/chafa-welcome/internal/global"
	"github.com/av1ppp/chafa-welcome/internal/sysinfo"
)

func main() {
	if err := innerMain(); err != nil {
		panic(err)
	}
}

func innerMain() error {
	gap := strings.Repeat(" ", 2)               // move to config file
	pictureMarginLeft := strings.Repeat(" ", 1) // move to config file

	homeDir := global.HomeDir()
	configPath := filepath.Join(homeDir, "config")
	config_, err := config.ParseFile(configPath)
	if err != nil {
		return err
	}

	fmt.Println("PicturePath:", config_.PicturePath)
	fmt.Println("ChafaBin:", config_.ChafaBin)
	fmt.Println("Width:", config_.Width)

	info, err := sysinfo.Collect()
	if err != nil {
		return err
	}
	infoLines := strings.Split(info.String(), "\n")
	infoNumberLines := len(infoLines)
	fmt.Println("infoNumberLines:", infoNumberLines)

	chafaOutput, err := chafaExecute(config_)
	if err != nil {
		return err
	}
	chafaLines := strings.Split(chafaOutput, "\n")
	chafaNumberLines := len(chafaLines) - 1
	chafaEmptyRow := strings.Repeat(" ", config_.Width)
	fmt.Println("chafaNumberLines:", chafaNumberLines)

	maxLines := 0
	if infoNumberLines > chafaNumberLines {
		maxLines = infoNumberLines
	} else {
		maxLines = chafaNumberLines
	}
	fmt.Println("maxLines:", maxLines)

	resultBuilder := strings.Builder{}

	for i := 0; i < maxLines; i++ {
		if i < chafaNumberLines {
			// with picture row
			if i < infoNumberLines {
				// with info row
				resultBuilder.WriteString(pictureMarginLeft + chafaLines[i] + gap + infoLines[i] + "\n")
			} else {
				// without info row
				resultBuilder.WriteString(pictureMarginLeft + chafaLines[i] + "\n")
			}
		} else {
			// without picture row
			resultBuilder.WriteString(pictureMarginLeft + chafaEmptyRow + gap + infoLines[i] + "\n")
		}
	}

	fmt.Println(resultBuilder.String())

	return nil
}

func chafaExecute(c *config.Config) (string, error) {
	args := []string{}

	if c.Width != 0 {
		args = append(args, fmt.Sprintf("--size=%d", c.Width))
	}

	args = append(args, c.PicturePath)

	cmd := exec.Command(c.ChafaBin, args...)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return stdout.String(), nil
}
