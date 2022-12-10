package Day02

import (
	"AdventOfCode/misc"
	"bufio"
	"os"
	"strings"
	"time"
)

type Game struct {
	Opponent string
	Own      string
	Points   int
}

func resultPoints(p1 string, opponent string) int {
	switch p1 {
	case "A":
		switch opponent {
		case "A":
			return 3
		case "B":
			return 0
		case "C":
			return 6
		}
	case "B":
		switch opponent {
		case "A":
			return 6
		case "B":
			return 3
		case "C":
			return 0
		}
	case "C":
		switch opponent {
		case "A":
			return 0
		case "B":
			return 6
		case "C":
			return 3
		}
	}
	return 0
}

func choicePoints(choice string) int {
	switch choice {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}
	return 0
}

func (g *Game) calculateScore(score chan int) {
	g.Points += resultPoints(g.Own, g.Opponent)
	g.Points += choicePoints(g.Own)
	score <- g.Points
}

func (g *Game) translate_p1() {
	switch g.Own {
	case "X":
		g.Own = "A"
	case "Y":
		g.Own = "B"
	case "Z":
		g.Own = "C"
	}
}

func (g *Game) translate_p2() {
	switch g.Own {
	case "X":
		switch g.Opponent {
		case "A":
			g.Own = "C"
		case "B":
			g.Own = "A"
		case "C":
			g.Own = "B"
		}
	case "Y":
		g.Own = g.Opponent
	case "Z":
		switch g.Opponent {
		case "A":
			g.Own = "B"
		case "B":
			g.Own = "C"
		case "C":
			g.Own = "A"
		}
	}
}

// readInput reads the input txt file for the challenge.
func readInput(path string) []Game {
	f := misc.GetInputFile(path)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var games []Game

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, " ")
		games = append(games, Game{letters[0], letters[1], 0})
	}
	return games
}

func part1() {
	games := readInput("Day02/input.txt")
	for idx := range games {
		go games[idx].translate_p1()
	}
	score := 0
	scorechan := make(chan int)
	for idx := range games {
		go games[idx].calculateScore(scorechan)
		score += <-scorechan
	}

	println("Day 2 - Part 1 solution:", score)
}

func part2() {
	games := readInput("Day02/input.txt")
	for idx := range games {
		go games[idx].translate_p2()
	}
	score := 0
	scorechan := make(chan int)
	for idx := range games {
		go games[idx].calculateScore(scorechan)
		score += <-scorechan
	}

	println("Day 2 - Part 2 solution:", score)
}

func Day2() {
	started := time.Now()
	part1()
	part2()
	println("Completed after", time.Since(started).String())
}
