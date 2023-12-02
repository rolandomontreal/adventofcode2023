package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filepath := "./actualdata.txt"
	bs, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	data := string(bs)

	restriction := map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	rows := strings.Split(data, "\n")

	sumOfPossibleRows := 0
	sumOfMinimumPowers := 0
	for _, row := range rows {
		possibleRow, id := possibleSet(row, restriction)
		if possibleRow {
			sumOfPossibleRows += id
		}

		// Least possible combination
		lpc := getLeastPossibleCombination(row)
		power := lpc["red"] * lpc["green"] * lpc["blue"]
		sumOfMinimumPowers += power
	}
	fmt.Println("Result: ", sumOfMinimumPowers)
}

func getLeastPossibleCombination(row string) map[string]int {
	output := map[string]int {
		"red": 0,
		"green": 0,
		"blue": 0,
	}

	splitted := strings.Split(row, ":")

	_, joinedSets := splitted[0], splitted[1]

	sets := strings.Split(joinedSets, ";")

	for _, set := range sets {
		perColor := strings.Split(set, ",")
		
		for _, color := range perColor {
			c := strings.TrimSpace(color)
			splitted := strings.Split(c, " ")
			n, err := strconv.Atoi(splitted[0])
			if err != nil {
				fmt.Println("Could not count instances in set")
				os.Exit(1)
			}
			color := splitted[1]
			if output[color] < n {
				output[color] = n
			}
		}
	}

	return output
}

func possibleSet(row string, restrictions map[string]int) (bool, int) {
	splitted := strings.Split(row, ":")

	game, joinedSets := splitted[0], splitted[1]
	gameId, err := strconv.Atoi(strings.Split(game, " ")[1])
	if err != nil {
		fmt.Println("Could not parse game it: ", err)
		os.Exit(1)
	}

	sets := strings.Split(joinedSets, ";")

	isPossible := true

	for _, set := range sets {
		perColor := strings.Split(set, ",")
		
		for _, color := range perColor {
			c := strings.TrimSpace(color)
			splitted := strings.Split(c, " ")
			n, err := strconv.Atoi(splitted[0])
			if err != nil {
				fmt.Println("Could not count instances in set")
				os.Exit(1)
			}
			color := splitted[1]
			restriction := restrictions[color]
			if n > restriction {
				isPossible = false
			}
		}
	}

	return isPossible, gameId
}