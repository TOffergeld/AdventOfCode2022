/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	day7 "AdventOfCode/Day7"
	"fmt"

	"github.com/spf13/cobra"
)

// day7Cmd represents the day5 command
var day7Cmd = &cobra.Command{
	Use: "day7",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day7 called")
		day7.Day7()
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
