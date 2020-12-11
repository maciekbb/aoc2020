package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day10/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var numbers []int

	numbers = append(numbers, 0)

	for _, row := range rows {
		voltage, err := strconv.Atoi(row)
		check(err)
		numbers = append(numbers, voltage)
	}

	sort.Ints(numbers)

	d1 := 0
	d3 := 1

	for i := 0; i < len(numbers)-1; i++ {
		d := numbers[i+1] - numbers[i]

		if d == 1 {
			d1++
		}

		if d == 3 {
			d3++
		}
	}

	fmt.Printf("Solution is %d %d %d\n", d1, d3, d1*d3)
}
