package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Range struct {
	a int
	b int
}

func (r Range) meets(n int) bool {
	return r.a <= n && n <= r.b
}

var rangeRegexp = regexp.MustCompile("(\\d+)-(\\d+)")

func meetsAny(ranges []Range, n int) bool {

	for _, r := range ranges {
		if r.meets(n) {
			return true
		}
	}

	return false
}

func main() {
	dat, err := ioutil.ReadFile("./day16/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	sections := strings.Split(content, "\n\n")

	var ranges []Range

	for _, row := range strings.Split(sections[0], "\n") {
		parts := strings.Split(row, ":")

		extracted := rangeRegexp.FindAllStringSubmatch(parts[1], -1)

		for _, r := range extracted {
			a, err := strconv.Atoi(r[1])
			check(err)

			b, err := strconv.Atoi(r[2])
			check(err)

			ranges = append(ranges, Range{a, b})
		}

	}

	nearbyTickets := strings.Split(sections[2], "\n")[1:]
	invalidNumsSum := 0

	for _, t := range nearbyTickets {
		ticketNums := strings.Split(t, ",")

		for _, tNum := range ticketNums {
			tNumAsInt, err := strconv.Atoi(tNum)
			check(err)

			if !meetsAny(ranges, tNumAsInt) {
				invalidNumsSum += tNumAsInt
			}
		}

	}

	fmt.Printf("Solution is %d\n", invalidNumsSum)
}
