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

var containerRegexp = regexp.MustCompile("(.*?) bags contain")
var contentRegexp = regexp.MustCompile("(\\d+) (.*?) bags?")

// BagDef defines how many bag of a given name fit into another bag
type BagDef struct {
	name  string
	count int
}

func main() {
	dat, err := ioutil.ReadFile("./day7/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	g := make(map[string][]BagDef)

	for _, row := range rows {
		container := containerRegexp.FindAllStringSubmatch(row, -1)
		content := contentRegexp.FindAllStringSubmatch(row, -1)

		containerName := container[0][1]

		for _, item := range content {
			count, err := strconv.Atoi(item[1])
			check(err)
			g[containerName] = append(g[containerName], BagDef{name: item[2], count: count})
		}
	}

	var traverse func(node string) int

	traverse = func(node string) int {
		r := 1

		for _, nb := range g[node] {
			r += nb.count * traverse(nb.name)
		}

		return r
	}

	r := traverse("shiny gold")

	fmt.Printf("Solution is %d\n", r-1)

}
