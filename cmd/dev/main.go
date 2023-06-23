package main

import (
	"fmt"
	"path/filepath"

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
	fmt.Println("ChafaArgs:", config_.ChafaArgs())

	return nil
}
