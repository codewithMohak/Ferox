package fingerprint

import "testing"

// TestFingerprint verifies that identical response bodies produce identical fingerprints.

func TestFingerprint(t *testing.T) {
	//Define the first response body
	bodyA := []byte("Hello, World!")

	bodyB := []byte("Hello, World!")

	fingerprintA := Fingerprint(bodyA)
	fingerprintB := Fingerprint(bodyB)

	if fingerprintA != fingerprintB {
		t.Fatalf(
			"fingerprint differ: %q != %q",
			fingerprintA,
			fingerprintB,
		)
	}
}
