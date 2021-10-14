package ch01

import (
	"strconv"
	"strings"
)

/**
 * Given a string, compress it into a readable format
 */

func stringCompression(s string) string {
	if len(s) == 0 {
		return ""
	}

	letter := s[0]
	count := 1
	builder := strings.Builder{}
	for i := 1; i < len(s); i++ {
		if letter == ' ' {
			letter = s[i]
			continue
		}

		if s[i] == letter {
			count++
			continue
		}

		builder.WriteString(string(letter) + strconv.Itoa(count))
		letter = s[i]
		count = 1
	}

	if letter != ' ' {
		builder.WriteString(string(letter) + strconv.Itoa(count))
	}

	return builder.String()
}