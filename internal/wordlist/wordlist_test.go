package wordlist // Places this test inside the wordlist package.

import (
	"os"      // Provides temporary file functionality.
	"testing" // Provides Go's testing framework.
)

// TestLoad verifies that the wordlist loader correctly processes a file.
func TestLoad(t *testing.T) {

	// Create temporary wordlist content for the test.
	content := `
# This is a comment.
admin

login
  api
`

	// Create a temporary file for the test.
	file, err := os.CreateTemp("", "ferox-wordlist-*.txt")

	// Stop the test if the temporary file could not be created.
	if err != nil {
		t.Fatal(err)
	}

	// Remove the temporary file after the test finishes.
	defer os.Remove(file.Name())

	// Write the test content into the temporary file.
	if _, err := file.WriteString(content); err != nil {
		t.Fatal(err)
	}

	// Close the file before loading it.
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	// Load the temporary wordlist.
	got, err := Load(file.Name())

	// Stop the test if loading failed.
	if err != nil {
		t.Fatal(err)
	}

	// Define the entries we expect after cleaning the file.
	want := []string{
		"admin",
		"login",
		"api",
	}

	// Check that both slices contain the same number of entries.
	if len(got) != len(want) {
		t.Fatalf(
			"got %d words, want %d",
			len(got),
			len(want),
		)
	}

	// Compare every loaded word with the expected word.
	for i := range want {

		// Fail if the current word doesn't match.
		if got[i] != want[i] {
			t.Fatalf(
				"word %d = %q, want %q",
				i,
				got[i],
				want[i],
			)
		}
	}
}
