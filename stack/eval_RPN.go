package stack

import (
	"strconv"
)

/*
You are given an array of strings tokens that represents a valid arithmetic expression in Reverse Polish Notation.

Return the integer that represents the evaluation of the expression.

The operands may be integers or the results of other operations.
The operators include '+', '-', '*', and '/'.
Assume that division between integers always truncates toward zero.
Example 1:

Input: tokens = ["1","2","+","3","*","4","-"]

Output: 5

Explanation: ((1 + 2) * 3) - 4 = 5
Constraints:

1 <= tokens.length <= 1000.
tokens[i] is "+", "-", "*", or "/", or a string representing an integer in the range [-100, 100].
*/

func evalRPN(tokens []string) int {
	res := 0

	eval := func(op string, a, b int) {
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		default:
			panic("unknown operator")
		}
	}

	var stack []int
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			n := len(stack)
			a := stack[n-2]
			b := stack[n-1]
			stack = stack[:n-2]
			eval(token, a, b)
			stack = append(stack, res)
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[len(stack)-1]
}
