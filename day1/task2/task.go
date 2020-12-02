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

func findTwoMatching(numbers []string, target int) (int, bool) {
	lookup := make(map[int]bool)

	for _, expense := range numbers {
		expenseAsInt, err := strconv.Atoi(expense)
		check(err)

		if _, ok := lookup[expenseAsInt]; ok {
			return expenseAsInt * (target - expenseAsInt), true
		}

		lookup[target-expenseAsInt] = true
	}

	return 0, false
}

func main() {
	dat, err := ioutil.ReadFile("./day1/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	numbers := strings.Split(content, "\n")

	for _, first := range numbers {
		expenseAsInt, err := strconv.Atoi(first)
		check(err)

		if match, ok := findTwoMatching(numbers, 2020-expenseAsInt); ok {
			fmt.Printf("Solution is: %d\n", expenseAsInt*match)
			break
		}
	}

}
