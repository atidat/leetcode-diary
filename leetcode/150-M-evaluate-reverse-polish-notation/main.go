package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	calc := make(map[string]bool, 4)
	calc["+"] = true
	calc["-"] = true
	calc["*"] = true
	calc["/"] = true

	for k := range tokens {
		if _, ok := calc[tokens[k]]; !ok {
			intoVal, _ := strconv.Atoi(tokens[k])
			stack = append(stack, intoVal)
		} else {
			new := 0
			if tokens[k] == "+" {
				new = stack[len(stack)-2] + stack[len(stack)-1]
			} else if tokens[k] == "-" {
				new = stack[len(stack)-2] - stack[len(stack)-1]
			} else if tokens[k] == "*" {
				new = stack[len(stack)-2] * stack[len(stack)-1]
			} else {
				new = stack[len(stack)-2] / stack[len(stack)-1]
			}
			stack = append(stack[0:len(stack)-2], new)
		}
		fmt.Println(stack)
	}
	fmt.Println(stack[0])
	return stack[0]
}

func main() {

}
