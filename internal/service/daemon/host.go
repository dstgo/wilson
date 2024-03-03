package daemon

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/dstgo/wilson/internal/conf"
	v1 "github.com/dstgo/wilson/internal/proto/api/v1"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"runtime"
	"time"
)

func NewHostHandler(dockerClient *client.Client, config *conf.WigfridConf) *HostHandler {
	return &HostHandler{
		dockerClient: dockerClient,
		config:       config,
	}
}

// HostHandler returns some information about host machine
type HostHandler struct {
	dockerClient *client.Client
	config       *conf.WigfridConf
}

// HostInfo returns host basic information
func (h *HostHandler) HostInfo(ctx context.Context) (*v1.SystemInfo, error) {

	// docker and host info
	dockerInfo, err := h.dockerClient.Info(ctx)
	if err != nil {
		return nil, err
	}

	hostInfo, err := host.Info()
	if err != nil {
		return nil, err
	}

	result := &v1.SystemInfo{
		Os:           hostInfo.OS,
		OsVersion:    hostInfo.KernelVersion,
		Arch:         hostInfo.KernelArch,
		GoVersion:    runtime.Version(),
		BuildVersion: h.config.BuildMeta.Version,
		Docker: &v1.DockerInfo{
			Containers:    int64(dockerInfo.Containers),
			Running:       int64(dockerInfo.ContainersRunning),
			Pause:         int64(dockerInfo.ContainersPaused),
			Stopped:       int64(dockerInfo.ContainersStopped),
			Images:        int64(dockerInfo.Images),
			Driver:        dockerInfo.Driver,
			Version:       dockerInfo.ServerVersion,
			KernelVersion: dockerInfo.KernelVersion,
		},
	}

	memoryInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	var totalDiskBytes int64
	partitions, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	for _, partition := range partitions {
		usage, _ := disk.Usage(partition.Mountpoint)
		totalDiskBytes += int64(usage.Total)
	}

	resource := &v1.Resource{
		Cpu:    int64(runtime.NumCPU()),
		Memory: int64(memoryInfo.Total),
		Disk:   totalDiskBytes,
	}

	result.Resource = resource

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	// cpu model name
	if len(cpuInfo) > 0 {
		result.CpuModel = cpuInfo[0].ModelName
	}

	return result, nil
}

// HostHealth returns the health check information of the host machine
func (h *HostHandler) HostHealth(ctx context.Context) (*v1.HealthInfo, error) {

	// cpu usage
	percent, err := cpu.Percent(time.Millisecond*500, false)
	if err != nil {
		return nil, err
	}

	cpuHealth := &v1.CpuHealth{
		Count: int64(runtime.NumCPU()),
	}

	if len(percent) > 0 {
		cpuHealth.Usage = percent[0]
	}

	// mem usage
	memory, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	memHealth := &v1.MemoryHealth{
		Total: int64(memory.Total),
		Used:  int64(memory.Used),
		Free:  int64(memory.Free),
		Usage: memory.UsedPercent,
	}

	return &v1.HealthInfo{
		Cpu: cpuHealth,
		Mem: memHealth,
	}, nil
}
