package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Chris-Greaves/gital/core"
)

var (
	ErrScanDelayUnparseable = fmt.Errorf("unable to parse %v", core.KeyScanDelay)
)

type Scheduler struct {
	config *core.Config
	ctx    context.Context
}

func CreateScheduler(config *core.Config, ctx context.Context) Scheduler {
	return Scheduler{
		config: config,
		ctx:    ctx,
	}
}

func (s *Scheduler) Run(done chan bool) error {
	wait := s.config.GetDuration(core.KeyScanDelay)
	if wait == time.Millisecond*0 {
		return ErrScanDelayUnparseable
	}

	slog.Info("Scheduler running", slog.Duration("scan_delay", wait))

	for {
		select {
		case <-s.ctx.Done():
			// Handle context cancellation (graceful shutdown)
			slog.Info("Scheduler stopping.", slog.Any("error", s.ctx.Err()))
			return nil
		default:
			// Do Work
			time.Sleep(wait)
		}
	}
}
