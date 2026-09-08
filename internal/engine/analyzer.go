package engine

import "github.com/codewithMohak/Ferox.git/internal/fingerprint"

// AnalyzeInput contains the information required to analyze one HTTP response.
type AnalyzeInput struct {
	StatusCode     int
	Body           []byte
	Baseline       string
	Length         int64
	BaselineLength int64
}

func Analyze(input AnalyzeInput) Result {
	responseFingerprint := fingerprint.Fingerprint(input.Body)

	blockPage := fingerprint.IsBlockPage(
		input.StatusCode,
		input.Body,
	)

	score := fingerprint.Score(fingerprint.ScoreInput{
		StatusCode:     input.StatusCode,     // Pass the HTTP status code to the scorer.
		Fingerprint:    responseFingerprint,  // Pass the calculated fingerprint.
		Baseline:       input.Baseline,       // Pass the known baseline fingerprint.
		IsBlockPage:    blockPage,            // Pass the block-page classification.
		Length:         input.Length,         // Pass the response length.
		BaselineLength: input.BaselineLength, // Pass the baseline response length.
	})
	return Result{
		StatusCode:   input.StatusCode,    // Preserve the original HTTP status code.
		Length:       input.Length,        // Preserve the response length.
		Body:         input.Body,          // Preserve the original response body.
		Fingerprint:  responseFingerprint, // Store the calculated fingerprint.
		IsBlockPage:  blockPage,           // Store the block-page classification.
		AnomalyScore: score,               // Store the calculated anomaly score.
	}
}
