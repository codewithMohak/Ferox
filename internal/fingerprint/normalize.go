package fingerprint

import (
	"regexp"
	"strings"
)

// UUID pattern used to identify common UUID values in responses.
var uuidPattern = regexp.MustCompile(
	`(?i)\b[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}\b`,
)

// Normalize removes common transient values from a response body
func Normalize(body []byte) []byte {
	text := string(body)
	text = uuidPattern.ReplaceAllString(text, "<UUID>")
	text = strings.TrimSpace(text)
	return []byte(text)
}
