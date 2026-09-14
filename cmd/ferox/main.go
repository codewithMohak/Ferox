package main // Defines the executable's main package.

import (
	"context" // Provides context for controlling the engine.
	"flag"
	"fmt" // Provides formatted terminal output.
	"log"
	"time"

	"github.com/codewithMohak/Ferox.git/internal/engine" // Imports Ferox's engine package.
	"github.com/codewithMohak/Ferox.git/internal/wordlist"
)

func main() {

	// Define the -u command-line argument for the wordlist path
	targetURL := flag.String(
		"u",
		"",
		"Traget URL to scan",
	)

	// Define the -w command-line argument for the worldlist path
	wordlistPath := flag.String(
		"w",
		"",
		"Path to the wordlist",
	)

	//Define the -c command-line argument for the worker count.
	concurrency := flag.Int(
		"c",
		10,
		"Number of Concurrent workers",
	)

	flag.Parse()

	if *targetURL == "" {
		log.Fatal("target URL is required use -u")
	}

	if *wordlistPath == "" {
		log.Fatal("wordlist is required: use -w")
	}

	if *concurrency < 1 {
		log.Fatal("concurrency must be at least 1")
	}

	words, err := wordlist.Load(*wordlistPath)

	if err != nil {
		log.Fatalf("load wordlist: %v", err)
	}

	// Create a context that remains active for the lifetime of this simple test.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Minute,
	)

	defer cancel()

	// Create the channel used to send jobs to the worker pool.
	jobs := make(chan engine.Job)

	// Start a goroutine that produces test jobs.
	go func() {

		defer close(jobs)

		for _, word := range words {
			url, err := engine.BuildURL(*targetURL, word)

			if err != nil {
				continue
			}
			// Send the generated URL to the worker pool.
			select {

			// Send the job when workers are ready to receive it.
			case jobs <- engine.Job{
				URL: url,
			}:

			// Stop generating jobs if the scan context is cancelled.
			case <-ctx.Done():
				return
			}
		}
	}()

	results := engine.Run(
		ctx,
		engine.Config{
			Concurrency: *concurrency,
		},
		jobs,
	)

	fmt.Printf(
		"Ferox scanning %s with %d workers\n\n",
		*targetURL,
		*concurrency,
	)

	for result := range results {
		if result.Err != nil {
			fmt.Printf(
				"ERROR %s %v\n",
				result.Job.URL,
				result.Err,
			)
			continue
		}
		fmt.Printf(
			"%d %-50s length=%d score=%d block=%v time=%v\n",
			result.StatusCode,
			result.Job.URL,
			result.Length,
			result.AnomalyScore,
			result.IsBlockPage,
			result.Duration,
		)
	}
}
