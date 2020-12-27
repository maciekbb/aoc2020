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

func computeKey(sNo, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= sNo
		val = val % 20201227
	}

	return val
}

func guessLoopSize(val, sNo int) int {
	idx := 1

	guess := sNo

	for {
		guess = guess * sNo
		guess = guess % 20201227

		if guess == val {
			return idx + 1
		}
		idx++
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day25/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	publicA, err := strconv.Atoi(rows[0])
	check(err)
	publicB, err := strconv.Atoi(rows[1])
	check(err)

	loopSize := guessLoopSize(publicA, 7)
	fmt.Printf("Loop size is %d\n", loopSize)
	fmt.Printf("Enc key: %d\n", computeKey(publicB, loopSize))

}
