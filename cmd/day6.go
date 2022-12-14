/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	day6 "AdventOfCode/Day06"
	"fmt"

	"github.com/spf13/cobra"
)

// day6Cmd represents the day5 command
var day6Cmd = &cobra.Command{
	Use: "day6",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day5 called")
		day6.Day6()
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
