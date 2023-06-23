package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"

	"github.com/av1ppp/chafa-welcome/internal/config"
	"github.com/av1ppp/chafa-welcome/internal/global"
)

func main() {
	if err := innerMain(); err != nil {
		panic(err)
	}
}

func innerMain() error {
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

	chafaOutput, err := chafaExecute(config_)
	if err != nil {
		return err
	}

	replaceNewValue := "  " + color.CyanString("Lorem:") + color.RedString("Ipsum") + "\n"
	chafaOutput = strings.Replace(chafaOutput, "\n", replaceNewValue, -1)
	fmt.Println(chafaOutput)

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
