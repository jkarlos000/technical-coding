package main

import (
	"fmt"
	"strconv"
	"strings"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	if s.Empty() {
		return s, -1
	}
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Empty() bool {
	if s == nil || len(s) == 0 {
		return true
	}
	return false
}

func main() {
	// x := Solution("13 DUP 4 POP 5 DUP + DUP + -")
	// x := Solution("4 5 6 - 7 +")
	x := Solution("5 6 + -")
	fmt.Println(x)
}

func Solution(S string) int {
	// write your code in Go 1.13.8
	stackInt := make(stack, 0)
	operation := strings.Fields(S)
	var a, b int
	for i := 0; i < len(operation); i++ {
		switch operation[i] {
		case "DUP":
			if stackInt.Empty() {
				break
			}
			stackInt, a = stackInt.Pop()
			stackInt = stackInt.Push(a)
			stackInt = stackInt.Push(a)
			break
		case "POP":
			if stackInt.Empty() {
				break
			}
			stackInt, _ = stackInt.Pop()
		case "+":
			stackInt, a = stackInt.Pop()
			if stackInt.Empty() {
				break
			}
			stackInt, b = stackInt.Pop()
			stackInt = stackInt.Push(a + b)
			break
		case "-":
			stackInt, a = stackInt.Pop()
			if stackInt.Empty() {
				break
			}
			stackInt, b = stackInt.Pop()
			stackInt = stackInt.Push(a - b)
			break
		default:
			n, err := strconv.ParseInt(operation[i], 10, 20)
			if err != nil {
				break
			}
			stackInt = stackInt.Push(int(n))
		}
	}
	_, last := stackInt.Pop()
	return last
}
