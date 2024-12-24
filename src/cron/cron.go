package cron

import (
	"github.com/Mabernetes/nc/src/services"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)

type Manager struct {
	services *services.Services
	log      *log.Logger
	cron     *cron.Cron
}

type JobsList []struct {
	Name    string
	Time    string
	JobInit func(manager *Manager) cron.Job
}

var jobs JobsList = JobsList{
	{
		Name:    "ConfigUpdate",
		Time:    "* * * * *",
		JobInit: func(manager *Manager) cron.Job { return &RunnerTaskUpdateJob{manager} },
	},
}

func Start(services *services.Services) {
	manager := &Manager{
		services: services,
		log:      log.New(os.Stdout, "[CRON] ", 0),
		cron:     cron.New(),
	}

	for _, job := range jobs {
		_, err := manager.cron.AddJob(job.Time, job.JobInit(manager))
		if err != nil {
			manager.log.Println("Cron error adding job:", err)
		} else {
			manager.log.Println("Cron added job:", job.Name)
		}
	}

	manager.log.Println("Cron run...")
	manager.cron.Start()
}
