package daythree

import (
	"bufio"
	"github.com/juliangruber/go-intersect"
	"os"
)

type Rucksack struct {
	A   []rune
	B   []rune
	All []rune
}

func getPrioritiesMap() map[rune]int {
	m := map[rune]int{}
	for i := 1; i <= 26; i++ {
		m[rune(i+96)] = i
		m[rune(i+64)] = i + 26
	}
	return m
}

func createRucksacks(path string) []Rucksack {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var rucksacks []Rucksack

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rucksacks = append(rucksacks, Rucksack{
			[]rune(line[:len(line)/2]), []rune(line[len(line)/2:]), []rune(line)})
	}
	return rucksacks
}

func Daythree() {
	rucksacks := createRucksacks("Day3/input.txt")
	priorities := getPrioritiesMap()
	value := 0
	for rs := range rucksacks {
		is := intersect.Hash(rucksacks[rs].A, rucksacks[rs].B)
		s, isInt := is[0].(int32)
		if !isInt {
			panic("")
		}
		value += priorities[s]
	}
	println("Day 3 - Part 1 solution:", value)
}
