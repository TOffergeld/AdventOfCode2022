package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
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
func findMax(elfs []Elf) int {
	max := 0
	for _, elf := range elfs {
		if elf.Possessions > max {
			max = elf.Possessions
		}
	}
	return max

}

func main() {
	elfs := readInput("Day1/input.txt")
	fmt.Printf("%+v\n", elfs)
	fmt.Println("Maximum:", findMax(elfs))
}
