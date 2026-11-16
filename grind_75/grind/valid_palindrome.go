package grind

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"unicode"
)

func isAlphaNumeric(c rune) bool {
	// Check if the byte value falls within the range of alphanumeric characters
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func IsPalindrome(s string) bool {
	s = strings.TrimSpace(s)
	// NOTE Explanation: s is an empty string "" after removing non-alphanumeric characters.
	// Since an empty string reads the same forward and backward, it is a palindrome.
	if len(s) < 1 {
		return true
	}

	var cleanString []rune
	// var reversedString []rune
	for _, char := range s {

		isAlphaNumeric := isAlphaNumeric(char)
		if isAlphaNumeric {
			cleanString = append(cleanString, char)
		}

		fmt.Println(char)
	}
	if len(cleanString) < 1 {
		return true
	}
	reversedString := make([]rune, len(cleanString))

	// NOTE copy https://yuminlee2.medium.com/golang-deep-and-shallow-copy-a-slice-7dbcd9eb044a
	// BUG this is not a deep copy!
	// 	// BUT Gotcha : I don't think `copy` actually deep-copies arrays. For example, if you `copy` an array of pointers, the source and destination array elements will point at the same memory addresses. A deep-copy would copy the referenced values to new memory addresses.
	// Shallow copy: https://goplay.tools/snippet/1OI45Flf_Vq
	// Deep copy: https://goplay.tools/snippet/j7dbA2geNIP
	// https://goplay.tools/snippet/65I9UCF6xlY
	// https://goplay.tools/snippet/j7dbA2geNIP

	copy(reversedString, cleanString)

	slices.Reverse(reversedString)

	fmt.Println(string(cleanString))
	fmt.Println(string(reversedString))

	return reflect.DeepEqual(strings.ToLower(string(cleanString)), strings.ToLower(string(reversedString)))
}

// NOTE I had feeling here is not idomatic approach. accorinding to AI

func isPalindrome(s string) bool {
	// Use strings.Builder to avoid repeatedly concatenating characters, and reverse runes so Unicode characters aren’t split:
	// https://medium.com/@raisulislamaupi/understanding-string-builder-reader-in-go-1486d75942ba

	var clean strings.Builder
	for _, r := range strings.ToLower(s) {
		// https://pkg.go.dev/unicode
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			// TODO: look at Builder methods and packeg tomorrow. I think I can learn a lot from it's source code
			// // WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
			// // It returns the number of bytes written and a nil error.
			// func (b *Builder) WriteRune(r rune) (int, error) {
			// 	b.copyCheck()
			// 	n := len(b.buf)
			// 	b.buf = utf8.AppendRune(b.buf, r)
			// 	return len(b.buf) - n, nil
			// }

			clean.WriteRune(r)
		}
	}

	cleanString := clean.String()
	runes := []rune(cleanString)

	var reversed strings.Builder
	for i := len(runes) - 1; i >= 0; i-- {
		reversed.WriteRune(runes[i])
	}

	return cleanString == reversed.String()
	// This uses strings and unicode from the standard library.
}
