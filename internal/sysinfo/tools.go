package sysinfo

import (
	"github.com/av1ppp/chafa-welcome/internal/config"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func readFile(name string) (string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}

func execute(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func findAndTrimLinePrefix(data string, prefix string) (string, bool) {
	for _, line := range strings.Split(data, "\n") {
		if strings.HasPrefix(line, prefix) {
			return strings.Trim(strings.TrimPrefix(line, prefix), "\""), true
		}
	}
	return "", false
}

func bToMB(b uint64) uint64 {
	return b / 1024 / 1024
}

var themeAttributes = map[string]color.Attribute{
	"bold":         color.Bold,
	"faint":        color.Faint,
	"italic":       color.Italic,
	"underline":    color.Underline,
	"blinkslow":    color.BlinkSlow,
	"blinkrapid":   color.BlinkRapid,
	"reversevideo": color.ReverseVideo,
	"concealed":    color.Concealed,
	"crossedout":   color.CrossedOut,

	"fgblack":   color.FgBlack,
	"fgred":     color.FgRed,
	"fggreen":   color.FgGreen,
	"fgyellow":  color.FgYellow,
	"fgblue":    color.FgBlue,
	"fgmagenta": color.FgMagenta,
	"fgcyan":    color.FgCyan,
	"fgwhite":   color.FgWhite,

	"fghiblack":   color.FgHiBlack,
	"fghired":     color.FgHiRed,
	"fghigreen":   color.FgHiGreen,
	"fghiyellow":  color.FgHiYellow,
	"fghiblue":    color.FgHiBlue,
	"fghimagenta": color.FgHiMagenta,
	"fghicyan":    color.FgHiCyan,
	"fghiwhite":   color.FgHiWhite,

	"bgblack":   color.BgBlack,
	"bgred":     color.BgRed,
	"bggreen":   color.BgGreen,
	"bgyellow":  color.BgYellow,
	"bgblue":    color.BgBlue,
	"bgmagenta": color.BgMagenta,
	"bgcyan":    color.BgCyan,
	"bgwhite":   color.BgWhite,

	"bghiblack":   color.BgHiBlack,
	"bghired":     color.BgHiRed,
	"bghigreen":   color.BgHiGreen,
	"bghiyellow":  color.BgHiYellow,
	"bghiblue":    color.BgHiBlue,
	"bghimagenta": color.BgHiMagenta,
	"bghicyan":    color.BgHiCyan,
	"bghiwhite":   color.BgHiWhite,
}

func themeToColor(theme string) *color.Color {
	theme = strings.TrimSpace(theme)
	if theme == "" {
		return color.New()
	}

	themeParts := strings.Split(theme, " ")
	attrs := make([]color.Attribute, len(themeParts))

	var (
		attr  color.Attribute
		found bool
	)

	for i, part := range themeParts {
		attr, found = themeAttributes[part]
		if !found {
			panic("incorrect theme part: " + part)
		}
		attrs[i] = attr
	}

	return color.New(attrs...)
}

type collectFunc = func(conf *config.Config) (string, error)

func collectIfInclude(include bool, conf *config.Config, collect collectFunc) (string, error) {
	if include {
		return collect(conf)
	}
	return "", nil
}
