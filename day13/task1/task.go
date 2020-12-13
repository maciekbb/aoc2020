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

func computeWaitTime(arrivalTime, bus int) int {
	multiplier := arrivalTime / bus

	firstDeparture := bus * (multiplier + 1)

	return firstDeparture - arrivalTime
}

func main() {
	dat, err := ioutil.ReadFile("./day13/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	arrivalTime, err := strconv.Atoi(rows[0])
	check(err)

	buses := strings.Split(rows[1], ",")

	var busesInService []int

	for _, b := range buses {
		if b == "x" {
			continue
		}

		busNumber, err := strconv.Atoi(b)
		check(err)

		busesInService = append(busesInService, busNumber)
	}

	minTime := computeWaitTime(arrivalTime, busesInService[0])
	busID := busesInService[0]

	for _, b := range busesInService[1:] {
		t := computeWaitTime(arrivalTime, b)

		if t < minTime {
			minTime = t
			busID = b
		}
	}

	fmt.Printf("Solution is %d\n", minTime*busID)

}
