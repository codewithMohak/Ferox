package engine // Places this test inside the engine package.

import "testing" // Provides Go's testing framework.

// TestBuildURL verifies that wordlist entries are correctly appended to a target.
func TestBuildURL(t *testing.T) {

	// Define table-driven URL generation test cases.
	tests := []struct {
		name string // Provides a readable test name.
		base string // Stores the base target URL.
		word string // Stores the wordlist entry.
		want string // Stores the expected final URL.
	}{
		{
			name: "normal path",
			base: "https://example.com",
			word: "admin",
			want: "https://example.com/admin",
		},
		{
			name: "base URL with trailing slash",
			base: "https://example.com/",
			word: "login",
			want: "https://example.com/login",
		},
		{
			name: "word with leading slash",
			base: "https://example.com",
			word: "/api",
			want: "https://example.com/api",
		},
	}

	// Run every test case independently.
	for _, tt := range tests {

		// Create a named sub-test.
		t.Run(tt.name, func(t *testing.T) {

			// Build the target URL.
			got, err := BuildURL(tt.base, tt.word)

			// Fail if URL construction returned an error.
			if err != nil {
				t.Fatal(err)
			}

			// Compare the generated URL with the expected URL.
			if got != tt.want {
				t.Fatalf(
					"got %q, want %q",
					got,
					tt.want,
				)
			}
		})
	}
}
