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

// Passport hold passport attributes.
type Passport map[string]string

func (p Passport) isValid() bool {
	if _, ok := p["cid"]; !ok {
		return len(p) == 7
	}

	return len(p) == 8
}

func prepareLookup(row string) Passport {
	items := strings.Split(strings.Join(strings.Split(row, "\n"), " "), " ")

	m := make(map[string]string)
	for _, item := range items {
		split := strings.Split(item, ":")
		k := split[0]
		v := split[1]
		m[k] = v
	}

	return m
}

func main() {
	dat, err := ioutil.ReadFile("./day4/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n\n")

	cnt := 0
	for _, row := range rows {
		m := prepareLookup(row)
		if m.isValid() {
			cnt++
		}
	}

	fmt.Printf("Solution is: %d\n", cnt)
}
