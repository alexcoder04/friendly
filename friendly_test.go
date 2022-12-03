package friendly

import (
	"fmt"
	"testing"
)

func TestArrayContains(t *testing.T) {
	inputArrInt := []int{1, 2, 3, 4, 5, 6, 7}
	inputElInt := []int{1, 3, 89}
	want := []bool{true, true, false}
	for i, n := range inputElInt {
		t.Run(fmt.Sprintf("ArrayContains=%d, int", i), func(t *testing.T) {
			got := ArrayContains(inputArrInt, n)
			if got != want[i] {
				t.Fatalf("got %v, want %v", got, want[i])
			}
		})
	}

	inputArrStr := []string{"hello", "world", "I", "am", "Go"}
	inputElStr := []string{"hello", "go", "pythoN"}
	want = []bool{true, false, false}
	for i, n := range inputElStr {
		t.Run(fmt.Sprintf("ArrayContains=%d, string", i), func(t *testing.T) {
			got := ArrayContains(inputArrStr, n)
			if got != want[i] {
				t.Fatalf("got %v, want %v", got, want[i])
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"1", true},
		{"26", true},
		{"1987", true},
		{"873.298", false},
		{"2873.iud", false},
		{"uebdu", false},
		{"87hvi87", false},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("IsInt=%d", i), func(t *testing.T) {
			got := IsInt(tc.input)
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}
