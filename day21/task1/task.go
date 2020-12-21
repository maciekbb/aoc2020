package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var allergenRegexp = regexp.MustCompile("(.*) \\(contains (.*)\\)")

// Food consists of ingredients.
type Food struct {
	ingredients []string
	allergens   []string
}

func (f Food) match(mapping map[string]string) bool {
	allergenSet := make(map[string]bool)

	for _, alg := range f.allergens {
		allergenSet[alg] = true
	}

	for _, ing := range f.ingredients {
		delete(allergenSet, mapping[ing])

	}

	return len(allergenSet) == 0
}

func intersection(a, b map[string]bool) map[string]bool {
	result := make(map[string]bool)

	for key := range a {
		if _, ok := b[key]; ok {
			result[key] = true
		}
	}

	return result
}

func keys(m map[string]bool) []string {
	var result []string

	for k := range m {
		result = append(result, k)
	}

	return result
}

func main() {
	dat, err := ioutil.ReadFile("./day21/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	var foods []Food

	allIngredients := make(map[string]bool)
	allAlergens := make(map[string]bool)
	possibleMappings := make(map[string]map[string]bool)

	for _, row := range rows {
		parsed := allergenRegexp.FindAllStringSubmatch(row, 1)

		ingredients := strings.Split(parsed[0][1], " ")
		allergens := strings.Split(parsed[0][2], ", ")

		ingredientsSet := make(map[string]bool)

		for _, ing := range ingredients {
			allIngredients[ing] = true
			ingredientsSet[ing] = true
		}

		for _, alg := range allergens {
			allAlergens[alg] = true

			if current, ok := possibleMappings[alg]; ok {
				possibleMappings[alg] = intersection(ingredientsSet, current)
			} else {
				possibleMappings[alg] = ingredientsSet
			}
		}

		foods = append(foods, Food{ingredients, allergens})
	}

	fmt.Printf("Possible mappings: %v\n", possibleMappings)

	unsafe := make(map[string]bool)

	for _, m := range possibleMappings {
		for ing := range m {
			unsafe[ing] = true
		}
	}

	fmt.Printf("Unsafe: %v\n", unsafe)

	cnt := 0

	for _, f := range foods {
		for _, ing := range f.ingredients {
			if !unsafe[ing] {
				cnt++
			}
		}
	}

	fmt.Printf("Solution is %d\n", cnt)
}
