package main

import (
	"fmt"
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
	filepath := "./actualdata.txt"
	bs, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Could not read file...", err)
		os.Exit(1)
	}

	data := string(bs)
	tickets := parseTickets(data)

	for i, ticket := range tickets {
		for k := 0; k < ticket.copies + 1; k++ {
			extraTickets := calculateExtraTickets(ticket)
			for j := i + 1; j < i + 1 + extraTickets && j < len(tickets); j++ {
				tickets[j].copies += 1
			}
		}
	}

	sum := 0
	for _, ticket := range tickets {
		sum += 1 + ticket.copies
	}
	fmt.Println(sum)
}

func calculateExtraTickets(t lotteryticket) int {
	winningNumsSet := map[int]bool {}
	for _, winningNum := range t.winningNums {
		winningNumsSet[winningNum] = true
	}

	occurrences := 0
	for _, givenNum := range t.givenNums {
		if winningNumsSet[givenNum] {
			occurrences++
		}
	}

	return occurrences
}

func parseTickets(data string) []lotteryticket {
	tickets := []lotteryticket{}
	rows := strings.Split(data, "\n")
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

// Part 1
// func getScoreRow(row string) int {
// 	nums := strings.Split(strings.Split(row, ":")[1], "|")
// 	winningNumsStr, givenNumsStr := nums[0], nums[1]
// 	winningsNums := strings.Split(strings.TrimSpace(winningNumsStr), " ")

// 	winningNumsSet := map[int]bool {}
// 	for _, winningNum := range winningsNums {
// 		if len(winningNum) > 0 {
// 			wn, err := strconv.Atoi(winningNum)
// 			if err != nil {
// 				fmt.Println("Could not parse wn: ", wn, err)
// 				os.Exit(1)
// 			}
// 			winningNumsSet[wn] = true
// 		}
// 	}
// 	fmt.Println("Numsset: ", winningNumsSet)

// 	givenNums := strings.Split(strings.TrimSpace(givenNumsStr), " ")
// 	occurrences := 0
// 	for _, givenNum := range givenNums {
// 		if len(givenNum) > 0 {
// 			gn, err := strconv.Atoi(givenNum)
// 			if err != nil {
// 				fmt.Println("Could not parse wn: ", gn, err)
// 				os.Exit(1)
// 			}
// 			if winningNumsSet[gn] {
// 				fmt.Printf("Match for %d\n", gn)
// 				occurrences++
// 			}
// 		}
// 	}

// 	fmt.Println("Occurences: ", occurrences)

// 	return int(math.Pow(2, float64(occurrences - 1)))
// }