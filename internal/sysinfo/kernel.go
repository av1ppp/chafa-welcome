package sysinfo

import "github.com/av1ppp/chafa-welcome/internal/config"

func collectKernel(conf *config.Config) (string, error) {
	return execute("uname", "-sr")
}
