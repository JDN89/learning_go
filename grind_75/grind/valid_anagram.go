package grind

import (
	"slices"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sRunes := []rune(s)
	slices.Sort(sRunes)
	tRunes := []rune(t)
	slices.Sort(tRunes)

	return slices.Equal(sRunes, tRunes)
}
