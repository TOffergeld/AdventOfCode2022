package Day6

import (
	"AdventOfCode/misc"
	"bufio"
	"fmt"
	"os"
	"time"
)

type Queue []rune

func (q *Queue) append(element rune) {
	l := cap(*q)
	if (*q)[l-1] != 0 {
		// queue is already full
		for i := 0; i < (l - 1); i++ {
			(*q)[i] = (*q)[i+1]
		}
		(*q)[l-1] = element
	} else {
	}
	// find last element
	lastElem := -1
	for i := 0; i < (l); i++ {
		if (*q)[i] == 0 {
			(*q)[lastElem+1] = element
			break
		} else {
			lastElem = i
		}
	}
}

func (q *Queue) allUnique() bool {
	isin := map[rune]bool{}
	for _, el := range *q {
		if el == 0 {
			// queue is not filled yet
			return false
		}
		if !isin[el] {
			isin[el] = true
		} else {
			return false
		}
	}
	return true
}

func readInput(path string) string {
	f := misc.GetInputFile(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	scanner := bufio.NewScanner(f)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	return text
}

func findMarker(text string, length int) int {
	q := make(Queue, length)
	var res int
	for idx, r := range text {
		res = idx
		q.append(r)
		if q.allUnique() {
			break
		}
	}
	marker := res + 1
	fmt.Println("Day 6 - Part 1 solution:", marker)
	return marker
}

func Day6() {
	started := time.Now()
	text := readInput("Day6/input.txt")
	findMarker(text, 4)
	findMarker(text, 14)
	println("Completed after", time.Since(started).String())
}
