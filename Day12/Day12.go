package Day12

import (
	"AdventOfCode/misc"
	"fmt"
	"image"
)

func bfs(p image.Point, s chan image.Point, grid map[image.Point]rune) {
	for _, move := range []image.Point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		if _, ok := grid[p.Add(move)]; ok {
			if testMove(p, p.Add(move), grid) {
				s <- p.Add(move)
			}
		}
	}
}

func testMove(src, dst image.Point, grid map[image.Point]rune) bool {
	var src_val, dst_val rune
	if grid[src] == 'S' {
		src_val = 'a'
		dst_val = grid[dst]
	} else if grid[dst]=='E' {
		src_val = grid[src]
		dst_val = 'z'
	} else {
		src_val = grid[src]
		dst_val = grid[dst]
	}
	return dst_val-src_val <= 1
}

func Day12() {
	txt := misc.ReadInputRows("input.txt")
	grid := map[image.Point]rune{}
	visited := map[image.Point]bool{}
	var start image.Point
	for r, line := range txt {
		for c, val := range line {
			grid[image.Point{X: r, Y: c}] = val
			if val == rune('S') {
				start = image.Point{X: r, Y: c}
			}
		}
	}
	q := []image.Point{}
	q = append(q, start)
	s := make(chan image.Point)
	var i int
	for i = 0; i < 1000; i++ {
		elements := len(q)
		for i := 0; i < elements; i++ {
			if grid[q[i]] == 'E' {
				break
			}
			if !visited[q[i]] {
				go bfs(q[i], s, grid)
				visited[q[i]] = true
			}

		}
		q = []image.Point{}
		for i := 0; i < elements; i++ {
			q = append(q, <-s)
		}
	}
	fmt.Println(start)
	fmt.Println(i)
	// ToDo: Set end to 'z'
	// ToDo: Don't move into already visited segments
}
