package engine

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// MaxResponseBodySize prevents Ferox from reading an unreasonably large response into memory.
const MaxResponseBodySize = 2 * 1024 * 1024

func worker(
	ctx context.Context,
	client *http.Client,
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
			processJobSafely(ctx, client, job, results)
		}
	}
}

func processJobSafely(
	ctx context.Context,
	client *http.Client,
	job Job,
	results chan<- Result) {
	defer func() {
		if r := recover(); r != nil {
			results <- Result{
				Job: job,
				Err: fmt.Errorf("panic while processing job: %v", r),
			}
		}
	}()

	result := processJob(ctx, client, job)
	results <- result
}

func processJob(ctx context.Context,
	client *http.Client,
	job Job,
) Result {
	start := time.Now()

	req, err := http.NewRequestWithContext(
		ctx,            // Allows cancellation to propagate into the HTTP request.
		http.MethodGet, // Use HTTP GET for this initial implementation.
		job.URL,        // Use the URL supplied by the current job.
		nil,            // The request does not have a body.
	)
	if err != nil {
		return Result{
			Job:      job,               // Preserve the original job.
			Duration: time.Since(start), // Record how long processing took before failure.
			Err:      err,               // Store the request creation error.
		}
	}
	if job.URL == "" {
		return Result{
			Job: job,
			Err: fmt.Errorf("empty URL"),
		}
	}

	resp, err := client.Do(req)
	// Check whether the HTTP request failed.
	if err != nil {

		// Return the error associated with this job.
		return Result{
			Job:      job,               // Preserve the original job.
			Duration: time.Since(start), // Record the elapsed time.
			Err:      err,               // Store the HTTP error.
		}
	}

	defer resp.Body.Close()

	limitedBody := io.LimitReader(
		resp.Body,           // Read from the HTTP response body.
		MaxResponseBodySize, // Stop reading after the configured maximum.
	)

	body, err := io.ReadAll(limitedBody)

	if err != nil {
		return Result{
			Job:        job,               // Preserve the original job.
			StatusCode: resp.StatusCode,   // Preserve the HTTP status received from the server.
			Duration:   time.Since(start), // Record total processing time.
			Err:        err,               // Store the body-reading error.
		}
	}
	// Analyze the response using the fingerprinting and scoring pipeline.
	analyzed := Analyze(AnalyzeInput{
		StatusCode: resp.StatusCode,  // Pass the HTTP status to the analyzer.
		Body:       body,             // Pass the response body to the analyzer.
		Length:     int64(len(body)), // Record the number of bytes read.
	})

	analyzed.Job = job

	// Record how long the complete request and analysis took.
	analyzed.Duration = time.Since(start)

	// Return the final result.
	return analyzed
}
