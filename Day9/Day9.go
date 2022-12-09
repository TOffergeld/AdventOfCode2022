package Day9

import (
	"AdventOfCode/misc"
	"fmt"
	"image"
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

func sign(v int) int {
	if v >= 0 {
		return 1
	} else {
		return -1
	}
}

// pInR determines if a point lies within a rectangle (including all borders - Point.In is half-open, (1,1) isn't contained in (-1, -1, 1, 1)
func pInR(p image.Point, r image.Rectangle) bool {
	return r.Min.X <= p.X && p.X <= r.Max.X &&
		r.Min.Y <= p.Y && p.Y <= r.Max.Y
}

var directions = map[string]image.Point{
	"R": {1, 0},
	"D": {0, -1},
	"L": {-1, 0},
	"U": {0, 1},
}

// updateKnot calculates the new position for tail knot t, influenced by head knot h.
func updateKnot(h image.Point, t image.Point) image.Point {
	distance := t.Sub(h)
	rect := image.Rect(-1, -1, 1, 1)
	if !pInR(distance, rect) {
		if abs(distance.X)+abs(distance.Y) > 2 {
			// move diagonally
			t = t.Add(image.Point{-sign(distance.X), -sign(distance.Y)})
		} else {
			// move horizontally or vertically
			tail_move := distance.Div(abs(distance.X) + abs(distance.Y))
			t = t.Sub(tail_move)
		}
	}
	return t
}

// visitedKnot simulates the movement of the entire rope according to the moves specified in txt and returns a map of points visited by the tail knot of the rope.
func visitedNodes(txt []string, rope []image.Point) map[image.Point]bool {
	visited := map[image.Point]bool{}
	h := image.Point{0, 0}
	t := h
	visited[t] = true

	for _, line := range txt {
		instruction := strings.Split(line, " ")
		dir := instruction[0]
		count, _ := strconv.Atoi(instruction[1])

		// split up into individual moves
		for move := 1; move <= count; move++ {
			h = rope[0]
			t = rope[1]
			// make head move first and update second segment
			h = h.Add(directions[dir])
			t = updateKnot(h, t)
			rope[0] = h
			rope[1] = t
			// move along the rope and update each segment
			for segment := range rope[1:] {
				// update head and tail and move tail
				h = rope[segment]
				t = rope[segment+1]
				t = updateKnot(h, t)
				rope[segment+1] = t
			}
			visited[rope[len(rope)-1]] = true
		}
	}
	return visited
}

func Day9() {
	started := time.Now()
	txt := misc.ReadInputRows("Day9/input.txt")

	shortRope := make([]image.Point, 2)
	longRope := make([]image.Point, 10)

	visited := visitedNodes(txt, shortRope)
	fmt.Println("Day 9 - Part 1 solution:", len(visited))

	visited = visitedNodes(txt, longRope)
	fmt.Println("Day 9 - Part 2 solution:", len(visited))
	println("Completed after", time.Since(started).String())

}
