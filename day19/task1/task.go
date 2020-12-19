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

// Rule represents any rule
type Rule interface {
	meets(lookup map[int]Rule, s string) (bool, []string)
}

// ComplexRule consists of subrules.
type ComplexRule struct {
	subrules []SequenceRule
}

// SequenceRule is a list of rule references
type SequenceRule struct {
	seq []int
}

func (r SequenceRule) meets(lookup map[int]Rule, s string) (bool, []string) {
	if len(r.seq) == 0 {
		return true, []string{s}
	}

	firstRule := lookup[r.seq[0]]
	nextRule := SequenceRule{r.seq[1:]}

	match, reminders := firstRule.meets(lookup, s)

	if match == false {
		return false, []string{}
	}

	var allReminders []string

	for _, reminder := range reminders {
		match, reminders := nextRule.meets(lookup, reminder)

		if match {
			allReminders = append(allReminders, reminders...)
		}
	}

	if len(allReminders) > 0 {
		return true, allReminders
	}

	return false, allReminders
}

func (r ComplexRule) meets(lookup map[int]Rule, s string) (bool, []string) {
	var allReminders []string

	for _, subrule := range r.subrules {
		if match, reminders := subrule.meets(lookup, s); match {
			allReminders = append(allReminders, reminders...)
		}
	}

	if len(allReminders) > 0 {
		return true, allReminders
	}

	return false, allReminders
}

// SimpleRule consist of a single char.
type SimpleRule struct {
	char rune
}

func (r SimpleRule) meets(_ map[int]Rule, s string) (bool, []string) {
	if strings.HasPrefix(s, string(r.char)) {
		return true, []string{s[1:]}
	}

	return false, []string{}
}

func main() {
	dat, err := ioutil.ReadFile("./day19/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	sections := strings.Split(content, "\n\n")

	rulesByID := make(map[int]Rule)

	rules := strings.Split(sections[0], "\n")

	for _, rule := range rules {
		parts := strings.Split(rule, ":")

		ruleID, err := strconv.Atoi(parts[0])
		check(err)

		parts[1] = strings.TrimSpace(parts[1])

		if strings.HasPrefix(parts[1], `"`) {
			rule := &SimpleRule{[]rune(parts[1])[1]}
			rulesByID[ruleID] = rule
		} else {
			rule := &ComplexRule{[]SequenceRule{}}
			rulesByID[ruleID] = rule

			subrules := strings.Split(parts[1], "|")

			for _, subrule := range subrules {
				subruleSeqence := strings.Split(strings.TrimSpace(subrule), " ")

				subrule := SequenceRule{}

				for _, element := range subruleSeqence {
					ref, err := strconv.Atoi(element)
					check(err)

					subrule.seq = append(subrule.seq, ref)
				}

				rule.subrules = append(rule.subrules, subrule)

			}
		}
	}

	cnt := 0
	for _, text := range strings.Split(sections[1], "\n") {
		match, reminders := rulesByID[0].meets(rulesByID, text)
		if match {
			for _, r := range reminders {
				if r == "" {
					cnt++
					break
				}
			}
		}
	}

	fmt.Printf("Solution is %d\n", cnt)
}
