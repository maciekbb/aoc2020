package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func countOccupiedAround(board [][]rune, x, y int) int {
	c := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			xx := x + i
			yy := y + j

			if xx < 0 || xx > len(board)-1 {
				continue
			}

			if yy < 0 || yy > len(board[0])-1 {
				continue
			}

			if board[xx][yy] == '#' {
				c++
			}
		}
	}

	return c
}

func makeNewBoard(x, y int) [][]rune {
	arr := make([][]rune, x)

	for i := 0; i < x; i++ {
		arr[i] = make([]rune, y)
	}

	return arr
}

func printBoard(board [][]rune) {
	for _, r := range board {
		fmt.Println(string(r))
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day11/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	n := len(rows)
	m := len(rows[0])

	state := makeNewBoard(n, m)

	for i := 0; i < n; i++ {
		state[i] = []rune(rows[i])
	}

	for {
		newState := makeNewBoard(n, m)

		changesCount := 0

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if state[i][j] == '.' {
					newState[i][j] = '.'
					continue
				}

				cnt := countOccupiedAround(state, i, j)

				if cnt == 0 {
					newState[i][j] = '#'

					if state[i][j] != '#' {
						changesCount++
					}
				} else if cnt >= 4 {
					newState[i][j] = 'L'

					if state[i][j] != 'L' {
						changesCount++
					}
				} else {
					newState[i][j] = state[i][j]
				}

			}
		}

		state = newState

		if changesCount == 0 {
			break
		}
	}

	cnt := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if state[i][j] == '#' {
				cnt++
			}
		}
	}

	fmt.Printf("Solution is %d\n", cnt)

}
