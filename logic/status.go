package logic

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"math"
	"node/utils"
)

type StatusLogic struct {
	cli *client.Client
}

func NewStatusLogic() *StatusLogic {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	return &StatusLogic{
		cli: cli,
	}
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

type DeploymentStatusData map[string]struct {
	Started int `json:"started"`
	Stopped int `json:"stopped"`
	Total   int `json:"total"`
}

func (l StatusLogic) Deployment(deployment, pod string) ([]types.Container, error) {
	var containers []types.Container
	var err error
	containers, err = l.cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return containers, err
	}

	if deployment != "" {
		containers = utils.ContainersDeploymentFilter(deployment, containers)
	}
	if pod != "" {
		containers = utils.ContainersPodFilter(pod, containers)
	}

	return containers, nil
}
