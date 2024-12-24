package cron

import (
	"os"
	"os/exec"
)

type RunnerTaskUpdateJob struct {
	*Manager
}

func (t *RunnerTaskUpdateJob) Run() {
	dir := os.Getenv("M8S_CONFIG_DIR")
	if dir == "" {
		dir = "~/m8s"
	}

	tree, err := t.services.Config.GetTree()
	if err != nil {
		return
	}
	var path string
	for _, file := range tree {
		path = t.services.Config.GetFilePath(file.Deployment, file.Pod)
		command := exec.Command("docker", "compose", "--file", path, "up", "-d")
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err = command.Run(); err != nil {
			t.log.Println("[CRON] error Cron.Runner.Start.ApplyConfigs:", err)
		}
	}
}
