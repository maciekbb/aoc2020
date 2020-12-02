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

type rule struct {
	min  int
	max  int
	char rune
}

func (r rule) validate(pswd string) bool {
	counter := 0

	for _, ch := range pswd {
		if ch == r.char {
			counter++
		}
	}

	return r.min <= counter && counter <= r.max
}

func toInt(s string) int {
	v, err := strconv.Atoi(s)
	check(err)
	return v
}

func main() {
	dat, err := ioutil.ReadFile("./day2/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	items := strings.Split(content, "\n")

	validCount := 0

	for _, item := range items {
		parsed := strings.Split(item, " ")
		char := []rune(parsed[1])
		pswd := parsed[2]

		bounds := strings.Split(parsed[0], "-")

		rule := &rule{min: toInt(bounds[0]), max: toInt(bounds[1]), char: char[0]}
		isValid := rule.validate(pswd)
		if isValid {
			validCount++
		}

	}

	fmt.Printf("Solution is %d\n", validCount)

}
