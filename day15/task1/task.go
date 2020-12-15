package main

import (
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
	dat, err := ioutil.ReadFile("./day15/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")

	startSeq := strings.Split(content, ",")

	mem := make(map[int]int)

	i := 0

	var last int

	for i < len(startSeq) {
		current, err := strconv.Atoi(startSeq[i])
		check(err)

		if i > 0 {
			mem[last] = i
		}

		last = current

		i++

	}

	for i < 30000000 {
		var current int
		if prevToLast, ok := mem[last]; ok {
			current = i - prevToLast
		} else {
			current = 0
		}

		mem[last] = i

		last = current

		i++

	}

}
