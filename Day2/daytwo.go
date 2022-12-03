package daytwo

import (
	"bufio"
	"os"
	"strings"
)

type Game struct {
	Opponent string
	Own      string
	Points   int
}

func (g *Game) calculateScore(score chan int) {
	switch g.Own {
	case "X":
		g.Points += 1
		if g.Opponent == "C" {
			g.Points += 6
		} else if g.Opponent == "A" {
			g.Points += 3
		}
	case "Y":
		g.Points += 2
		if g.Opponent == "A" {
			g.Points += 6
		} else if g.Opponent == "B" {
			g.Points += 3
		}
	case "Z":
		g.Points += 3
		if g.Opponent == "B" {
			g.Points += 6
		} else if g.Opponent == "C" {
			g.Points += 3
		}
	}
	score <- g.Points
}

// readInput reads the input txt file for the challenge.
func readInput(path string) []Game {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
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

func Daytwo() {
	games := readInput("Day2/input.txt")
	score := 0
	scorechan := make(chan int)
	for idx := range games {
		go games[idx].calculateScore(scorechan)
		score += <-scorechan
	}

	println("Day 2 - Part 1 solution:", score)
}
