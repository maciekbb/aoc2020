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

// Passport hold passport attributes.
type Passport map[string]string

func (p Passport) getNumValue(k string) (int, error) {
	v, ok := p[k]

	if !ok {
		return 0, fmt.Errorf("key not found %s", k)
	}

	numV, err := strconv.Atoi(v)

	if err != nil {
		return 0, err
	}

	return numV, nil

}

func (p Passport) checkNumValue(k string, a, b int) error {
	v, err := p.getNumValue(k)
	if err != nil {
		return err
	}

	if v < a || v > b {
		return fmt.Errorf("value out of range for %s", k)
	}

	return nil
}

func (p Passport) checkRegexp(k string, re *regexp.Regexp) error {
	val, ok := p[k]
	if !ok {
		return fmt.Errorf("key %s not found", k)
	}

	if !re.Match([]byte(val)) {
		return fmt.Errorf("%s for key %s does not match %v", val, k, re)
	}

	return nil
}

var hgtRe = regexp.MustCompile("^(\\d+)(cm|in)$")
var hclRe = regexp.MustCompile("^#[0-9a-f]{6}$")
var eclRe = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
var pidRe = regexp.MustCompile("^[0-9]{9}$")

func (p Passport) isValid() bool {
	// 	byr (Birth Year) - four digits; at least 1920 and at most 2002.
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	// cid (Country ID) - ignored, missing or not.

	var err error

	err = p.checkNumValue("byr", 1920, 2002)
	if err != nil {
		return false
	}

	err = p.checkNumValue("iyr", 2010, 2020)
	if err != nil {
		return false
	}

	err = p.checkNumValue("eyr", 2020, 2030)
	if err != nil {
		return false
	}

	hgt, ok := p["hgt"]
	if !ok {
		return false
	}

	hgtMatches := hgtRe.FindAllStringSubmatch(hgt, -1)

	if hgtMatches == nil {
		return false
	}

	hgtVal, err := strconv.Atoi(hgtMatches[0][1])

	if err != nil {
		return false
	}

	if hgtMatches[0][2] == "cm" && (hgtVal < 150 || hgtVal > 193) {
		return false
	}

	if hgtMatches[0][2] == "in" && (hgtVal < 53 || hgtVal > 76) {
		return false
	}

	err = p.checkRegexp("hcl", hclRe)
	if err != nil {
		return false
	}

	err = p.checkRegexp("ecl", eclRe)
	if err != nil {
		return false
	}

	err = p.checkRegexp("pid", pidRe)
	if err != nil {
		return false
	}

	return true
}

func prepareLookup(row string) Passport {
	items := strings.Split(strings.Join(strings.Split(row, "\n"), " "), " ")

	m := make(map[string]string)
	for _, item := range items {
		split := strings.Split(item, ":")
		k := split[0]
		v := split[1]
		m[k] = v
	}

	return m
}

func main() {
	dat, err := ioutil.ReadFile("./day4/task1/input.txt")
	check(err)

	content := strings.TrimSuffix(string(dat), "\n")
	rows := strings.Split(content, "\n\n")

	cnt := 0
	for _, row := range rows {
		m := prepareLookup(row)
		if m.isValid() {
			cnt++
		}
	}

	fmt.Printf("Solution is: %d\n", cnt)
}
