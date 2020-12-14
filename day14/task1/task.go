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

var instructionRegexp = regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")

type instruction struct {
	addr int
	val  int
	mask string
}

func padWithZeros(val []rune, n int) []rune {
	var result []rune

	for i := 0; i < n-len(val); i++ {
		result = append(result, '0')
	}

	for i := 0; i < len(val); i++ {
		result = append(result, val[i])
	}

	return result
}

func applyMask(mask []rune, val int) string {
	valAsBin := padWithZeros([]rune(strconv.FormatInt(int64(val), 2)), len(mask))

	var result []rune

	for i := 0; i < len(mask); i++ {
		maskVal := mask[i]
		if maskVal != 'X' {
			result = append(result, maskVal)
		} else {
			result = append(result, valAsBin[i])
		}
	}

	return string(result)

}

func main() {
	dat, err := ioutil.ReadFile("./day14/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var instructions []instruction

	var currentMask string

	for _, row := range rows {
		if strings.HasPrefix(row, "mask") {
			currentMask = row[7:]
			continue
		}

		parsed := instructionRegexp.FindAllStringSubmatch(row, -1)

		addr, err := strconv.Atoi(parsed[0][1])
		check(err)

		val, err := strconv.Atoi(parsed[0][2])
		check(err)

		instructions = append(instructions, instruction{addr, val, currentMask})

	}

	mem := make(map[int]string)

	for _, ins := range instructions {
		mem[ins.addr] = applyMask([]rune(ins.mask), ins.val)
	}

	var s int64 = 0

	for _, v := range mem {
		val, err := strconv.ParseInt(v, 2, 64)
		check(err)
		s += val
	}

	fmt.Printf("Solution is %d\n", s)
}
