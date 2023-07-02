package sysinfo

import (
	"github.com/shirou/gopsutil/v3/mem"
	"strconv"
)

func collectMemory() (string, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}

	usedMbStr := strconv.FormatUint(bToMB(vm.Used), 10)
	totalMbStr := strconv.FormatUint(bToMB(vm.Total), 10)
	usedPercentMbStr := strconv.Itoa(int(vm.UsedPercent))

	return usedMbStr + "MB / " + totalMbStr + "MB" + " (" + usedPercentMbStr + "%)", nil
}
