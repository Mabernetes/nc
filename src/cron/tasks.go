package cron

import "github.com/Mabernetes/nc/src/logic"

type Runner interface {
	Start()
}

type Cron struct {
	Runner Runner
}

func NewCronHandler(logic logic.Config) *Cron {
	return &Cron{
		Runner: NewRunnerTask(logic),
	}
}

func (c *Cron) Start() {
	c.Runner.Start()
}
