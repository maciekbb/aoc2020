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

// Node represents a node in a linked list.
type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func printList(head, current *Node) {
	t := head
	for {
		if t == current {
			fmt.Printf("(%d) ", t.Value)
		} else {
			fmt.Printf("%d ", t.Value)
		}

		t = t.Next

		if t == head {
			break
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("./day23/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")

	var head *Node
	var current *Node
	var prev *Node

	highest := 0

	// values in nodes are uniq
	lookup := make(map[int]*Node)

	for _, ch := range content {
		val, err := strconv.Atoi(string(ch))
		check(err)

		node := &Node{
			Value: val,
			Prev:  prev,
			Next:  nil,
		}

		lookup[val] = node

		if prev != nil {
			prev.Next = node
		} else {
			head = node
		}

		prev = node

		if val > highest {
			highest = val
		}
	}

	// form a circle
	prev.Next = head

	current = head

	for i := 1; i <= 100; i++ {
		currentValue := current.Value

		fmt.Printf("Move -- %d --\n", i)

		printList(head, current)

		fmt.Println()

		i1 := current.Next
		i2 := current.Next.Next
		i3 := current.Next.Next.Next

		fmt.Printf("pick up %d %d %d\n", i1.Value, i2.Value, i3.Value)

		var destination *Node

		k := 1
		for destination == nil || destination == i1 || destination == i2 || destination == i3 {
			if currentValue-k > 0 {
				destination = lookup[currentValue-k]
			} else {
				destination = lookup[highest+currentValue-k]

			}

			k++
		}

		fmt.Printf("destination %d\n", destination.Value)

		// disconnect i1 - i3
		i1.Prev.Next = i3.Next
		i3.Next.Prev = i1.Prev

		// connect i1-i3
		i3.Next = destination.Next
		destination.Next.Prev = i3
		destination.Next = i1
		i1.Prev = destination

		current = current.Next

	}

	printList(head, current)

}
