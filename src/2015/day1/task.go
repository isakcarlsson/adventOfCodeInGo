package main

import (
	"fmt"

	"github.com/isakcarlsson/adventOfCodeInGo/src"
)

func task1(data string) int {
	floor := 0

	for _, v := range data {
		if v == '(' {
			floor++
		} else if v == ')' {
			floor--
		}
	}

	return floor
}

func task2(data string) int {
	floor := 0

	for i, v := range data {
		if v == '(' {
			floor++
		} else if v == ')' {
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}

	return -1
}

func main() {
	data := src.ReadData(2015, 1)

	fmt.Println(task1(data))
	fmt.Println(task2(data))
}
