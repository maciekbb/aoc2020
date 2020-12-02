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
	a    int
	b    int
	char rune
}

func (r rule) validate(pswd []rune) bool {

	p := pswd[r.a-1] == r.char
	q := pswd[r.b-1] == r.char

	return (p || q) && !(p && q)
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
		pswd := []rune(parsed[2])

		bounds := strings.Split(parsed[0], "-")

		rule := &rule{a: toInt(bounds[0]), b: toInt(bounds[1]), char: char[0]}
		isValid := rule.validate(pswd)
		if isValid {
			validCount++
		}

	}

	fmt.Printf("Solution is %d\n", validCount)

}
