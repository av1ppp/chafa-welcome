package sysinfo

import "strings"

func collectLocalIP() (string, error) {
	localIP, err := execute("sh", "-c", "ip route get 1.2.3.4 | awk '{print $7}'")
	if err != nil {
		return "-", nil
	}
	return strings.TrimSpace(localIP), nil
}

func collectGlobalIP() (string, error) {
	return execute("curl", "ifconfig.me")
}
