package day5

import (
	"AdventOfCode/misc"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Crate struct {
	Letter string
}

type Stack struct {
	numOfCrates int
	crates      []Crate
}

func (s *Stack) pop(count int) []Crate {
	c := s.crates[s.numOfCrates-count:]
	s.crates = s.crates[:s.numOfCrates-count]
	s.numOfCrates -= count
	return c
}

func (s *Stack) put(c []Crate) {
	// this inserts the Crates in their original order
	// Required for Part 2. In Part 1, put needs to be called for each crate individually
	for _, newCrate := range c {
		s.crates = append(s.crates, newCrate)
		s.numOfCrates++
	}
}

// popLine takes a slice of lines and pops the last one off
func popLine(sl []string) (string, []string) {
	elem := sl[len(sl)-1]
	rem := sl[:len(sl)-1]
	return elem, rem
}

func solve(path, mode string) {
	f := misc.GetInputFile(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)
	stacks := buildInitialStacks(scanner)

	for scanner.Scan() {
		line := scanner.Text()
		expr := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
		res := expr.FindStringSubmatch(line)
		count, _ := strconv.Atoi(res[1])
		from, _ := strconv.Atoi(res[2])
		to, _ := strconv.Atoi(res[3])
		if mode == "9000" {
			for i := 0; i < count; i++ {
				// Part 1: count = 1
				stacks[to-1].put(stacks[from-1].pop(1))
			}
		} else if mode == "9001" {
			stacks[to-1].put(stacks[from-1].pop(count))
		}

	}
	var result string
	for _, s := range stacks {
		result += s.crates[s.numOfCrates-1].Letter
	}
	if mode == "9000" {
		println("Day 5 - Part 1 solution:", result)
	} else if mode == "9001" {
		println("Day 5 - Part 2 solution:", result)
	}
}

// buildInitialStacks creates a list of stacks that represent the initial situation
func buildInitialStacks(scanner *bufio.Scanner) []Stack {
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		} else {
			break
		}
	}
	// last line is stack index
	// Number of stacks is (len(string)+2)/4 or len = 3*n + (n-2)
	stackIdxStr, lines := popLine(lines)
	numStacks := int((len(stackIdxStr) + 2) / 4)
	stacks := make([]Stack, numStacks)
	var line string
	for true {
		if len(lines) == 0 {
			break
		}
		line, lines = popLine(lines)
		for i := 0; i < numStacks; i++ {
			letterIdx := i*4 + 1
			if letterIdx > len(line) {
				// line ends prematurely (if no token for the remaining columns exists)
				continue
			}
			letter := string(line[letterIdx])
			if (letter < "A" || letter > "Z") && letter != " " {
				panic("Invalid letter")
			} else if letter == " " {
				// no crate in this position
				continue
			}
			stacks[i].put([]Crate{{letter}})
		}
	}
	return stacks
}

func Day5() {
	started := time.Now()
	solve("Day5/input.txt", "9000")
	solve("Day5/input.txt", "9001")
	println("Completed after", time.Since(started).String())
}
