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

// Position represents a position in 3D grid.
type Position struct {
	x, y, w, z int
}

func (p Position) hashCode() string {
	return fmt.Sprintf("%d,%d,%d,%d", p.x, p.y, p.z, p.w)
}

func (p Position) neighbors() []Position {
	var result []Position
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					result = append(result, Position{p.x + i, p.y + j, p.z + k, p.w + l})
				}
			}
		}
	}

	return result
}

func positionFromHash(hash string) Position {
	nums := strings.Split(hash, ",")

	x, err := strconv.Atoi(nums[0])
	check(err)

	y, err := strconv.Atoi(nums[1])
	check(err)

	z, err := strconv.Atoi(nums[2])
	check(err)

	w, err := strconv.Atoi(nums[3])
	check(err)

	return Position{x, y, z, w}
}

func main() {
	dat, err := ioutil.ReadFile("./day17/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	active := make(map[string]bool)

	for i, row := range rows {
		for j, cell := range row {
			if cell == '#' {
				p := Position{0, 0, i, j}
				active[p.hashCode()] = true
			}

		}
	}

	fmt.Printf("Active after epoch %d = %d\n", 0, len(active))

	for epoch := 1; epoch < 7; epoch++ {
		g := make(map[string]int)

		for hash := range active {
			position := positionFromHash(hash)
			neighbors := position.neighbors()
			for _, n := range neighbors {
				if current, ok := g[n.hashCode()]; ok {
					g[n.hashCode()] = current + 1
				} else {
					g[n.hashCode()] = 1
				}
			}
		}

		var pointsOfIntrest []string

		for hash := range g {
			pointsOfIntrest = append(pointsOfIntrest, hash)
		}

		for hash := range active {
			pointsOfIntrest = append(pointsOfIntrest, hash)
		}

		for _, hash := range pointsOfIntrest {
			value := g[hash]

			if isActive := active[hash]; isActive {
				if value != 2 && value != 3 {
					delete(active, hash)
				}
			} else {
				if value == 3 {
					active[hash] = true
				}
			}
		}

		fmt.Printf("Active after epoch %d = %d\n", epoch, len(active))

	}

}
