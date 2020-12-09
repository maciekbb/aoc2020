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

var windowSize = 25

func main() {
	dat, err := ioutil.ReadFile("./day9/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var numbers []int

	for _, row := range rows {
		n, err := strconv.Atoi(row)
		check(err)
		numbers = append(numbers, n)
	}

	lookup := make(map[int]int)

	for i := 0; i < windowSize; i++ {
		for j := i + 1; j < windowSize; j++ {
			s := numbers[i] + numbers[j]
			if c, ok := lookup[s]; ok {
				lookup[s] = c + 1
			} else {
				lookup[s] = 1
			}
		}
	}

	i := windowSize

	for i < len(numbers) {
		if c, ok := lookup[numbers[i]]; !ok || c == 0 {

			break
		}

		for j := 1; j < windowSize; j++ {
			s := numbers[i-windowSize] + numbers[i-j]
			lookup[s] = lookup[s] - 1
		}

		for j := 1; j < windowSize; j++ {
			s := numbers[i] + numbers[i-j]
			if c, ok := lookup[s]; ok {
				lookup[s] = c + 1
			} else {
				lookup[s] = 1
			}
		}

		i++
	}

	toFind := numbers[i]
	fmt.Printf("Looking for %d\n", toFind)

	runningSum := numbers[0]

	a := 0
	b := 1

	for {
		if runningSum == toFind {
			smallest := numbers[a]
			biggest := numbers[a]

			for i := a; i < b; i++ {
				if numbers[i] < smallest {
					smallest = numbers[i]
				}
				if numbers[i] > biggest {
					biggest = numbers[i]
				}
			}

			fmt.Printf("Solution is %d\n", smallest+biggest)
			break
		}

		if runningSum < toFind {
			runningSum += numbers[b]
			b++
		} else {
			runningSum -= numbers[a]
			a++
		}
	}

}
