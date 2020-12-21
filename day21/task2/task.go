package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
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

type Pair struct {
	alg string
	ing string
}

type byAlg []Pair

func (s byAlg) Len() int {
	return len(s)
}
func (s byAlg) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byAlg) Less(i, j int) bool {
	return strings.Compare(s[i].alg, s[j].alg) == -1
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

		// fmt.Printf("%v: %v\n", ingredients, allergens)

		foods = append(foods, Food{ingredients, allergens})
	}

	fmt.Printf("Possible mappings: %v\n", possibleMappings)

	var validMappings []map[string]string

	var distribute func(ingredients map[string]bool, allergens []string)

	isValidMapping := func(mapping map[string]string) bool {
		for _, food := range foods {
			if !food.match(mapping) {
				return false
			}
		}

		return true
	}

	mapping := make(map[string]string)

	distribute = func(ingredients map[string]bool, allergens []string) {
		if len(allergens) == 0 {
			if isValidMapping(mapping) {
				persistedMapping := make(map[string]string)
				for k, v := range mapping {
					persistedMapping[k] = v
				}

				validMappings = append(validMappings, persistedMapping)

			}
			return
		}

		alg := allergens[0]

		for ing := range ingredients {
			if _, ok := possibleMappings[alg][ing]; !ok {
				continue
			}

			if _, ok := mapping[ing]; ok {
				continue
			}

			mapping[ing] = alg
			distribute(ingredients, allergens[1:])
			delete(mapping, ing)
		}
	}

	distribute(allIngredients, keys(allAlergens))

	var unsafePairs []Pair

	for ing, alg := range validMappings[0] {
		unsafePairs = append(unsafePairs, Pair{alg, ing})
	}

	sort.Sort(byAlg(unsafePairs))

	var result []string

	for _, pair := range unsafePairs {
		result = append(result, pair.ing)
	}

	fmt.Printf("Solution is %s\n", strings.Join(result, ","))

}
