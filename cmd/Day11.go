/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"AdventOfCode/Day11"
	"fmt"

	"github.com/spf13/cobra"
)

// Day11Cmd represents the Day11 command
var Day11Cmd = &cobra.Command{
	Use: "Day11",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Day11 called")
		Day11.Day11()
	},
}

func init() {
	rootCmd.AddCommand(Day11Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Day11Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// Day11Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
