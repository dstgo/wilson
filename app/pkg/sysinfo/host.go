package sysinfo

import (
	"github.com/dstgo/size"
	"github.com/shirou/gopsutil/host"
)

type HostInfo struct {
	Id       string
	Name     string
	Version  string
	Arch     string
	Platform string
	Os       string
	Family   string
	Cpu      int
	Mem      size.SizeMeta
}

func GetHostInfo() HostInfo {
	var info HostInfo
	stat, err := host.Info()
	if err != nil {
		return info
	}

	info = HostInfo{
		Id:       stat.HostID,
		Name:     stat.Hostname,
		Version:  stat.KernelVersion,
		Arch:     stat.KernelArch,
		Platform: stat.Platform,
		Os:       stat.OS,
		Family:   stat.PlatformFamily,
	}
	return info
}
