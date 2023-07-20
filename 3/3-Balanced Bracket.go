package main

import (
	"fmt"
)

func isBracket(char rune) bool {
	return char == '{' || char == '}' || char == '[' || char == ']' || char == '(' || char == ')'
}

func isOpeningBracket(char rune) bool {
	return char == '{' || char == '[' || char == '('
}

func isMatchingPair(open, close rune) bool {
	switch open {
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	case '(':
		return close == ')'
	}
	return false
}

func balancedBracket(brackets string) string {
	var stack []rune

	for _, char := range brackets {
		if isBracket(char) {
			if isOpeningBracket(char) {
				stack = append(stack, char)
			} else {
				if len(stack) == 0 || !isMatchingPair(stack[len(stack)-1], char) {
					return "NO"
				}
				stack = stack[:len(stack)-1]
			}
		} else {
			return "NO"
		}
	}

	if len(stack) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	var brackets string
	fmt.Scan(&brackets)

	fmt.Println(balancedBracket(brackets))
}
