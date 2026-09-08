package main // Defines the executable's main package.

import (
	"context" // Provides context for controlling the engine.
	"fmt"     // Provides formatted terminal output.

	"github.com/codewithMohak/Ferox.git/internal/engine" // Imports Ferox's engine package.
)

func main() {

	// Create a context that remains active for the lifetime of this simple test.
	ctx := context.Background()

	// Create the channel used to send jobs to the worker pool.
	jobs := make(chan engine.Job)

	// Start a goroutine that produces test jobs.
	go func() {

		// Close the jobs channel after all URLs have been submitted.
		defer close(jobs)

		// Define URLs that Ferox will request during this integration test.
		urls := []string{
			"http://127.0.0.1:8080/",       // Test the local root endpoint.
			"http://127.0.0.1:8080/admin",  // Test a local admin path.
			"http://127.0.0.1:8080/random", // Test a local random path.
		}

		// Submit every URL as an engine job.
		for _, url := range urls {

			// Send the current URL to the worker pool.
			jobs <- engine.Job{
				URL: url, // Store the URL inside the job.
			}
		}
	}()

	// Start two concurrent workers.
	results := engine.Run(
		ctx, // Pass the active context.
		engine.Config{
			Concurrency: 2, // Allow two requests to be processed concurrently.
		},
		jobs, // Give the worker pool the job channel.
	)

	// Process results until every worker has finished.
	for result := range results {

		// Print the important information collected from the HTTP response.
		fmt.Printf(
			"URL=%s STATUS=%d LENGTH=%d SCORE=%d BLOCK=%v TIME=%v ERROR=%v\n",
			result.Job.URL,      // Print the URL that was requested.
			result.StatusCode,   // Print the HTTP status code.
			result.Length,       // Print the response body length.
			result.AnomalyScore, // Print the calculated anomaly score.
			result.IsBlockPage,  // Print whether a block page was detected.
			result.Duration,     // Print request and processing duration.
			result.Err,          // Print any error associated with the job.
		)
	}
}
