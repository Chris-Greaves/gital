package main

import "github.com/Chris-Greaves/gital/core"

type Scheduler struct {
	config *core.Config
}

func CreateScheduler(config *core.Config) Scheduler {
	return Scheduler{
		config: config,
	}
}

func (s *Scheduler) Run() error {
	return nil
}
