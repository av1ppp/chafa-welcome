package sysinfo

func collectUptime() (string, error) {
	return execute("uptime", "-p")
}
