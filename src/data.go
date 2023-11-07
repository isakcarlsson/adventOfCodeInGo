package src

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadData(year int, day int) string {
	data, err := os.ReadFile(fmt.Sprintf("src/%d/day%d/input.txt", year, day))
	check(err)

	return string(data)
}
