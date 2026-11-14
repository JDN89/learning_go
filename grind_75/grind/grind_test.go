package grind_test

import (
	"slices"
	"testing"

	"grind_75/grind"
)

func TestTwoSum(t *testing.T) {
	got := grind.TwoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if !slices.Equal(got, want) {
		t.Errorf(" TwoSum result got %d != want %d", got, want)
	}
}

func TestIsValid(t *testing.T) {
	input := "()[]{}"
	got := grind.IsValid(input)

	want := true
	// NOTE print boolean https://pkg.go.dev/fmt
	if got != want {
		t.Errorf("wanted %t , but got %t ", got, want)
	}
}
