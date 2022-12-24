package friendly

import (
	"fmt"
	"os"
	"strconv"
)

// Checks if an element is present in an array.
// Can be used with any type thanks to generics.
func ArrayContains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// Checks whether a string represents a valid integer.
func IsInt(num string) bool {
	_, err := strconv.Atoi(num)
	return err == nil
}

// Prints the provided error message and exits immediately.
// `Fatal error: <message>`
func Die(message string, args ...any) {
	fmt.Printf("Fatal error: %s\n", fmt.Sprintf(message, args...))
	os.Exit(1)
}

// Prints a warning message.
// `Warning: <message>`
func Warn(message string, args ...any) {
	fmt.Printf("Warning: %s\n", fmt.Sprintf(message, args...))
}
