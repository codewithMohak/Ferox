package fingerprint

import (
	"bytes"
	"strings"
)

var blockIndicators = [][]byte{
	[]byte("access denied"),
	[]byte("request blocked"),
	[]byte("forbidden"),
	[]byte("security check"),
	[]byte("blocked by security"),
}

func IsBlockPage(statusCode int, body []byte) bool {
	if len(body) == 0 {
		return false
	}
	normalizedBody := strings.ToLower(string(body))
	for _, indicator := range blockIndicators {
		if bytes.Contains(
			[]byte(normalizedBody),
			bytes.ToLower(indicator),
		) {

			// An indicator was found, so classify this response as a possible block page.
			return true
		}
	}

	if statusCode == 403 {
		return true
	}
	return false	
}
