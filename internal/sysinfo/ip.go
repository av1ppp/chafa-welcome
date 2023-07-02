package sysinfo

import (
	"strings"

	"github.com/av1ppp/chafa-welcome/internal/config"
)

func collectLocalIP(conf *config.Config) (string, error) {
	localIP, err := execute("sh", "-c", "ip route get 1.2.3.4 | awk '{print $7}'")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(localIP), nil
}

func collectGlobalIP(conf *config.Config) (string, error) {
	return execute("curl", "ifconfig.me")
}
