package cmd

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/spf13/cobra"
)

// day01Cmd represents the day01 command
var day01Cmd = &cobra.Command{
	Use:   "day01",
	Short: "Day01",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(day01Cmd)
	day01Cmd.PersistentFlags().String("input", "input/1_1.txt", "input file (default is ./input/1_1.txt)")
	day01Cmd.AddCommand(p1Cmd)
	day01Cmd.AddCommand(p2Cmd)
}

var p1Cmd = &cobra.Command{
	Use:   "p1",
	Short: "Day01 Part 1",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("input")
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}

		result, err := day01_1(name)
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}
		cmd.Printf("Result: %s\n", result)

	},
}

var p2Cmd = &cobra.Command{
	Use:   "p2",
	Short: "Day01 Part 2",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("input")
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}

		result, err := day01_2(name)
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}
		cmd.Printf("Result: %s\n", result)
	},
}

func day01_1(file string) (string, error) {
	data, err := lib.ReadInputLines(file)
	if err != nil {
		return "", err
	}

	sonarData, err := lib.ParseLines2Int(data)
	if err != nil {
		return "", err
	}

	var prev int64 = sonarData[0]
	var counter int64 = 0
	for i := 1; i < len(sonarData); i++ {
		if sonarData[i] > prev {
			counter++
		}
		prev = sonarData[i]
	}

	return fmt.Sprintf("%d", counter), nil

}

func day01_2(file string) (string, error) {
	data, err := lib.ReadInputLines(file)
	if err != nil {
		return "", err
	}

	sonarData, err := lib.ParseLines2Int(data)
	if err != nil {
		return "", err
	}

	var prev int64 = sonarData[0] + sonarData[1] + sonarData[2]
	var counter int64 = 0
	for i := 1; i < len(sonarData)-2; i++ {
		window := sonarData[i] + sonarData[i+1] + sonarData[i+2]
		if window > prev {
			counter++
		}
		prev = window
	}

	return fmt.Sprintf("%d", counter), nil

}
