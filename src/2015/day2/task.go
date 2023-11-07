package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/isakcarlsson/adventOfCodeInGo/src"
)

func task1(presents [][]int) int {
	totalArea := 0
	for _, present := range presents {
		l, w, h := present[0], present[1], present[2]

		s1 := l * w
		s2 := w * h
		s3 := h * l

		area := 2*s1 + 2*s2 + 2*s3 + int(math.Min(float64(s1), math.Min(float64(s2), float64(s3))))

		totalArea += area
	}

	return totalArea
}

func task2(presents [][]int) int {
	totalLength := 0

	for _, present := range presents {
		l, w, h := present[0], present[1], present[2]

		side1 := 2 * (l + w)
		side2 := 2 * (w + h)
		side3 := 2 * (h + l)

		length := int(math.Min(float64(side1), math.Min(float64(side2), float64(side3)))) + (l * w * h)

		totalLength += length
	}

	return totalLength
}

func main() {
	data := src.ReadData(2015, 2)
	lines := strings.Split(data, "\n")
	presents := make([][]int, len(lines))

	for i, line := range lines {
		presents[i] = make([]int, 3)

		dimensions := strings.Split(line, "x")

		s, err := strconv.Atoi(dimensions[0])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		} else {
			presents[i][0] = s
		}

		// we can also convert negative numbers
		s1, err1 := strconv.Atoi(dimensions[1])
		if err1 != nil {
			fmt.Println("Can't convert this to an int!")
		} else {
			presents[i][1] = s1
		}

		// now, we will try to convert a not valid
		// string and handle the error
		s2, err2 := strconv.Atoi(dimensions[2])
		if err2 != nil {
			fmt.Println("Can't convert this to an int!")
		} else {
			presents[i][2] = s2
		}
	}

	fmt.Println(task1(presents))
	fmt.Println(task2(presents))
}
