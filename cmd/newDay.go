/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"AdventOfCode/misc"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// newDayCmd represents the newDay command
var newDayCmd = &cobra.Command{
	Use:  "newDay",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		day, err := strconv.Atoi(args[0])
		if err != nil {
			panic("Invalid argument: " + args[0])
		}
		misc.NewDay(day)
		fmt.Println("Created structure for new day: " + args[0])
	},
}

func init() {
	rootCmd.AddCommand(newDayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newDayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newDayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
