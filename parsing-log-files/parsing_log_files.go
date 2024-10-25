package parsinglogfiles

import "regexp"

// Task 1: Check if a log line is valid
func IsValidLine(line string) bool {
	// Regex to check if the line starts with one of the valid prefixes
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(line)
}

// Task 2: Split the log line by custom separators
func SplitLogLine(line string) []string {
	// Regex to match separators like "<*>", "<~~~>", etc.
	re := regexp.MustCompile(`<[-=~*]*>`)
	// Split the log line by the regex pattern
	return re.Split(line, -1)
}

// Task 3: Count the number of lines containing "password" in quoted text
func CountQuotedPasswords(lines []string) int {
	// Regex to match "password" inside double quotes, case insensitive
	re := regexp.MustCompile(`"(?i).*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

// Task 4: Remove occurrences of "end-of-line" followed by digits
func RemoveEndOfLineText(line string) string {
	// Regex to match "end-of-line" followed by digits
	re := regexp.MustCompile(`end-of-line\d+`)
	// Replace occurrences with an empty string
	return re.ReplaceAllString(line, "")
}

func TagWithUserName(lines []string) []string {
	// Regex to find "User " followed by the username (non-whitespace)
	re := regexp.MustCompile(`User\s+(\S+)`)
	result := []string{}
	for _, line := range lines {
		if matches := re.FindStringSubmatch(line); matches != nil {
			username := matches[1] // Captured username
			taggedLine := "[USR] " + username + " " + line
			result = append(result, taggedLine)
		} else {
			result = append(result, line)
		}
	}
	return result
}
