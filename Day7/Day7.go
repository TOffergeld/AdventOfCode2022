package day5

import (
	"time"
)

func Day7() {
	started := time.Now()
	solve("Day7/input.txt")
	solve("Day7/input.txt")
	println("Completed after", time.Since(started).String())
}
