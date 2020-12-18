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

type Operation interface {
	exec(v int) int
}

type AddOperation struct {
	value int
}

func (p AddOperation) exec(v int) int {
	return p.value + v
}

type MultiplyOperation struct {
	value int
}

func (p MultiplyOperation) exec(v int) int {
	return p.value * v
}

func tokenize(s string) []string {

	var result []string

	var current []rune
	for _, t := range s {
		switch t {
		case '(', ')', '*', '+':
			{
				if len(current) > 0 {
					result = append(result, string(current))
					current = []rune{}
				}

				result = append(result, string(t))
			}
		case ' ':
			{
				if len(current) > 0 {
					result = append(result, string(current))
					current = []rune{}
				}
			}
		default:
			current = append(current, t)
		}

	}

	if len(current) > 0 {
		result = append(result, string(current))
	}

	return result

}

func evaluate(s string) int {
	var stack []Operation

	tokens := tokenize(s)
	// fmt.Printf("Tokens are %v\n", tokens)

	r := 0

	var prevToken string = "+"

	for _, t := range tokens {
		switch t {
		case "(":
			{
				if prevToken == "+" {
					stack = append(stack, AddOperation{r})
				}

				if prevToken == "*" {
					stack = append(stack, MultiplyOperation{r})
				}
				r = 0
				prevToken = "+"
			}

		case ")":
			{
				op := stack[len(stack)-1]

				stack[len(stack)-1] = nil
				stack = stack[:len(stack)-1]

				r = op.exec(r)
			}

		case "+", "*":
			{
				prevToken = t
			}

		default:
			v, err := strconv.Atoi(t)
			check(err)

			if prevToken == "+" {
				r += v
				// fmt.Printf("Adding %d result is %d\n", v, r)
			}

			if prevToken == "*" {
				r *= v
				// fmt.Printf("Multiply %d result is %d\n", v, r)
			}

		}
	}

	return r
}

func main() {
	dat, err := ioutil.ReadFile("./day18/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	r := 0
	for _, row := range rows {
		r += evaluate(row)

	}

	fmt.Printf("Solution is %d\n", r)
}
