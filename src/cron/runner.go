package cron

import (
	"log"
	"node/src/logic"
	"os"
	"os/exec"
)

type RunnerTask struct {
	configsDir string
	logic      logic.Config
}

func NewRunnerTask(logic logic.Config) *RunnerTask {
	dir := os.Getenv("M8S_CONFIG_DIR")
	if dir == "" {
		dir = "~/m8s"
	}
	return &RunnerTask{
		configsDir: dir,
		logic:      logic,
	}
}

func (t *RunnerTask) Start() {
	t.ApplyConfigs()
}

func (t *RunnerTask) ApplyConfigs() {
	tree, err := t.logic.GetTree()
	if err != nil {
		return
	}
	var path string
	for _, file := range tree {
		path = t.logic.GetFilePath(file.Deployment, file.Pod)
		command := exec.Command("docker", "compose", "--file", path, "up", "-d")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err = command.Run(); err != nil {
			log.Println("error Cron.Runner.Start.ApplyConfigs:", err)
		}
	}
}
