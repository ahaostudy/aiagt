package utils

import "unicode/utf8"

func TruncateString(s string, maxLength int) string {
	if utf8.RuneCountInString(s) <= maxLength {
		return s
	}

	truncated := []rune(s)[:maxLength]
	return string(truncated)
}
