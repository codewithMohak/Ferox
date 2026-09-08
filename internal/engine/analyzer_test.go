package engine // Places this test inside the engine package.

import (
	"testing" // Provides Go's testing framework.

	"github.com/codewithMohak/Ferox.git/internal/fingerprint"
)

// TestAnalyze verifies that the complete response-analysis pipeline works.
func TestAnalyze(t *testing.T) {

	// Create a baseline response body.
	baselineBody := []byte("Access denied")

	// Generate the baseline fingerprint using the same fingerprinting logic as the analyzer.
	baselineFingerprint := fingerprint.Fingerprint(baselineBody)

	// Analyze a response that is identical to the baseline.
	result := Analyze(AnalyzeInput{
		StatusCode:     403,                      // The response returned HTTP 403.
		Body:           baselineBody,             // Use the baseline response body.
		Baseline:       baselineFingerprint,      // Tell the analyzer this is the baseline fingerprint.
		Length:         int64(len(baselineBody)), // Calculate the response body length.
		BaselineLength: int64(len(baselineBody)), // Use the same length as the baseline.
	})

	// The analyzer should recognize the response as a block page.
	if !result.IsBlockPage {
		t.Fatal("expected response to be detected as a block page")
	}

	// The response should have a fingerprint.
	if result.Fingerprint == "" {
		t.Fatal("expected fingerprint to be generated")
	}
}
