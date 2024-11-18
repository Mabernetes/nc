package logic

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"math"
)

type StatusLogic struct {
}

func NewStatusLogic() *StatusLogic {
	return &StatusLogic{}
}

type ServerStatusData struct {
	Cpu struct {
		Max int `json:"max"`
		Use int `json:"use"`
	} `json:"cpu"`
	Mem struct {
		Max int `json:"max"`
		Use int `json:"use"`
	} `json:"mem"`
	Disk struct {
		Max int `json:"max"`
		Use int `json:"use"`
	} `json:"disk"`
}

func (l StatusLogic) Server() ServerStatusData {
	var status ServerStatusData

	cpuInfo, err := cpu.Info()
	if err == nil && len(cpuInfo) > 0 {
		status.Cpu.Max = len(cpuInfo) * 100 // 100% на каждое ядро
	}

	cpuUsage, err := cpu.Percent(0, false)
	if err == nil && len(cpuUsage) > 0 {
		status.Cpu.Use = int(math.Round(cpuUsage[0]))
	}

	// Память
	memInfo, err := mem.VirtualMemory()
	if err == nil {
		status.Mem.Max = int(memInfo.Total / 1024 / 1024) // в MB
		status.Mem.Use = int(memInfo.Used / 1024 / 1024)  // в MB
	}

	// Диск
	diskInfo, err := disk.Usage("/")
	if err == nil {
		status.Disk.Max = int(diskInfo.Total / 1024 / 1024) // в MB
		status.Disk.Use = int(diskInfo.Used / 1024 / 1024)  // в MB
	}

	return status
}
