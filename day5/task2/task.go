package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convert(s []rune, zero, one rune) int {
	result := 0

	for _, ch := range s {
		result *= 2
		if ch == one {
			result++
		}
	}

	return result
}

func main() {
	dat, err := ioutil.ReadFile("./day5/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	codes := strings.Split(content, "\n")

	seatids := make(map[int]bool)

	m := 0
	for _, code := range codes {
		row := convert([]rune(code)[0:7], 'F', 'B')
		column := convert([]rune(code)[7:10], 'L', 'R')

		r := row*8 + column
		seatids[r] = true

		if r > m {
			m = r
		}
	}

	for i := 0; i <= m; i++ {
		if _, ok := seatids[i]; !ok {
			fmt.Printf("Solution is: %d\n", i)
		}
	}

}
