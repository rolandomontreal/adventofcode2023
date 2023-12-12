package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type handtype struct {
	name string
	rank int
}

type hand struct {
	cards string
	bid int
	handtype
}

var handTypes = []handtype{
	{
		name: "fiveofakind",
		rank: 7,
	},
	{
		name: "fourofakind",
		rank: 6,
	},
	{
		name: "fullhouse",
		rank: 5,
	},
	{
		name: "threeofakind",
		rank: 4,
	},
	{
		name: "twopairs",
		rank: 3,
	},
	{
		name: "onepair",
		rank: 2,
	},
	{
		name: "highcard",
		rank: 1,
	},
}

func main() {
	bs, err := os.ReadFile("./testdata.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hands := parseHands(string(bs))
	fmt.Println(hands)
}

func sortHands(hs []hand) []hand {
	output := []hand{}
	// TODO - implement sorting here
	return output
}

func parseHands(puzzleInput string) []hand {
	hands := []hand{}
	rows := strings.Split(puzzleInput, "\n")
	for _, row := range rows {
		fields := strings.Fields(row)
		h, b := fields[0], fields[1]
		bid, _ := strconv.Atoi(b)
		fmt.Println(h, b)
		handtype := getHandType(h)
		fmt.Println(handtype, "\n")
		hand := hand{
			cards: h,
			bid: bid,
			handtype: handtype,
		}
		hands = append(hands, hand)
	}
	return hands
}

func getHandType(h string) handtype {
	

	fmt.Println("Running get hand type for: ", h)
	cards := map[string]int{}
	for i := 0; i < len(h); i++ {
		c := h[i:i + 1]
		_, ok := cards[c]
		if ok {
			fmt.Println("Has key, will increment")
			cards[c] += 1
		} else {
			fmt.Println("Does not have..")
			cards[c] = 1
		}
	}

	// High card
	if len(cards) == 5 {
		return handTypes[6]
	} else if len(cards) == 4 { // one pair
		return handTypes[5]
	} else if len(cards) == 1 { // fiveof a kind
		return handTypes[0]
	}

	// Four of a kind or full house
	if len(cards) == 2 {
		fourofakind := true
		for _, v := range cards {
			if v == 3 {
				fourofakind = false
			}
		}
		if fourofakind {
			return handTypes[1]
		}
		return handTypes[2]
	} else { // Length is three, three of a kind or pairs
		threeofakind := true
		for _, v := range cards {
			if v == 2 {
				threeofakind = false
			}
		}
		if threeofakind {
			return handTypes[3]
		}
	}

	return handTypes[4]
}

