package main

import (
	"AdventOfCode/misc"
	"fmt"
	"strconv"
	"time"
)

type Tree struct {
	height int
	north bool
	south bool
	west bool
	east bool
}

func (t *Tree) visible() bool {
	return t.north || t.west || t.south || t.east
}

func printForest(arr [][]Tree) {
	width := len(arr[0])
	height := len(arr)
	for ttb:=0; ttb<height;ttb++ {
		for ltr:=0; ltr<width;ltr++ {
			if arr[ttb][ltr].visible() {
				fmt.Print("x")
			} else {
				fmt.Print("o")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	started := time.Now()

	rows := misc.ReadInputRows("Day8/input.txt")
	width := len(rows[0])
	height := len(rows)
	arr := make([][]Tree, height)
	for idx := range arr {
		arr[idx] = make([]Tree, width)
	}
	// parse and set visible from west
	for ttb, row := range rows {
		highest := -1
		for ltr, v := range row {
			v_int, _ := strconv.Atoi(string(v))
			arr[ttb][ltr].height = v_int
			if v_int > highest {
				arr[ttb][ltr].west = true
				highest = v_int
			}
		}
	}
	// visible from north
	for ltr:=0; ltr<width;ltr++ {
		highest := -1
		for ttb:=0; ttb<height;ttb++ {
			if arr[ttb][ltr].height > highest {
				arr[ttb][ltr].north = true
				highest = arr[ttb][ltr].height
			}
		}
	}
	// visible from south
	for ltr:=width-1; ltr>=0;ltr-- {
		highest := -1
		for ttb:=height-1; ttb>=0;ttb-- {
			if arr[ttb][ltr].height > highest {
				arr[ttb][ltr].south = true
				highest = arr[ttb][ltr].height
			}
		}
	}
	// visible from east
	for ttb:=0; ttb<height;ttb++ {
		highest := -1
		for ltr:=width-1; ltr>=0;ltr-- {
			if arr[ttb][ltr].height > highest {
				arr[ttb][ltr].east = true
				highest = arr[ttb][ltr].height
			}
		}
	}
	visible := 0
	for ltr:=0; ltr<width;ltr++ {
		for ttb:=0; ttb<height;ttb++ {
			if arr[ttb][ltr].visible() {
				visible++
			}
		}
	}

	fmt.Println("Day 8 - Part 1 solution:", visible)
	printForest(arr)
	println("Completed after", time.Since(started).String())
}