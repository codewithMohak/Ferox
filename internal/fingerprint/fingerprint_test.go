package fingerprint // Defines this test as part of the fingerprint package.

import "testing" // Provides Go's testing framework.

// TestFingerprint verifies that identical bodies produce identical fingerprints.
func TestFingerprint(t *testing.T) {

	// Define the first response body.
	bodyA := []byte("hello world")

	// Define another response body with the same content.
	bodyB := []byte("hello world")

	// Generate a fingerprint for the first body.
	fingerprintA := Fingerprint(bodyA)

	// Generate a fingerprint for the second body.
	fingerprintB := Fingerprint(bodyB)

	// Identical bodies should produce identical fingerprints.
	if fingerprintA != fingerprintB {
		t.Fatalf(
			"fingerprints differ: %q != %q",
			fingerprintA,
			fingerprintB,
		)
	}
}

// TestFingerprint_DifferentBodies verifies different bodies produce different fingerprints.
func TestFingerprint_DifferentBodies(t *testing.T) {

	// Define the first response body.
	bodyA := []byte("hello world")

	// Define a different response body.
	bodyB := []byte("hello ferox")

	// Generate a fingerprint for the first body.
	fingerprintA := Fingerprint(bodyA)

	// Generate a fingerprint for the second body.
	fingerprintB := Fingerprint(bodyB)

	// Different bodies should produce different fingerprints.
	if fingerprintA == fingerprintB {
		t.Fatalf(
			"expected different fingerprints, got %q",
			fingerprintA,
		)
	}
}

// TestFingerprint_NormalizesUUID verifies that different UUIDs produce the same fingerprint.
func TestFingerprint_NormalizesUUID(t *testing.T) {

	// Create the first response containing a UUID.
	bodyA := []byte(
		`Request ID: 550e8400-e29b-41d4-a716-446655440000`,
	)

	// Create another response containing a different UUID.
	bodyB := []byte(
		`Request ID: 123e4567-e89b-42d3-a456-426614174000`,
	)

	// Generate a fingerprint for the first response.
	fingerprintA := Fingerprint(bodyA)

	// Generate a fingerprint for the second response.
	fingerprintB := Fingerprint(bodyB)

	// Different UUIDs should become the same <UUID> value after normalization.
	if fingerprintA != fingerprintB {
		t.Fatalf(
			"fingerprints differ: %q != %q",
			fingerprintA,
			fingerprintB,
		)
	}
}
