package sysinfo

import (
	"errors"
	"os"

	"github.com/av1ppp/chafa-welcome/internal/config"
)

func collectTerminal(conf *config.Config) (string, error) {
	termProgram, exists := os.LookupEnv("TERM_PROGRAM")
	if exists {
		switch termProgram {
		case "iTerm.app":
			return "iTerm2", nil
		case "Terminal.app":
			return "Apple Terminal", nil
		default:
			return termProgram, nil
		}
	}

	term, termExists := os.LookupEnv("TERM")
	if term == "tw52" || term == "tw100" {
		return "TosWin2", nil
	}

	if _, exists = os.LookupEnv("SSH_CONNECTION"); exists {
		return os.Getenv("SSH_TTY"), nil
	}

	if _, exists = os.LookupEnv("WT_SESSION"); exists {
		return "Windows Terminal", nil
	}

	if !termExists {
		return "", errors.New("TERM environment not set")
	}
	return term, nil
}
