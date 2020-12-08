package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type state struct {
	acc     int
	pointer int
}

type instruction interface {
	run(state *state)
}

type nop struct{}

func (ins nop) run(s *state) {
	s.pointer = s.pointer + 1
}

type acc struct {
	change int
}

func (ins acc) run(s *state) {
	s.acc = s.acc + ins.change

	s.pointer = s.pointer + 1
}

type jmp struct {
	change int
}

func (ins jmp) run(s *state) {
	s.pointer = s.pointer + ins.change
}

type program struct {
	instructions []instruction
	state        state
	executed     []bool
}

func (p *program) run() {
	for {
		if p.executed[p.state.pointer] == true {
			break
		}

		p.executed[p.state.pointer] = true

		p.instructions[p.state.pointer].run(&p.state)

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func buildInstruction(s string) instruction {
	instruction := strings.Split(s, " ")

	switch instruction[0] {
	case "nop":
		return &nop{}
	case "acc":
		change, err := strconv.Atoi(instruction[1])
		check(err)
		return &acc{change: change}
	case "jmp":
		change, err := strconv.Atoi(instruction[1])
		check(err)
		return &jmp{change: change}
	}

	return nil
}

func main() {
	dat, err := ioutil.ReadFile("./day8/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var instructions []instruction

	for _, r := range rows {
		instructions = append(instructions, buildInstruction(r))
	}

	p := &program{
		instructions: instructions,
		state:        state{acc: 0, pointer: 0},
		executed:     make([]bool, len(instructions)),
	}

	p.run()

	fmt.Printf("Solution is %d\n", p.state.acc)
}
