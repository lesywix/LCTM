package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, len(temperatures))
	for i, v := range temperatures {
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 71, 72, 76, 73}))
}
