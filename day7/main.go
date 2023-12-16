package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type handtype struct {
	name string
	value int
}

type hand struct {
	cards string
	bid int
	handtype
}

var handTypes = []handtype{
	{
		name: "fiveofakind",
		value: 7,
	},
	{
		name: "fourofakind",
		value: 6,
	},
	{
		name: "fullhouse",
		value: 5,
	},
	{
		name: "threeofakind",
		value: 4,
	},
	{
		name: "twopairs",
		value: 3,
	},
	{
		name: "onepair",
		value: 2,
	},
	{
		name: "highcard",
		value: 1,
	},
}

func main() {
	bs, err := os.ReadFile("./actualdata.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hands := parseHands(string(bs))
	slices.SortFunc(hands,
		func(a, b hand) int {
			return a.handtype.value - b.handtype.value
		})
	hands = sortHands(hands)
	winnings := 0
	for i := 0; i < len(hands); i++ {
		rank := i + 1
		fmt.Println("For hand: ", hands[i], ", I have rank: ", rank)
		winnings += rank * hands[i].bid
	}
	fmt.Println("Winnings: ", winnings)
}

var values = map[string]int {
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	// "J": 11, // for pt 1
	"J": 1, // for pt 2
	"Q": 12,
	"K": 13,
	"A": 14,
}

func sortHands(hs []hand) []hand {
	byhandtype := [][]hand{}
	for i := 0; i < len(hs); i++ {
		h := hs[i]
		endindex := i + 1
		for endindex < len(hs) && h.handtype.value == hs[endindex].handtype.value {
			endindex++
		}
		subsection := hs[i:endindex]
		i = endindex - 1
		byhandtype = append(byhandtype, subsection)
	}

	output := []hand{}
	for _, cardsbyhandtype := range byhandtype {
		slices.SortFunc(cardsbyhandtype,
			func(a, b hand) int {
				i := 0
				result := 0
				for result == 0 && i < len(a.cards) {
					aval := values[a.cards[i:i+1]]
					bval := values[b.cards[i:i+1]]
					result = aval - bval
					i++
				}
				return result
			},
		)
		for _, card := range cardsbyhandtype {
			output = append(output, card)
		}
	}

	return output
}

func parseHands(puzzleInput string) []hand {
	hands := []hand{}
	rows := strings.Split(puzzleInput, "\n")
	fmt.Println("Rows length: ", len(rows))
	for _, row := range rows {
		fields := strings.Fields(row)
		h, b := fields[0], fields[1]
		bid, _ := strconv.Atoi(b)
		handtype := getHandTypePt2(h)
		hand := hand{
			cards: h,
			bid: bid,
			handtype: handtype,
		}
		hands = append(hands, hand)
	}
	return hands
}

func getHandTypePt2(h string) handtype {
	cards := map[string]int{}
	for i := 0; i < len(h); i++ {
		c := h[i:i + 1]
		_, ok := cards[c]
		if ok {
			cards[c] += 1
		} else {
			cards[c] = 1
		}
	}
	_, ok := cards["J"]
	if ok {
		// fmt.Println("Has jack...", cards)
		njs := cards["J"]
		delete(cards, "J")
		keyWithHighestValue := ""
		highestValue := -1
		for k, v := range cards {
			if v > highestValue {
				highestValue = v
				keyWithHighestValue = k
			}
		}
		// fmt.Println(keyWithHighestValue, highestValue)
		cards[keyWithHighestValue] += njs
	}

	// fmt.Println(cards)
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


// For pt 1
func getHandType(h string) handtype {
	cards := map[string]int{}
	for i := 0; i < len(h); i++ {
		c := h[i:i + 1]
		_, ok := cards[c]
		if ok {
			cards[c] += 1
		} else {
			cards[c] = 1
		}
	}

	// High card
	if len(cards) == 5 {
		return handTypes[6]
	} else if len(cards) == 4 { // one pair
		return handTypes[5]
	} else if len(cards) == 1 { // fiveof a kind

		fmt.Println("FIVE OF A KIND", cards)
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

