package engine // Defines this file as part of the engine package.

import (
	"fmt"     // Provides formatted string creation.
	"net/url" // Provides URL parsing and validation.
	"strings" // Provides string manipulation functions.
)

// BuildURL creates a target URL by combining a base URL and a wordlist entry.
func BuildURL(baseURL string, word string) (string, error) {

	// Parse the supplied base URL.
	parsed, err := url.Parse(baseURL)

	// Check whether the base URL could be parsed.
	if err != nil {

		// Return an error describing the invalid URL.
		return "", fmt.Errorf("parse base URL: %w", err)
	}

	// Make sure the URL includes a scheme such as http or https.
	if parsed.Scheme != "http" && parsed.Scheme != "https" {

		// Reject URLs without a supported HTTP scheme.
		return "", fmt.Errorf("unsupported URL scheme: %q", parsed.Scheme)
	}

	// Make sure the URL contains a hostname.
	if parsed.Host == "" {

		// Reject URLs that don't specify a target host.
		return "", fmt.Errorf("URL has no host")
	}

	// Remove any trailing slash from the base path.
	basePath := strings.TrimRight(parsed.Path, "/")

	// Remove any leading slash from the wordlist entry.
	word = strings.TrimLeft(word, "/")

	// Construct the final path.
	parsed.Path = basePath + "/" + word

	// Return the complete URL.
	return parsed.String(), nil
}
