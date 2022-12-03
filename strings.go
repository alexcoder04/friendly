package friendly

import (
	"strconv"
	"strings"
	"unicode"
)

// Capitalizes a first (makes first letter capital).
func CapitalizeWord(word string) string {
	return string(append([]rune{unicode.ToUpper([]rune(word)[0])}, []rune(word)[1:]...))
}

// Capitalizes every part of between periods of the email address before the @-sign.
// Joins them with " ".
// Example: "john.doe@example.com" -> "John Doe".
func GetFullNameFromMailAddress(address string) string {
	parts := strings.Split(strings.Split(address, "@")[0], ".")
	res := []string{}
	for _, p := range parts {
		if len(p) < 2 {
			continue
		}
		res = append(res, CapitalizeWord(p))
	}
	return strings.Join(res, " ")
}

// Checks if first semantic version is greater than another.
func SemVersionGreater(v1 string, v2 string) bool {
	v1a := strings.Split(v1, ".")
	v2a := strings.Split(v2, ".")

	for i := 0; i < len(v1a); i++ {
		num1, err := strconv.Atoi(v1a[i])
		if err != nil {
			num1 = 0
		}
		num2, err := strconv.Atoi(v2a[i])
		if err != nil {
			num2 = 0
		}
		if num1 == num2 {
			continue
		}
		if num1 > num2 {
			return true
		}
		return false
	}
	return false
}
