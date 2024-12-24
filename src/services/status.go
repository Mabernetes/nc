package services

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"math"
)

type Status interface {
	Server() ServerStatusData
	Runner() (map[string]DeploymentStatusData, error)
}

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

	// Процессор
	status.Cpu.Max, _ = cpu.Counts(true)
	cpuUsage, _ := cpu.Percent(0, false)
	status.Cpu.Use = int(math.Round(cpuUsage[0]))

	// Память
	memInfo, _ := mem.VirtualMemory()
	status.Mem.Max = int(memInfo.Total / 1024 / 1024) // в MB
	status.Mem.Use = int(memInfo.Used / 1024 / 1024)  // в MB

	// Диск
	diskInfo, _ := disk.Usage("/")
	status.Disk.Max = int(diskInfo.Total / 1024 / 1024) // в MB
	status.Disk.Use = int(diskInfo.Used / 1024 / 1024)  // в MB

	return status
}

type DeploymentStatusData struct {
	Started int `json:"started"`
	Stopped int `json:"stopped"`
	Total   int `json:"total"`
}

func (l StatusLogic) Runner() (map[string]DeploymentStatusData, error) {
	var out map[string]DeploymentStatusData
	var containers []types.Container
	var err error
	containers, err = l.cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return out, err
	}

	for _, c := range containers {
		if label, ok := c.Labels["ru.m8s.deployment.name"]; ok {
			var deployStatus DeploymentStatusData = out[label]

			deployStatus.Total = deployStatus.Total + 1
			switch c.State {
			case "running":
				deployStatus.Started = deployStatus.Started + 1
			default:
				deployStatus.Started = deployStatus.Started + 1
			}

			out[label] = deployStatus
		}
	}

	return out, nil
}
