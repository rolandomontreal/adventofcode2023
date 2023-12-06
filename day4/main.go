package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type lotteryticket struct {
	num int
	winningNums []int
	givenNums []int
	copies int
}

func main() {
	filepath := "./testdata.txt"
	bs, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Could not read file...", err)
		os.Exit(1)
	}

	data := string(bs)
	rows := strings.Split(data, "\n")

	// tickets := parseTickets(data)
	// fmt.Println(tickets)

	sum := 0
	for _, row := range rows {
		fmt.Println(row)
		score := getScoreRow(row)
		fmt.Println("Score calculated: ", score, "\n")
		sum += score
	}

	fmt.Println("Solution: ", sum)
}

func parseTickets(data string) []lotteryticket {
	tickets := []lotteryticket{}
	rows := strings.Split(data, "\n")[0:1]
	for i, row := range rows {
		nums := strings.Split(strings.Split(row, ":")[1], "|")
		winningNumsStr, givenNumsStr := nums[0], nums[1]
		wns := parseNums(winningNumsStr)
		gns := parseNums(givenNumsStr)
		ticket := lotteryticket{
			num: i + 1,
			winningNums: wns,
			givenNums: gns,
			copies: 0,
		}
		tickets = append(tickets, ticket)
	}
	return tickets
}

func parseNums(numarr string) []int {
	nstr := strings.Fields(numarr)
	fmt.Println(nstr)
	output := []int{}
	for _, ns := range nstr {
		n, err := strconv.Atoi(ns)
		if err != nil {
			fmt.Println("Fail parsing num: ", n, err)
			os.Exit(1)
		}
		output = append(output, n)
	}
	return output
}

func getScoreRow(row string) int {
	nums := strings.Split(strings.Split(row, ":")[1], "|")
	winningNumsStr, givenNumsStr := nums[0], nums[1]
	winningsNums := strings.Split(strings.TrimSpace(winningNumsStr), " ")

	winningNumsSet := map[int]bool {}
	for _, winningNum := range winningsNums {
		if len(winningNum) > 0 {
			wn, err := strconv.Atoi(winningNum)
			if err != nil {
				fmt.Println("Could not parse wn: ", wn, err)
				os.Exit(1)
			}
			winningNumsSet[wn] = true
		}
	}
	fmt.Println("Numsset: ", winningNumsSet)

	givenNums := strings.Split(strings.TrimSpace(givenNumsStr), " ")
	occurrences := 0
	for _, givenNum := range givenNums {
		if len(givenNum) > 0 {
			gn, err := strconv.Atoi(givenNum)
			if err != nil {
				fmt.Println("Could not parse wn: ", gn, err)
				os.Exit(1)
			}
			if winningNumsSet[gn] {
				fmt.Printf("Match for %d\n", gn)
				occurrences++
			}
		}
	}

	fmt.Println("Occurences: ", occurrences)

	return int(math.Pow(2, float64(occurrences - 1)))
}