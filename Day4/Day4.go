package day4

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	Ranges [2][2]int
}

func part1(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)
	part1 := 0
	part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		item := Pair{Ranges: [2][2]int{
			toInt(*(*[2]string)(strings.Split(s[0], "-"))),
			toInt(*(*[2]string)(strings.Split(s[1], "-")))}}
		if (item.Ranges[0][0] <= item.Ranges[1][0] && item.Ranges[0][1] >= item.Ranges[1][1]) ||
			(item.Ranges[0][0] >= item.Ranges[1][0] && item.Ranges[0][1] <= item.Ranges[1][1]) {
			part1++
		}
		if overlap(item) {
			part2++
		}
	}
	println("Day 4 - Part 1 solution:", part1)
	println("Day 4 - Part 2 solution:", part2)
}

func overlap(p Pair) bool {
	A, B := p.Ranges[0], p.Ranges[1]
	if (A[0] <= B[0] && B[0] <= A[1]) ||
		(A[0] <= B[1] && B[1] <= A[1]) ||
		(B[0] <= A[0] && A[0] <= B[1]) ||
		(B[0] <= A[1] && A[1] <= B[1]) {
		return true
	} else {
		return false
	}
}

func toInt(in [2]string) [2]int {
	res := [2]int{}
	var err error
	res[0], err = strconv.Atoi(in[0])
	if err != nil {
		panic(err)
	}
	res[1], err = strconv.Atoi(in[1])
	if err != nil {
		panic(err)
	}
	return res
}

func Day4() {
	part1("Day4/input.txt")
}
