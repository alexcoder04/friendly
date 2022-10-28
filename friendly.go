package friendly

import (
	"strconv"
	"strings"
)

func ArrayContains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func IsInt(num string) bool {
	_, err := strconv.Atoi(num)
	return err == nil
}

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
