package logs

import (
	"strings"
	"unicode/utf8"
)

var myRunes = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		if value, ok := myRunes[char]; ok {
			return value
		}
	}

	return "default"
}

// Replace replaces all occurrences of the corrupted character with the original value in the log.
func Replace(log string, corrupted, replacement rune) string {
	// Convert corrupted and replacement runes to strings
	corruptedStr := string(corrupted)
	replacementStr := string(replacement)

	// Use strings.ReplaceAll to replace all instances of the corrupted character
	return strings.ReplaceAll(log, corruptedStr, replacementStr)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log)
}
