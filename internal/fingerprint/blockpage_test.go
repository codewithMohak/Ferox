package fingerprint // Places this test in the fingerprint package.

import "testing" // Provides Go's testing framework.

// TestIsBlockPage tests responses that should be recognized as block pages.
func TestIsBlockPage(t *testing.T) {

	// Define test cases for different response scenarios.
	tests := []struct {
		name       string // Gives each test case a readable name.
		statusCode int    // Represents the HTTP status returned by the server.
		body       []byte // Represents the HTTP response body.
		want       bool   // Represents whether we expect the response to be classified as a block page.
	}{
		{
			name:       "access denied page",    // Tests a common access-denied message.
			statusCode: 403,                     // Uses the HTTP 403 status code.
			body:       []byte("Access Denied"), // Provides a generic security-block message.
			want:       true,                    // This should be recognized as a block page.
		},
		{
			name:       "request blocked page",    // Tests another common block message.
			statusCode: 403,                       // Uses the HTTP 403 status code.
			body:       []byte("Request Blocked"), // Provides a generic block message.
			want:       true,                      // This should be recognized as a block page.
		},
		{
			name:       "normal page",                               // Tests a response that doesn't look like a block page.
			statusCode: 200,                                         // Uses a successful HTTP status.
			body:       []byte("<html><body>Welcome</body></html>"), // Provides normal application content.
			want:       false,                                       // This should not be classified as a block page.
		},
		{
			name:       "empty response", // Tests a response without a body.
			statusCode: 200,              // Uses a successful HTTP status.
			body:       nil,              // Represents an empty response body.
			want:       false,            // An empty response cannot contain a block-page indicator.
		},
	}

	// Execute every test case independently.
	for _, tt := range tests {

		// Create a named sub-test for the current test case.
		t.Run(tt.name, func(t *testing.T) {

			// Run the block-page detector against the test response.
			got := IsBlockPage(tt.statusCode, tt.body)

			// Compare the actual result with the expected result.
			if got != tt.want {

				// Fail the test and display the unexpected result.
				t.Fatalf(
					"IsBlockPage() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
