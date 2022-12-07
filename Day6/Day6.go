package day6

import (
	"AdventOfCode/misc"
	"bufio"
	"fmt"
	"os"
	"time"
)

type PacketQueue [4]rune
type MessageQueue [14]rune

type Queue interface {
	PacketQueue | MessageQueue
}

func (q *PacketQueue) append(element rune) {
	l := len(q)
	if q[l-1] != 0 {
		// queue is already full
		l = len(q)
		for i := 0; i < (l - 1); i++ {
			q[i] = q[i+1]
		}
		q[l-1] = element
	} else {
	}
	// find last element
	lastElem := -1
	for i := 0; i < (l); i++ {
		if q[i] == 0 {
			q[lastElem+1] = element
			break
		} else {
			lastElem = i
		}
	}
}

func (q *PacketQueue) allUnique() bool {
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
func (q *MessageQueue) append(element rune) {
	l := len(q)
	if q[l-1] != 0 {
		// queue is already full
		l = len(q)
		for i := 0; i < (l - 1); i++ {
			q[i] = q[i+1]
		}
		q[l-1] = element
	} else {
	}
	// find last element
	lastElem := -1
	for i := 0; i < (l); i++ {
		if q[i] == 0 {
			q[lastElem+1] = element
			break
		} else {
			lastElem = i
		}
	}
}

func (q *MessageQueue) allUnique() bool {
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

func findPacketMarker(text string) int {
	q := PacketQueue{}
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

func findMessageMarker(text string) int {
	q := MessageQueue{}
	var res int
	for idx, r := range text {
		res = idx
		q.append(r)
		if q.allUnique() {
			break
		}
	}
	marker := res + 1
	fmt.Println("Day 6 - Part 2 solution:", marker)
	return marker
}

func Day6() {
	started := time.Now()
	text := readInput("Day6/input.txt")
	findPacketMarker(text)
	findMessageMarker(text)
	println("Completed after", time.Since(started).String())
}
