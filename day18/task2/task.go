package main

import (
	"container/list"
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

func tokenize(s string) *list.List {

	result := list.New()

	var current []rune
	for _, t := range s {
		switch t {
		case '(', ')', '*', '+':
			{
				if len(current) > 0 {
					result.PushBack(string(current))
					current = []rune{}
				}

				result.PushBack(string(t))
			}
		case ' ':
			{
				if len(current) > 0 {
					result.PushBack(string(current))
					current = []rune{}
				}
			}
		default:
			current = append(current, t)
		}

	}

	if len(current) > 0 {
		result.PushBack(string(current))
	}

	return result

}

func insertParenthesesBefore(tokens *list.List, t *list.Element) {
	n := t.Prev()

	if n.Value != ")" {
		tokens.InsertBefore("(", n)
	} else {
		cnt := 1

		for cnt != 0 {
			n = n.Prev()

			if n.Value == ")" {
				cnt++
			}

			if n.Value == "(" {
				cnt--
			}
		}

		tokens.InsertBefore("(", n)
	}
}

func insertParenthesesAfter(tokens *list.List, t *list.Element) {
	n := t.Next()

	if n.Value != "(" {
		tokens.InsertAfter(")", n)
	} else {
		cnt := 1

		for cnt != 0 {
			n = n.Next()

			if n.Value == "(" {
				cnt++
			}

			if n.Value == ")" {
				cnt--
			}
		}

		tokens.InsertAfter(")", n)
	}
}

func prioritize(tokens *list.List) {
	for t := tokens.Front(); t != nil; t = t.Next() {
		if t.Value == "+" {
			insertParenthesesBefore(tokens, t)
			insertParenthesesAfter(tokens, t)
		}
	}

}

func evaluate(s string) int {
	var stack []Operation

	tokens := tokenize(s)

	prioritize(tokens)

	r := 0

	var prevToken string = "+"

	for t := tokens.Front(); t != nil; t = t.Next() {
		switch t.Value {
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
				prevToken = t.Value.(string)
			}

		default:
			v, err := strconv.Atoi(t.Value.(string))
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
