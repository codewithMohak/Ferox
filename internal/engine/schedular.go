package engine

import (
	"context"
	"sync"
)

type Config struct {
	Concurrency int
}

func Run(ctx context.Context, cfg Config, jobs <-chan Job) <-chan Result {
	results := make(chan Result)

	var wg sync.WaitGroup

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, jobs, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}
