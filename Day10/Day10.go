package Day10

import (
	"AdventOfCode/misc"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func printSignal(signal []int) {
	for _, row := range []int{0, 1, 2, 3, 4, 5} {
		offset := 40 * row
		for col := 0; col < 40; col++ {
			var char string
			if abs(signal[offset+col]-col) <= 1 {
				char = "#"
			} else {
				char = "."
			}
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
}

func Day10() {
	started := time.Now()

	txt := misc.ReadInputRows("Day10/input.txt")
	var signal []int
	var inc int
	for _, row := range txt {
		switch row[:4] {
		case "addx":
			var items []int
			if len(signal) == 0 {
				items = []int{1, 1}
			} else {
				items = []int{signal[len(signal)-1], signal[len(signal)-1]}
			}
			for x := range items {
				items[x] += inc
			}
			signal = append(signal, items...)
			inc, _ = strconv.Atoi(strings.TrimPrefix(row, "addx "))
		case "noop":
			if len(signal) == 0 {
				signal = append(signal, 1)
			} else {
				signal = append(signal, signal[len(signal)-1]+inc)
			}
			inc = 0
		}
	}
	totals := 0
	for _, idx := range []int{20, 60, 100, 140, 180, 220} {
		//fmt.Println("Register at position", idx, "is", signal[idx-1], "Result:", idx*signal[idx-1])
		totals += idx * signal[idx-1]
	}

	fmt.Println("Day 10 - Part 1 solution:", totals)
	fmt.Println("Day 10 - Part 2 solution:")

	printSignal(signal)

	println("Completed after", time.Since(started).String())

}
