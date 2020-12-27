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

// Tile is a hexagon.
type Tile struct {
	e  int
	ne int
	se int
}

func main() {
	dat, err := ioutil.ReadFile("./day24/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	flipped := make(map[string]bool)

	for _, row := range rows {
		cnt := make(map[string]int)

		i := 0

		for i < len(row) {

			if []rune(row)[i] == 'w' || []rune(row)[i] == 'e' {
				cnt[string([]rune(row)[i])]++
				i++
			} else {
				cnt[string([]rune(row)[i:i+2])]++
				i += 2
			}
		}

		e := cnt["e"] - cnt["w"]
		ne := cnt["ne"] - cnt["sw"]
		se := cnt["se"] - cnt["nw"]

		for se > 0 {
			se--
			ne--
			e++
		}

		for se < 0 {
			se++
			ne++
			e--
		}

		code := fmt.Sprintf("%d,%d", e, ne)

		if _, ok := flipped[code]; ok {
			// fmt.Printf("Delete %s\n", code)
			delete(flipped, code)
		} else {
			// fmt.Printf("Add %s\n", code)
			flipped[code] = true
		}

	}

	// fmt.Printf("Solution is %d\n", len(flipped))

	for i := 0; i < 100; i++ {
		contrib := make(map[string]int)
		for key := range flipped {
			parts := strings.Split(key, ",")
			e, _ := strconv.Atoi(parts[0])
			ne, _ := strconv.Atoi(parts[1])

			if _, ok := contrib[fmt.Sprintf("%d,%d", e, ne)]; !ok {
				contrib[fmt.Sprintf("%d,%d", e, ne)] = 0
			}

			contrib[fmt.Sprintf("%d,%d", e, ne+1)]++
			contrib[fmt.Sprintf("%d,%d", e, ne-1)]++

			contrib[fmt.Sprintf("%d,%d", e+1, ne)]++
			contrib[fmt.Sprintf("%d,%d", e-1, ne)]++

			contrib[fmt.Sprintf("%d,%d", e-1, ne+1)]++
			contrib[fmt.Sprintf("%d,%d", e+1, ne-1)]++
		}

		for key := range contrib {
			if flipped[key] {
				if contrib[key] == 0 || contrib[key] > 2 {
					delete(flipped, key)
				}
			} else {
				if contrib[key] == 2 {
					flipped[key] = true
				}
			}
		}

	}

	fmt.Printf("Solution is %d\n", len(flipped))

}
