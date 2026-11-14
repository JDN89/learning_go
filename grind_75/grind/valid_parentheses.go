package grind

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
//     Open brackets must be closed by the same type of brackets.
//     Open brackets must be closed in the correct order.
//     Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// Input: s = "()"
//
// Output: true
//
// Example 2:
//
// Input: s = "()[]{}"
//
// Output: true
//
// Example 3:
//
// Input: s = "(]"
//
// Output: false
//
// Example 4:
//
// Input: s = "([])"
//
// Output: true
//
// Example 5:
//
// Input: s = "([)]"
//
// Output: false
//
// Constraints:
//
//     1 <= s.length <= 104
//     s consists of parentheses only '()[]{}'.

// SOLUTION
// ( [ { chars you push on the stack
// encounter '}' pop from stack and see that char is '{'
// encounter ']' pop from stack and see that char is '['
// encounter ')' pop from stack and see that char is '('

func IsValid(s string) bool {
	stack := []rune{}

	for idx, char := range s {
		// index 0 and closing, false -> fast return
		if (idx == 0 && char == '}') || (idx == 0 && char == ']') || (idx == 0 && char == ')') {
			return false
		}

		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		}

		if char == ')' || char == ']' || char == '}' {
			// Encoutner closing but no elements left on stack -> no match == FALSE
			if len(stack) == 0 {
				return false
			}

			if len(stack) > 0 {

				last := stack[len(stack)-1]

				if char == ')' && last != '(' {
					return false
				}
				if char == ']' && last != '[' {
					return false
				}
				if char == '}' && last != '{' {
					return false
				}

				stack = stack[:len(stack)-1]
			}
		}

	}
	return len(stack) == 0
}
