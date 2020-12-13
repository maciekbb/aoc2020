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

type bus struct {
	value int
	delay int
}

func (b bus) t0(x int) int {
	return x*b.value + b.delay
}

func (b bus) leavesAt(x int) bool {
	return (x-b.delay)%b.value == 0
}

func (b bus) arg(x int) int {
	return (x - b.delay) / b.value
}

func main() {
	dat, err := ioutil.ReadFile("./day13/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var buses []bus

	for i, b := range strings.Split(rows[1], ",") {
		if b == "x" {
			continue
		}

		busNumber, err := strconv.Atoi(b)
		check(err)

		buses = append(buses, bus{busNumber, -i})
	}

	formula := buses[0]
	for _, b := range buses[1:] {
		// fmt.Printf("Formula is %v\n", formula)
		// fmt.Printf("Looking at formula %v\n", b)
		i := 0
		matches := make([]int, 0)
		for {
			x := formula.t0(i)
			if b.leavesAt(x) {
				// fmt.Printf("Found match at i, x = %d, %d\n", i, x)
				matches = append(matches, b.arg(x))
				if len(matches) == 2 {
					formula = bus{
						value: b.value * (matches[0] - matches[1]),
						delay: b.delay + b.value*matches[0],
					}

					break
				}
			}
			i--
		}
	}

	// fmt.Printf("End formula is %v\n", formula)
	fmt.Printf("Solution is %d\n", formula.t0(1))

}
