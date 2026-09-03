package fingerprint

import "testing"

func TestNormalize_UUID(t *testing.T) {
	input := []byte(
		`User ID: 550e8400-e29b-41d4-a716-446655440000`,
	)

	got := string(Normalize(input))
	want := `User ID: <UUID>`

	if got != want {
		t.Fatalf(
			"got %q, want %q",
			got,
			want,
		)
	}
}

func TestNormalize_Whitespace(t *testing.T) {
	input := []byte("   hello world   ")

	got := string(Normalize(input))
	want := "hello world"

	if got !=want {
		t.Fatalf(
			"got %q, want %q",
			got,
			want,
		)
	}
}