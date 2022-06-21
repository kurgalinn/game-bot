package service

import (
	"time"
)

const (
	SyncDuration = 1
	JobPeriod    = 2
)

type Worker interface {
	Add(key string, period time.Duration, callback func())
	Remove(key string)
	Execute(key string)
	TimeLeft(key string) int
}

type job struct {
	time     time.Time
	callback func()
}

type worker struct {
	jobs map[string]job
}

func NewWorker() *worker {
	w := &worker{jobs: make(map[string]job)}
	ticker := time.NewTicker(JobPeriod * time.Second)
	go func() {
		for {
			select {
			case _ = <-ticker.C:
				// TODO: optimize it, sorted list with break
				for key, job := range w.jobs {
					if job.timeLeft()+SyncDuration <= 0 {
						job.callback()
						w.Remove(key)
					}
				}
			}
		}
	}()
	return w
}

func (w worker) Add(key string, period time.Duration, callback func()) {
	w.jobs[key] = job{
		time:     time.Now().Add(period),
		callback: callback,
	}
}

func (w worker) Remove(key string) {
	delete(w.jobs, key)
}

func (w worker) Execute(key string) {
	w.jobs[key].callback()
	w.Remove(key)
}

func (w worker) TimeLeft(key string) int {
	return w.jobs[key].timeLeft()
}

func (j job) timeLeft() int {
	return int(j.time.Unix() - time.Now().Unix())
}
