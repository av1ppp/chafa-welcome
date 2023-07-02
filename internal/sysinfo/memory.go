package sysinfo

import (
	"strconv"

	"github.com/shirou/gopsutil/v3/mem"

	"github.com/av1ppp/chafa-welcome/internal/config"
)

func collectMemory(conf *config.Config) (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	usedMbStr := strconv.FormatUint(bToMB(vm.Used), 10)
	totalMbStr := strconv.FormatUint(bToMB(vm.Total), 10)
	usedPercentMbStr := strconv.Itoa(int(vm.UsedPercent))

	if conf.Body.Memory.Percent {
		return usedMbStr + "MB / " + totalMbStr + "MB" + " (" + usedPercentMbStr + "%)", nil
	}
	return usedMbStr + "MB / " + totalMbStr + "MB", nil
}
