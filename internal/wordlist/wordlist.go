package wordlist

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Load(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Open Wordlist: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())

		if word == "" {
			continue
		}

		if strings.HasPrefix(word, "#") {
			continue
		}
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read wordlist: %w", err)
	}
	return words, nil
}
