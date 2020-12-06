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

func makeMap(s string) map[rune]bool {
	m := make(map[rune]bool)

	for _, code := range s {
		m[code] = true
	}

	return m
}

func intersect(a, b map[rune]bool) map[rune]bool {
	r := make(map[rune]bool)

	for k := range a {
		if _, ok := b[k]; ok {
			r[k] = true
		}
	}

	return r
}

func main() {
	dat, err := ioutil.ReadFile("./day6/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	groups := strings.Split(content, "\n\n")

	cnt := 0
	for _, group := range groups {
		people := strings.Split(group, "\n")
		m := makeMap(people[0])

		for _, person := range people[1:] {
			m = intersect(m, makeMap(person))
		}

		cnt += len(m)
	}

	fmt.Printf("Solution is: %d", cnt)
}
