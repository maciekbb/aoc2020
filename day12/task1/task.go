package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var directions = [4]rune{'N', 'E', 'S', 'W'}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Instruction changes ship state.
type Instruction interface {
	visit(s *ship)
}

// Rotate describes ship rotation.
type Rotate struct {
	change int
}

func (r Rotate) visit(s *ship) {
	shift := r.change / 90

	for i := range directions {
		if s.direction == directions[i] {
			newDirection := directions[(i+shift)%len(directions)]
			s.direction = newDirection
			break
		}
	}
}

// Forward describes ship move in its current direction.
type Forward struct {
	change int
}

func (f Forward) visit(s *ship) {
	switch s.direction {
	case 'N':
		s.y = s.y + f.change
	case 'E':
		s.x = s.x + f.change
	case 'S':
		s.y = s.y - f.change
	case 'W':
		s.x = s.x - f.change
	}
}

// Move describes ship move in a specified direction.
type Move struct {
	direction rune
	change    int
}

func (m Move) visit(s *ship) {
	switch m.direction {
	case 'N':
		s.y = s.y + m.change
	case 'E':
		s.x = s.x + m.change
	case 'S':
		s.y = s.y - m.change
	case 'W':
		s.x = s.x - m.change
	}
}

type ship struct {
	x         int
	y         int
	direction rune
}

func (s *ship) command(ins Instruction) {
	ins.visit(s)
}

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}

func main() {
	dat, err := ioutil.ReadFile("./day12/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var instructions []Instruction

	for _, row := range rows {
		direction := []rune(row)[0]
		change, err := strconv.Atoi(row[1:])
		check(err)

		switch direction {
		case 'R':
			instructions = append(instructions, Rotate{change})
		case 'L':
			instructions = append(instructions, Rotate{360 - change})
		case 'F':
			instructions = append(instructions, Forward{change})
		case 'N', 'E', 'S', 'W':
			instructions = append(instructions, Move{direction, change})
		}

	}

	ship := ship{x: 0, y: 0, direction: 'E'}

	for _, ins := range instructions {
		ship.command(ins)

		// fmt.Printf("New postion is %d, %d, %c\n", ship.x, ship.y, ship.direction)
	}

	fmt.Printf("Solution is %d\n", abs(ship.x)+abs(ship.y))

}
