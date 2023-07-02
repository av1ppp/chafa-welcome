package sysinfo

import (
	"os"
)

func collectTerminal() (string, error) {
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

	term, exists := os.LookupEnv("TERM")
	if term == "tw52" || term == "tw100" {
		return "TosWin2", nil
	}

	if _, exists = os.LookupEnv("SSH_CONNECTION"); exists {
		return os.Getenv("SSH_TTY"), nil
	}

	if _, exists = os.LookupEnv("WT_SESSION"); exists {
		return "Windows Terminal", nil
	}

	if term != "" {
		return term, nil
	}
	return "-", nil
}
