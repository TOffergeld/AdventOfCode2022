package Day07

import (
	"AdventOfCode/misc"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Directory struct {
	files       map[string]*File
	directories map[string]*Directory
	name        string
	parent      *Directory
	size        int
}

type File struct {
	name   string
	size   int
	parent *Directory
}

func (d *Directory) getSize() int {
	if d.size != 0 {
		// size is already known (doesn't work if size is calculated before tree is complete)
		return d.size
	} else {
		size := 0
		for _, dir := range d.directories {
			size += dir.getSize()
		}
		for _, f := range d.files {
			size += f.size
		}
		d.size = size
		return size
	}
}

func (d *Directory) getAllDirs() []*Directory {
	dirs := []*Directory{}
	for _, dir := range d.directories {
		dirs = append(dirs, dir)
		dirs = append(dirs, dir.getAllDirs()...)
	}
	return dirs
}

func buildTree(text []string) *Directory {
	root := new(Directory)
	currentDir := root
	currentDir.name = "/"
	currentDir.directories = map[string]*Directory{}
	currentDir.files = map[string]*File{}
	for _, row := range text {
		switch row[:4] {
		case "$ cd":
			dirname := strings.TrimPrefix(row, "$ cd ")
			if dir, ok := currentDir.directories[dirname]; ok {
				currentDir = dir
			} else {
				if dirname == ".." {
					currentDir = currentDir.parent
				} else {
					fmt.Printf(dirname, "not found.")
				}
			}
		case "$ ls":
			continue
		case "dir ":
			dirname := strings.TrimPrefix(row, "dir ")
			newDir := new(Directory)
			newDir.name = dirname
			newDir.parent = currentDir
			newDir.directories = map[string]*Directory{}
			newDir.files = map[string]*File{}
			currentDir.directories[dirname] = newDir
		default:
			x := strings.Split(row, " ")
			s, fname := x[0], x[1]
			size, _ := strconv.Atoi(s)
			f := File{fname, size, currentDir}
			currentDir.files[fname] = &f
		}
	}
	return root
}

func solve(text []string) {
	root := buildTree(text)
	root.getSize()
	all := root.getAllDirs()

	p1 := 0
	for _, dir := range all {
		if dir.size <= 100000 {
			p1 += dir.size
		}
	}
	fmt.Println("Day 7 - Part 1 solution:", p1)

	fs := 70000000
	total := root.size
	req := 30000000
	ideal := req - (fs - total)

	selected := root
	for _, dir := range all {
		if dir.size >= ideal {
			if dir.size < selected.size {
				selected = dir
			}
		}
	}
	fmt.Println("Day 7 - Part 2 solution:", selected.size)
}

func Day7() {
	started := time.Now()
	text := misc.ReadInputRows("Day07/input.txt")
	solve(text)

	println("Completed after", time.Since(started).String())
}
