package main

import "fmt"

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := make([]int32, 0)
	for _, char := range s {
		switch {
		case len(stack) == 0:
			stack = append(stack, char)
		case !isMatch(stack[len(stack)-1], char):
			stack = append(stack, char)
		case isMatch(stack[len(stack)-1], char):
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func isMatch(left, right int32) bool {
	switch {
	case left == '(' && right == ')':
		return true
	case left == '[' && right == ']':
		return true
	case left == '{' && right == '}':
		return true
	default:
		return false
	}
}
func main() {
	res := isValid("()")
	fmt.Print(res)
	//fmt.Printf("%T", '(')
}
