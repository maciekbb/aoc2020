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
		if maskVal == '0' {
			result = append(result, valAsBin[i])
		} else {
			result = append(result, maskVal)
		}
	}

	// fmt.Println(string(valAsBin))
	// fmt.Println(string(mask))
	// fmt.Println(string(result))

	return string(result)
}

func generateAll(addr []rune) []string {
	for i, x := range addr {
		if x == 'X' {
			a := make([]rune, len(addr))
			b := make([]rune, len(addr))

			copy(a, addr)
			copy(b, addr)

			a[i] = '1'
			b[i] = '0'

			var result []string

			// fmt.Printf("Split %v %v\n", string(a), string(b))

			result = append(result, generateAll(a)...)
			result = append(result, generateAll(b)...)

			// fmt.Printf("Partial results %v\n", result)

			return result
		}
	}

	return []string{string(addr)}
}

func getAddresses(addr int, mask []rune) []string {
	return generateAll([]rune(applyMask(mask, addr)))
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

	mem := make(map[string]int)

	for _, ins := range instructions {
		addresses := getAddresses(ins.addr, []rune(ins.mask))
		for _, addr := range addresses {
			// fmt.Printf("Writing %d to %s\n", ins.val, addr)
			mem[addr] = ins.val
		}
	}

	// fmt.Printf("mem = %v\n", mem)

	s := 0

	for _, v := range mem {

		s += v
	}

	fmt.Printf("Solution is %d\n", s)
}
