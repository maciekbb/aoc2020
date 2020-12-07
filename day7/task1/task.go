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

var containerRegexp = regexp.MustCompile("(.*?) bags contain")
var contentRegexp = regexp.MustCompile("\\d+ (.*?) bags?")

func main() {
	dat, err := ioutil.ReadFile("./day7/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n")

	g := make(map[string][]string)

	for _, row := range rows {
		container := containerRegexp.FindAllStringSubmatch(row, -1)
		content := contentRegexp.FindAllStringSubmatch(row, -1)

		containerName := container[0][1]

		for _, item := range content {
			g[item[1]] = append(g[item[1]], containerName)
		}
	}

	var traverse func(node string)

	visited := make(map[string]bool)

	traverse = func(node string) {
		visited[node] = true

		for _, nb := range g[node] {
			if _, ok := visited[nb]; !ok {
				traverse(nb)
			}
		}
	}

	traverse("shiny gold")

	fmt.Printf("Solution is %d\n", len(visited)-1)

}
