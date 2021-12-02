package cmd

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/spf13/cobra"
)

// day01Cmd represents the day01 command
var day02Cmd = &cobra.Command{
	Use:   "day02",
	Short: "Day02",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(day02Cmd)
	day02Cmd.PersistentFlags().String("input", "input/2_1.txt", "input file (default is ./input/2_1.txt)")
	day02Cmd.AddCommand(d2p1Cmd)
	day02Cmd.AddCommand(d2p2Cmd)
}

var d2p1Cmd = &cobra.Command{
	Use:   "p1",
	Short: "Day02 Part 1",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("input")
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}

		result, err := day02_1(name)
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}
		cmd.Printf("Result: %s\n", result)

	},
}

var d2p2Cmd = &cobra.Command{
	Use:   "p2",
	Short: "Day02 Part 2",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("input")
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}

		result, err := day02_2(name)
		if err != nil {
			cmd.PrintErrf("error: %+v", err)
		}
		cmd.Printf("Result: %s\n", result)
	},
}

func day02_1(file string) (string, error) {
	data, err := lib.ReadInputLines(file)
	if err != nil {
		return "", err
	}

	hor, depth := lib.ParseLines2Nav1(data)

	return fmt.Sprintf("%d", hor*depth), nil
}

func day02_2(file string) (string, error) {
	data, err := lib.ReadInputLines(file)
	if err != nil {
		return "", err
	}

	hor, depth := lib.ParseLines2Nav2(data)

	return fmt.Sprintf("%d", hor*depth), nil
}
