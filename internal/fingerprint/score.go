package fingerprint

// ScoreInput contains the signals used by the anomaly scorer.
type ScoreInput struct {
	StatusCode     int
	Fingerprint    string
	Baseline       string
	IsBlockPage    bool
	Length         int64
	BaselineLength int64
}

func Score(input ScoreInput) int {
	score := 0
	if input.Fingerprint != "" &&
		input.Baseline != "" &&
		input.Fingerprint != input.Baseline {
		score += 50
	}

	if input.IsBlockPage {
		score -= 30
	}

	if input.StatusCode >= 200 &&
		input.StatusCode < 300 {

		// Add points for a successful response.
		score += 20
	}
	lengthDifference := input.Length - input.BaselineLength

	if lengthDifference < 0 {

		// Negate the value so we can compare the magnitude of the difference.
		lengthDifference = -lengthDifference
	}
	if input.BaselineLength > 0 &&
		lengthDifference > input.BaselineLength/2 {

		// Add points because the response size differs substantially from the baseline.
		score += 10
	}
	if score < 0 {
		score = 0
	}
	return score
}
