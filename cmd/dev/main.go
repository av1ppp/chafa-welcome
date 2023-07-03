package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/av1ppp/chafa-welcome/internal/chafa"
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
	homeDir := global.HomeDir()
	confPath := filepath.Join(homeDir, "config")
	conf, err := config.ParseFile(confPath)
	if err != nil {
		return err
	}

	gap := strings.Repeat(" ", conf.Body.Gap)
	pictureMarginLeft := strings.Repeat(" ", conf.Offset.X)

	fmt.Println("Source:", conf.Image.Source)
	fmt.Println("Chafa:", conf.ChafaBin)
	fmt.Println("Siz:", conf.Image.Size)

	info, err := sysinfo.Collect(conf)
	if err != nil {
		return err
	}
	infoLines := strings.Split(info.String(), "\n")
	infoNumberLines := len(infoLines)
	fmt.Println("infoNumberLines:", infoNumberLines)

	chafaOutput, err := chafa.Execute(conf)
	if err != nil {
		return err
	}
	chafaLines := strings.Split(chafaOutput, "\n")
	chafaNumberLines := len(chafaLines) - 1
	chafaEmptyRow := strings.Repeat(" ", conf.Image.Size)
	fmt.Println("chafaNumberLines:", chafaNumberLines)

	maxLines := 0
	if infoNumberLines > chafaNumberLines {
		maxLines = infoNumberLines
	} else {
		maxLines = chafaNumberLines
	}
	fmt.Println("maxLines:", maxLines)

	resultBuilder := strings.Builder{}

	for i := 0; i < conf.Offset.Y; i++ {
		resultBuilder.WriteByte('\n')
	}

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
