package sysinfo

import (
	"github.com/dstgo/size"
	"github.com/shirou/gopsutil/mem"
)

type Mem struct {
	Total size.SizeMeta
	Avi   size.SizeMeta
	Used  size.SizeMeta
}

type MemInfo struct {
	Virtual Mem
	Swap    Mem
}

func GetMemInfo() MemInfo {
	var info MemInfo
	memory, err := mem.VirtualMemory()
	if err != nil {
		return info
	}
	info.Virtual = Mem{
		Total: size.NewSize(float64(int64(memory.Total)), size.B),
		Avi:   size.NewSize(float64(int64(memory.Available)), size.B),
		Used:  size.NewSize(float64(int64(memory.Used)), size.B),
	}
	swap, err := mem.SwapMemory()
	if err != nil {
		return info
	}
	info.Swap = Mem{
		Total: size.NewSize(float64(int64(swap.Total)), size.B),
		Avi:   size.NewSize(float64(int64(swap.Free)), size.B),
		Used:  size.NewSize(float64(int64(swap.Used)), size.B),
	}
	return info
}
