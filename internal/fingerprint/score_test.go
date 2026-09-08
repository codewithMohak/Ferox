package fingerprint // Places this test in the fingerprint package.

import "testing" // Provides Go's testing framework.

// TestScore uses table-driven tests to verify different scoring scenarios.
func TestScore(t *testing.T) {

	// Define the scenarios we want to test.
	tests := []struct {
		name  string     // Gives the test scenario a readable name.
		input ScoreInput // Contains the signals provided to the scorer.
		want  int        // Contains the expected anomaly score.
	}{
		{
			name: "normal baseline response", // Represents a response matching the baseline.
			input: ScoreInput{
				StatusCode:     404,   // The response returned HTTP 404.
				Fingerprint:    "A",   // The response fingerprint is A.
				Baseline:       "A",   // The baseline fingerprint is also A.
				IsBlockPage:    false, // The response isn't identified as a block page.
				Length:         1000,  // The response body is 1000 bytes.
				BaselineLength: 1000,  // The baseline response is also 1000 bytes.
			},
			want: 0, // No signal indicates an interesting deviation.
		},
		{
			name: "different successful response", // Represents a response different from the baseline.
			input: ScoreInput{
				StatusCode:     200,   // The response returned HTTP 200.
				Fingerprint:    "B",   // The response has fingerprint B.
				Baseline:       "A",   // The normal baseline fingerprint is A.
				IsBlockPage:    false, // The response isn't a block page.
				Length:         2000,  // The response is significantly larger.
				BaselineLength: 1000,  // The baseline response is 1000 bytes.
			},
			want: 80, // Different fingerprint (+50) + 200 (+20) + length difference (+10).
		},
		{
			name: "blocked response", // Represents a generic security block page.
			input: ScoreInput{
				StatusCode:     403,  // The response returned HTTP 403.
				Fingerprint:    "B",  // The response has a different fingerprint.
				Baseline:       "A",  // The baseline fingerprint is A.
				IsBlockPage:    true, // The response looks like a block page.
				Length:         1000, // The response has the same length as baseline.
				BaselineLength: 1000, // The baseline response is 1000 bytes.
			},
			want: 20, // Different fingerprint (+50) minus block-page penalty (-30).
		},
	}

	// Run every scoring scenario independently.
	for _, tt := range tests {

		// Create a named sub-test for the current scenario.
		t.Run(tt.name, func(t *testing.T) {

			// Calculate the score using the scenario's input.
			got := Score(tt.input)

			// Compare the calculated score against the expected score.
			if got != tt.want {

				// Fail the test if the values don't match.
				t.Fatalf(
					"Score() = %d, want %d",
					got,
					tt.want,
				)
			}
		})
	}
}
