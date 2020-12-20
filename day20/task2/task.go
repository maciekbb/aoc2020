package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Tile represents a single tile.
type Tile struct {
	id      int
	symbols [][]rune
}

func (t Tile) String() string {
	result := fmt.Sprint(t.id) + "\n"

	for _, row := range t.symbols {
		result += string(row) + "\n"
	}

	return result
}

func (t Tile) flipVertical() Tile {
	n := len(t.symbols)
	symbols := make([][]rune, n)

	for i := 0; i < n; i++ {
		symbols[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			symbols[i][j] = t.symbols[i][n-j-1]
		}
	}

	return Tile{t.id, symbols}
}

func (t Tile) flipHorizontal() Tile {
	n := len(t.symbols)
	symbols := make([][]rune, n)

	for i := 0; i < n; i++ {
		symbols[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			symbols[i][j] = t.symbols[n-i-1][j]
		}
	}

	return Tile{t.id, symbols}
}

func (t Tile) rotate() Tile {
	n := len(t.symbols)
	symbols := make([][]rune, n)

	for i := 0; i < n; i++ {
		symbols[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			symbols[i][j] = t.symbols[j][i]
		}
	}

	return Tile{t.id, symbols}.flipVertical()
}

func fitVertically(t1, t2 Tile) bool {
	n := len(t1.symbols)

	for i := 0; i < n; i++ {
		if t1.symbols[n-1][i] != t2.symbols[0][i] {
			return false
		}
	}

	return true
}

func fitHorizontally(t1, t2 Tile) bool {
	n := len(t1.symbols)

	for i := 0; i < n; i++ {
		if t1.symbols[i][n-1] != t2.symbols[i][0] {
			return false
		}
	}

	return true
}

func areTheSame(t1, t2 Tile) bool {
	n := len(t1.symbols)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if t1.symbols[i][j] != t2.symbols[i][j] {
				return false
			}
		}
	}

	return true
}

func generateExtra(t Tile) []Tile {
	var result []Tile

	maybeAppend := func(newTile Tile) {
		for _, tile := range result {
			if areTheSame(newTile, tile) {
				return
			}
		}

		result = append(result, newTile)
	}

	maybeAppend(t.flipVertical())
	maybeAppend(t.flipHorizontal())

	for i := 1; i < 4; i++ {
		rotated := t
		for j := 0; j < i; j++ {
			rotated = rotated.rotate()
		}

		maybeAppend(rotated)
		maybeAppend(rotated.flipVertical())
		maybeAppend(rotated.flipHorizontal())
	}

	return result
}

func main() {
	dat, err := ioutil.ReadFile("./day20/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	sections := strings.Split(content, "\n\n")

	n := int(math.Sqrt(float64(len(sections))))

	var tiles []Tile

	for _, s := range sections {
		rows := strings.Split(s, "\n")

		id, err := strconv.Atoi(rows[0][5 : len(rows[0])-1])
		check(err)
		var symbols [][]rune

		for _, r := range rows[1:] {
			symbols = append(symbols, []rune(r))
		}

		tile := Tile{id, symbols}

		tiles = append(tiles, tile)
	}

	var extraTiles []Tile
	for _, tile := range tiles {
		extras := generateExtra(tile)

		extraTiles = append(extraTiles, extras...)
	}

	tiles = append(tiles, extraTiles...)
	fmt.Printf("In total there are %d tiles\n", len(tiles))

	image := make([][]*Tile, n)

	for i := 0; i < n; i++ {
		image[i] = make([]*Tile, n)
	}

	takenTiles := make(map[int]bool)

	possibleTiles := func(i, j int) []*Tile {
		var result []*Tile

		for k := 0; k < len(tiles); k++ {
			if takenTiles[tiles[k].id] {
				continue
			}

			if i > 0 {
				toTop := image[i-1][j]
				if !fitVertically(*toTop, tiles[k]) {
					continue
				}
			}

			if j > 0 {
				toLeft := image[i][j-1]
				if !fitHorizontally(*toLeft, tiles[k]) {
					continue
				}
			}

			result = append(result, &tiles[k])
		}

		return result
	}

	var distribute func(i, j int) bool

	distribute = func(i, j int) bool {
		possibilities := possibleTiles(i, j)

		for _, p := range possibilities {

			image[i][j] = p

			takenTiles[p.id] = true

			if j < n-1 {
				if distribute(i, j+1) {
					return true
				}
			} else if i < n-1 {
				if distribute(i+1, 0) {
					return true
				}
			} else {
				return true
			}

			image[i][j] = nil
			delete(takenTiles, p.id)
		}

		return false
	}

	distribute(0, 0)

	m := len(tiles[0].symbols) - 2

	picture := make([][]rune, n*m)

	for i := 0; i < m*n; i++ {
		picture[i] = make([]rune, n*m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tile := image[i][j]
			for k := 1; k < len(tile.symbols)-1; k++ {
				for l := 1; l < len(tile.symbols)-1; l++ {
					s := tile.symbols[k][l]
					picture[i*m+k-1][j*m+l-1] = s
				}
			}
		}
	}

	pictureTile := Tile{0, picture}

	pattern := [][]rune{
		[]rune("                  # "),
		[]rune("#    ##    ##    ###"),
		[]rune(" #  #  #  #  #  #   "),
	}

	roughtCount := 0

	for i := 0; i < m*n; i++ {
		for j := 0; j < m*n; j++ {
			if pictureTile.symbols[i][j] == '#' {
				roughtCount++
			}
		}
	}

	toSearchMonsterIn := []Tile{pictureTile}
	extras := generateExtra(pictureTile)

	toSearchMonsterIn = append(toSearchMonsterIn, extras...)

	for _, tile := range toSearchMonsterIn {
		// fmt.Printf("%v\n", tile)
		monsterCount := 0
		foundAny := false

		for i := 0; i < len(tile.symbols)-len(pattern); i++ {
			for j := 0; j < len(tile.symbols)-len(pattern[0]); j++ {
				// fmt.Printf("Checking offset (%d, %d)\n", i, j)
				found := true
				cnt := 0

				for k := 0; k < len(pattern); k++ {
					for l := 0; l < len(pattern[0]); l++ {
						if tile.symbols[i+k][j+l] == '#' && pattern[k][l] == '#' {
							cnt++
						}

						if tile.symbols[i+k][j+l] != '#' && pattern[k][l] == '#' {
							found = false
						}
					}
				}

				if found {
					foundAny = true
					monsterCount += cnt
				}

			}
		}

		if foundAny {
			fmt.Printf("Solution is %d\n", roughtCount-monsterCount)
		}
	}
}
