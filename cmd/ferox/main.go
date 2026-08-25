package main

import (
	"context"
	"fmt"
	"time"

	"github.com/codewithMohak/Ferox.git/internal/engine"
)

func main() {
	ctx := context.Background()

	jobs := make(chan engine.Job)

	go func() {
		defer close(jobs)

		urls := []string{
			"https://example.com/",
			"https://example.com/admin",
			"https://example.com/login",
			"https://example.com/api",
		}

		for _, url := range urls {
			jobs <- engine.Job{URL: url}
		}
	}()

	results := engine.Run(ctx, engine.Config{
		Concurrency: 2,
	}, jobs)

	for result := range results {
		fmt.Printf(
			"URL=%s STATUS=%d ERROR=%v TIME=%v\n",
			result.Job.URL,
			result.StatusCode,
			result.Err,
			result.Duration,
		)
	}
	time.Sleep(100 * time.Millisecond)
}
