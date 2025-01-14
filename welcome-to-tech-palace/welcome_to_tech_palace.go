package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	repeatStars := strings.Repeat("*", numStarsPerLine)

	return fmt.Sprintf(`
    %s
    %s
    %s
    `, repeatStars, welcomeMsg, repeatStars)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.ReplaceAll(oldMsg, "*", "")
}
