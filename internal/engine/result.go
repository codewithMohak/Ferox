package engine

import "time"

type Job struct {
	URL string
}

type Result struct {
	Job        Job
	StatusCode int
	Length     int64
	Duration   time.Duration
	Err        error
}
