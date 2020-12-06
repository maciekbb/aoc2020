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

func main() {
	dat, err := ioutil.ReadFile("./day6/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	groups := strings.Split(content, "\n\n")

	cnt := 0
	for _, group := range groups {
		people := strings.Split(group, "\n")
		m := make(map[rune]bool)
		for _, person := range people {
			for _, code := range person {
				m[code] = true
			}
		}
		cnt += len(m)

	}

	fmt.Printf("Solution is: %d", cnt)
}
