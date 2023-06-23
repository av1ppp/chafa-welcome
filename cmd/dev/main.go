package main

import (
	"bytes"
	"fmt"
	"github.com/av1ppp/chafa-welcome/internal/sysinfo"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/av1ppp/chafa-welcome/internal/config"
	"github.com/av1ppp/chafa-welcome/internal/global"
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

	fmt.Println("PicturePath:", config_.PicturePath())
	fmt.Println("ChafaBin:", config_.ChafaBin())
	fmt.Println("Width:", config_.Width())

	// ====
	// ====
	// ====

	info, err := sysinfo.New()
	if err != nil {
		return err
	}
	infoLines := strings.Split(info.String(), "\n")

	chafaOutput, err := chafaExecute(config_)
	if err != nil {
		return err
	}
	//replaceNewValue := "  " + color.CyanString("Lorem:") + color.RedString("Ipsum") + "\n"
	//chafaOutput = strings.Replace(chafaOutput, "\n", replaceNewValue, -1)

	resultBuilder := strings.Builder{}

	for i, line := range strings.Split(chafaOutput, "\n") {
		if i < len(infoLines) {
			resultBuilder.WriteString(pictureMarginLeft + line + gap + infoLines[i] + "\n")
		} else {
			resultBuilder.WriteString(pictureMarginLeft + line + "\n")
		}
	}
	fmt.Println(resultBuilder.String())

	return nil
}

func chafaExecute(c *config.Config) (string, error) {
	args := []string{}

	if c.Width() != 0 {
		args = append(args, fmt.Sprintf("--size=%d", c.Width()))
	}

	args = append(args, c.PicturePath())

	cmd := exec.Command(c.ChafaBin(), args...)
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return stdout.String(), nil
}
