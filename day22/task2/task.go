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

func (d *Deck) hashCode() string {
	var arr []string

	for t := d.cards.Back(); t != nil; t = t.Prev() {
		arr = append(arr, fmt.Sprintf("%d", t.Value.(int)))
	}

	return strings.Join(arr, ",")
}

func (d *Deck) copy(amount int) *Deck {
	copiedDeck := list.New()

	i := 0
	for c := d.cards.Front(); c != nil && i < amount; c = c.Next() {
		copiedDeck.PushBack(c.Value)
		i++
	}

	return &Deck{copiedDeck}
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

func play(p1, p2 *Deck) int {
	memo1 := make(map[string]bool)
	memo2 := make(map[string]bool)

	for p1.cards.Len() > 0 && p2.cards.Len() > 0 {
		if memo1[p1.hashCode()] {
			return 1
		}

		if memo2[p2.hashCode()] {
			return 1
		}

		memo1[p1.hashCode()] = true
		memo2[p2.hashCode()] = true

		v1 := p1.take()
		v2 := p2.take()

		if p1.cards.Len() >= v1 && p2.cards.Len() >= v2 {
			winner := play(p1.copy(v1), p2.copy(v2))

			if winner == 1 {
				p1.cards.PushBack(v1)
				p1.cards.PushBack(v2)
			} else {
				p2.cards.PushBack(v2)
				p2.cards.PushBack(v1)
			}
		} else {
			if v1 > v2 {
				p1.cards.PushBack(v1)
				p1.cards.PushBack(v2)
			} else {
				p2.cards.PushBack(v2)
				p2.cards.PushBack(v1)
			}
		}

	}

	if p2.cards.Len() == 0 {
		return 1
	}

	return 2

}

func main() {
	dat, err := ioutil.ReadFile("./day22/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	sections := strings.Split(content, "\n\n")

	p1 := Deck{getCards(sections[0])}
	p2 := Deck{getCards(sections[1])}

	play(&p1, &p2)

	fmt.Printf("Game end %d %d\n", p1.score(), p2.score())

}
