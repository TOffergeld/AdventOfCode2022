package dayone

import (
	"AdventOfCode/misc"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

type Elf struct {
	Position    int
	Possessions int
}

// sumPossessions calculates the sum of each elfs possessions from a given slice of integers
func sumPossessions(possessions []int) int {
	sum := 0
	for _, v := range possessions {
		sum += v
	}
	return sum
}

// readInput reads the input txt file for the challenge.
func readInput(path string) []Elf {
	f := misc.GetInputFile(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	position := 1

	elfs := []Elf{}
	stack := []int{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			i_line, _ := strconv.Atoi(line)
			stack = append(stack, i_line)
		} else {
			elfs = append(elfs,
				Elf{position, sumPossessions(stack)})
			position += 1
			stack = nil
		}
	}
	return elfs
}

// findMax finds the maximum of possessions of the elves
func findMax(elfs []Elf) Elf {
	var max Elf
	for _, elf := range elfs {
		if elf.Possessions > max.Possessions {
			max = elf
		}
	}
	return max

}

func sortElfs(elfs []Elf) {
	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].Possessions > elfs[j].Possessions
	})
}

func Day1() {
	started := time.Now()
	elfs := readInput("Day1/input.txt")
	sortElfs(elfs)
	fmt.Println("Day 1 - Part 1 solution:", elfs[0].Possessions)
	fmt.Println("Day 1 - Part 2 solution:",
		elfs[0].Possessions+
			elfs[1].Possessions+
			elfs[2].Possessions)
	println("Completed after", time.Since(started).Milliseconds())
}
