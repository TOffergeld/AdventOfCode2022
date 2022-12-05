package daythree

import (
	"AdventOfCode/misc"
	"bufio"
	"github.com/juliangruber/go-intersect"
	"log"
	"os"
)

type Group struct {
	Members [3]Rucksack
}

type Rucksack struct {
	A   []rune
	B   []rune
	All []rune
}

func (g *Group) getBadge(c chan rune) {
	is := intersect.Hash(g.Members[0].All, g.Members[1].All)
	is = intersect.Hash(is, g.Members[2].All)
	s, isRune := is[0].(rune)
	if !isRune {
		log.Fatal("No common rune")
	}
	c <- s
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
	var rucksacks []Rucksack
	f := misc.GetInputFile(path)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rucksacks = append(rucksacks, Rucksack{
			[]rune(line[:len(line)/2]), []rune(line[len(line)/2:]), []rune(line)})
	}
	return rucksacks
}

func createGroups(path string) []Group {
	var groups []Group
	var rucksackBuffer []Rucksack
	f := misc.GetInputFile(path)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	scanner := bufio.NewScanner(f)
	counter := 0
	for scanner.Scan() {
		counter++
		line := scanner.Text()
		rucksackBuffer = append(rucksackBuffer, Rucksack{
			[]rune(line[:len(line)/2]), []rune(line[len(line)/2:]), []rune(line)})
		if counter == 3 {
			counter = 0
			rucksacks := (*[3]Rucksack)(rucksackBuffer[0:3])
			groups = append(groups, Group{Members: *rucksacks})
			rucksackBuffer = nil
		}
	}
	return groups
}

func part1() {
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

func part2() {
	groups := createGroups("Day3/input.txt")
	value := 0
	priorities := getPrioritiesMap()
	c := make(chan rune)
	for idx := range groups {
		go groups[idx].getBadge(c)
		value += priorities[<-c]
	}
	println("Day 3 - Part 2 solution:", value)

}

func Day3() {
	part1()
	part2()
}
