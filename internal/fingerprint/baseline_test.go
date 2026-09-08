package fingerprint

import "testing"

func TestBaseLine_Add(t *testing.T) {
	baseline := &Baseline{Counts: make(map[string]int)}

	baseline.Add("A")
	baseline.Add("A")
	baseline.Add("B")

	if baseline.Counts["A"] != 2 {
		t.Fatalf(
			"count for A = %d, want 2",
			baseline.Counts["A"],
		)
	}

	if baseline.Counts["B"] != 1 {
		t.Fatalf(
			"count for B = %d, want 1",
			baseline.Counts["B"],
		)
	}
}

func TestBaseline_Dominant(t *testing.T) {
	baseline := &Baseline{Counts: make(map[string]int)}

	baseline.Add("A")
	baseline.Add("A")
	baseline.Add("A")

	baseline.Add("B")

	got := baseline.Dominant()

	if got != "A" {
		t.Fatalf(
			"dominant = %q, want %q",
			got,
			"A",
		)
	}
}
