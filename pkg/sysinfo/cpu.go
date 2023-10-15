package sysinfo

import "github.com/shirou/gopsutil/cpu"

type CpuInfo struct {
	Count int
	Name  string
}

func GetCpuInfo() CpuInfo {
	var info CpuInfo
	counts, err := cpu.Counts(true)
	if err != nil {
		return info
	}
	info.Count = counts
	stats, err := cpu.Info()
	if err != nil {
		return info
	}
	if len(stats) > 0 {
		info.Name = stats[0].ModelName
	}
	return info
}
