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
	config, err := config.ParseFile(configPath)
	if err != nil {
		return err
	}

	fmt.Println("PicturePath:", config.PicturePath())
	fmt.Println("ChafaBin:", config.ChafaBin())
	fmt.Println("ChafaArgs:", config.ChafaArgs())

	return nil
}
