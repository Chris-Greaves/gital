package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Chris-Greaves/gital/core"
	"github.com/Chris-Greaves/gital/core/db"
)

var (
	ErrScanDelayUnparseable = fmt.Errorf("unable to parse %v", core.KeyScanDelay)
)

type Scheduler struct {
	config *core.Config
	ctx    context.Context
}

type Job func(ctx context.Context, db db.Database, directories []string) error

func CreateScheduler(config *core.Config, ctx context.Context) Scheduler {
	return Scheduler{
		config: config,
		ctx:    ctx,
	}
}

func (s *Scheduler) Run(job Job, db db.Database, done chan bool) error {
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
			err := job(s.ctx, db, s.config.GetStringSlice(core.KeyDirectories))
			if err != nil {
				slog.Error("Error occured while running scheduled job", slog.Any("error", err))
			}
			sleepWithContext(s.ctx, wait)
		}
	}
}

func sleepWithContext(ctx context.Context, dur time.Duration) {
	select {
	case <-time.After(dur):
		return
	case <-ctx.Done():
		return
	}
}
