package main

import (
	"container/list"
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

// Deck is a deck of cards.
type Deck struct {
	cards *list.List
}

func (d *Deck) take() int {
	first := d.cards.Front()
	d.cards.Remove(first)
	return first.Value.(int)
}

func (d *Deck) score() int {
	s := 0
	i := 1

	for t := d.cards.Back(); t != nil; t = t.Prev() {
		s += t.Value.(int) * i
		i++
	}

	return s
}

func getCards(section string) *list.List {
	cards := list.New()
	for _, card := range strings.Split(section, "\n")[1:] {
		card, err := strconv.Atoi(card)
		check(err)
		cards.PushBack(card)
	}

	return cards
}

func main() {
	dat, err := ioutil.ReadFile("./day22/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	sections := strings.Split(content, "\n\n")

	p1 := Deck{getCards(sections[0])}
	p2 := Deck{getCards(sections[1])}

	for p1.cards.Len() > 0 && p2.cards.Len() > 0 {
		v1 := p1.take()
		v2 := p2.take()

		if v1 > v2 {
			p1.cards.PushBack(v1)
			p1.cards.PushBack(v2)
		} else {
			p2.cards.PushBack(v2)
			p2.cards.PushBack(v1)
		}
	}

	fmt.Printf("Game end %d %d\n", p1.score(), p2.score())

}
