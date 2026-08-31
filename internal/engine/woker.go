package engine

import (
	"context"
	"fmt"
)

func worker(
	ctx context.Context,
	jobs <-chan Job,
	results chan<- Result,
) {
	for {
		select {
		case <-ctx.Done():
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}
			result := processJob(ctx, job)
			results <- result
		}
	}
}

func processJobsSafely(ctx context.Context, job Job, results chan<- Result) {
	defer func() {
		if r := recover(); r != nil {
			results <- Result{
				Job: job,
				Err: fmt.Errorf("panic while processing job: %v", r),
			}
		}
	}()

	result := processJob(ctx, job)
	results <- result
}

func processJob(ctx context.Context, job Job) Result {
	if job.URL == "" {
		return Result{
			Job: job,
			Err: fmt.Errorf("empty URL"),
		}
	}
	return Result{
		Job: job,
	}
}
