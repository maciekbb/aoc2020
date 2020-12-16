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

// BetweenCondition describes condition to be between a and b.
type BetweenCondition struct {
	a, b int
}

// Field is a ticket field.
type Field struct {
	label      string
	conditions []BetweenCondition
}

func (r Field) meets(n int) bool {
	for _, c := range r.conditions {
		if c.a <= n && n <= c.b {
			return true
		}
	}

	return false
}

var rangeRegexp = regexp.MustCompile("(\\d+)-(\\d+)")

func meetsAny(fields []Field, n int) bool {
	for _, r := range fields {
		if r.meets(n) {
			return true
		}
	}

	return false
}

func isValidTicket(fields []Field, t string) bool {
	ticketNums := strings.Split(t, ",")

	for _, tNum := range ticketNums {
		tNumAsInt, err := strconv.Atoi(tNum)
		check(err)

		if !meetsAny(fields, tNumAsInt) {
			return false
		}
	}

	return true
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

func getValidLabels(fields []Field, ticket string, attrNo int) map[string]bool {
	tNums := strings.Split(ticket, ",")

	result := make(map[string]bool)

	tNumAsInt, err := strconv.Atoi(tNums[attrNo])
	check(err)

	for _, f := range fields {
		if f.meets(tNumAsInt) {
			result[f.label] = true
		}
	}

	return result
}

func main() {
	dat, err := ioutil.ReadFile("./day16/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	sections := strings.Split(content, "\n\n")

	var fields []Field

	for _, row := range strings.Split(sections[0], "\n") {
		parts := strings.Split(row, ":")

		extracted := rangeRegexp.FindAllStringSubmatch(parts[1], -1)

		field := Field{parts[0], []BetweenCondition{}}

		for _, r := range extracted {
			a, err := strconv.Atoi(r[1])
			check(err)

			b, err := strconv.Atoi(r[2])
			check(err)

			field.conditions = append(field.conditions, BetweenCondition{a, b})
		}

		fields = append(fields, field)
	}

	myTicket := strings.Split(sections[1], "\n")[1]
	nearbyTickets := strings.Split(sections[2], "\n")[1:]

	var validTickets []string
	validTickets = append(validTickets, myTicket)

	for _, t := range nearbyTickets {

		if isValidTicket(fields, t) {
			validTickets = append(validTickets, t)
		}
	}

	nAttributes := len(strings.Split(validTickets[0], ","))

	mapping := make(map[int]map[string]bool)

	for i := 0; i < nAttributes; i++ {

		potentialFieldLabels := getValidLabels(fields, validTickets[0], i)

		for _, t := range validTickets[1:] {
			labels := getValidLabels(fields, t, i)

			// fmt.Printf("Labels for ticket %s field %d: %v\n", t, i, labels)

			potentialFieldLabels = intersection(potentialFieldLabels, labels)
		}

		mapping[i] = potentialFieldLabels

	}

	finalMapping := make(map[string]int)

	unmappedFields := len(mapping)

	for unmappedFields > 0 {
		for field, possible := range mapping {
			if len(possible) == 1 {
				for onlyPossibleForThis := range possible {
					finalMapping[onlyPossibleForThis] = field
					unmappedFields--

					for _, possibleForOthers := range mapping {
						delete(possibleForOthers, onlyPossibleForThis)
					}
				}

				break
			}
		}
	}

	r := 1

	for _, f := range fields {
		if strings.HasPrefix(f.label, "departure") {
			myTicketNums := strings.Split(myTicket, ",")
			v, err := strconv.Atoi(myTicketNums[finalMapping[f.label]])
			check(err)
			r *= v
		}
	}

	fmt.Printf("Solution is %d\n", r)

}
