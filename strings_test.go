package friendly

import (
	"fmt"
	"testing"
)

func stringStringTest(t *testing.T, f func(string) string, fname string, input []string, want []string) {
	for i, w := range input {
		t.Run(fmt.Sprintf("%s=%d", fname, i), func(t *testing.T) {
			got := f(w)
			if got != want[i] {
				t.Fatalf("want %s, got %s", want[i], got)
			}
		})
	}
}

func TestCapitalizeWord(t *testing.T) {
	input := []string{"hello", "world", "Akshan", "computer", "34tz"}
	want := []string{"Hello", "World", "Akshan", "Computer", "34tz"}

	stringStringTest(t, CapitalizeWord, "CapitalizeWord", input, want)
}

func TestGetFullNameFromMailAddress(t *testing.T) {
	input := []string{"john.doe@example.com", "farrukh.alhem@gmail.com", "user87236@hotmail.com", "PeterBrown@outlook.com"}
	want := []string{"John Doe", "Farrukh Alhem", "User87236", "PeterBrown"}

	stringStringTest(t, GetFullNameFromMailAddress, "GetFullNameFromMailAddress", input, want)
}

func TestSafeCharsOnly(t *testing.T) {
	strs := []string{"Hello", "Good afternoon", ".e", "#`", "123"}
	allowed := []byte("abcdefghjklmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ1234567890")
	want := []bool{true, false, false, false, true}

	for i := range strs {
		t.Run(fmt.Sprintf("SafeCharsOnly=%d", i), func(t *testing.T) {
			got := SafeCharsOnly(strs[i], allowed)
			if got != want[i] {
				t.Fatalf("want %v, got %v", want[i], got)
			}
		})
	}
}

func TestSemVersionGreater(t *testing.T) {
	input := [][]string{
		{"1.2.3", "1.2.3"},
		{"1.0.1", "1.0.0"},
		{"4.0.3", "4.1.8"},
		{"2.0.0", "1.9.9"},
	}
	want := []bool{false, true, false, true}

	for i, a := range input {
		t.Run(fmt.Sprintf("SemVersionGreater=%d", i), func(t *testing.T) {
			got := SemVersionGreater(a[0], a[1])
			if got != want[i] {
				t.Fatalf("want %v, got %v", want[i], got)
			}
		})
	}
}
