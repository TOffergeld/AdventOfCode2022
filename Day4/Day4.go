package day4

import (
	"AdventOfCode/misc"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part1(path string) {
	f := misc.GetInputFile(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		item := [2][2]int{
			stringToInt(*(*[2]string)(strings.Split(s[0], "-"))),
			stringToInt(*(*[2]string)(strings.Split(s[1], "-")))}
		if (item[0][0] <= item[1][0] && item[0][1] >= item[1][1]) ||
			(item[0][0] >= item[1][0] && item[0][1] <= item[1][1]) {
			count++
		}
	}
	println("Day 4 - Part 1 solution:", count)
}

func part2(path string) {
	f := misc.GetInputFile(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		item := [2][2]int{
			stringToInt(*(*[2]string)(strings.Split(s[0], "-"))),
			stringToInt(*(*[2]string)(strings.Split(s[1], "-")))}
		if checkForOverlap(item) {
			count++
		}
	}
	println("Day 4 - Part 2 solution:", count)
}

func checkForOverlap(pair [2][2]int) bool {
	A, B := pair[0], pair[1]
	if (A[0] <= B[0] && B[0] <= A[1]) || // A[0] <= B[0] <= A[1]
		(A[0] <= B[1] && B[1] <= A[1]) || // A[0] <= B[1] <= A[1]
		(B[0] <= A[0] && A[0] <= B[1]) || // B[0] <= A[0] <= B[1]
		(B[0] <= A[1] && A[1] <= B[1]) { // B[0] <= A[1] <= B[1]
		return true
	} else {
		return false
	}
}

func stringToInt(in [2]string) [2]int {
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
	part2("Day4/input.txt")
}
