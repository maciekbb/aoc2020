package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day1/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	numbers := strings.Split(content, "\n")

	lookup := make(map[int]bool)

	for _, expense := range numbers {
		expenseAsInt, err := strconv.Atoi(expense)
		check(err)

		if _, ok := lookup[expenseAsInt]; ok {
			fmt.Printf("Solution is: %d\n", expenseAsInt*(2020-expenseAsInt))
			break
		}

		lookup[2020-expenseAsInt] = true
	}

}
