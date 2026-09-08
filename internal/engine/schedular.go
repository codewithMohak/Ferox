package engine

import (
	"context"
	"net/http"
	"sync"

	"github.com/codewithMohak/Ferox.git/internal/httpclient"
)

type Config struct {
	Concurrency int
	Client      *http.Client
}

func Run(ctx context.Context, cfg Config, jobs <-chan Job) <-chan Result {
	results := make(chan Result)

	if cfg.Client == nil {
		cfg.Client = httpclient.New(httpclient.Options{})
	}

	var wg sync.WaitGroup

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, cfg.Client, jobs, results)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}
