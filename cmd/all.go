/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	day1 "AdventOfCode/Day01"
	day2 "AdventOfCode/Day02"
	day3 "AdventOfCode/Day03"
	day4 "AdventOfCode/Day04"
	day5 "AdventOfCode/Day05"
	day6 "AdventOfCode/Day06"
	"AdventOfCode/Day07"
	"AdventOfCode/Day08"
	"AdventOfCode/Day09"
	"AdventOfCode/Day10"
	"fmt"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var allCmd = &cobra.Command{
	Use: "all",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called")
		day1.Day1()
		day2.Day2()
		day3.Day3()
		day4.Day4()
		day5.Day5()
		day6.Day6()
		Day07.Day7()
		Day08.Day8()
		Day09.Day9()
		Day10.Day10()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
