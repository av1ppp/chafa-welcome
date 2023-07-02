package sysinfo

import "github.com/av1ppp/chafa-welcome/internal/config"

func collectUptime(conf *config.Config) (string, error) {
	return execute("uptime", "-p")
}
