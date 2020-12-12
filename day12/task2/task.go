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

// Instruction changes ship state.
type Instruction interface {
	visit(s *ship)
}

// Rotate describes ship rotation.
type Rotate struct {
	change int
}

func (r Rotate) visit(s *ship) {
	sin := func(a int) int {
		switch a {
		case 90:
			return 1
		case 270:
			return -1
		}

		return 0
	}

	cos := func(a int) int {
		if a == 180 {
			return -1
		}

		return 0
	}

	wXNew := s.wX*cos(r.change) + s.wY*sin(r.change)
	wYNew := -s.wX*sin(r.change) + s.wY*cos(r.change)

	s.wX = wXNew
	s.wY = wYNew
}

// Forward describes ship move in its current direction.
type Forward struct {
	change int
}

func (f Forward) visit(s *ship) {
	s.x += f.change * s.wX
	s.y += f.change * s.wY
}

// Move describes ship move in a specified direction.
type Move struct {
	direction rune
	change    int
}

func (m Move) visit(s *ship) {
	switch m.direction {
	case 'N':
		s.wY = s.wY + m.change
	case 'E':
		s.wX = s.wX + m.change
	case 'S':
		s.wY = s.wY - m.change
	case 'W':
		s.wX = s.wX - m.change
	}
}

type ship struct {
	x  int
	y  int
	wX int
	wY int
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

	ship := ship{x: 0, y: 0, wX: 10, wY: 1}

	for _, ins := range instructions {
		ship.command(ins)

		// fmt.Printf("New postion is (%d, %d), waypoint (%d, %d)\n", ship.x, ship.y, ship.wX, ship.wY)
	}

	fmt.Printf("Solution is %d\n", abs(ship.x)+abs(ship.y))

}
