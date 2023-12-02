package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type value struct {
	numString string
	stringRep string
}

var possibleValues = []value{
	{
		numString: "1",
		stringRep: "one",
	},
	{
		numString: "2",
		stringRep: "two",
	},
	{
		numString: "3",
		stringRep: "three",
	},
	{
		numString: "4",
		stringRep: "four",
	},
	{
		numString: "5",
		stringRep: "five",
	},
	{
		numString: "6",
		stringRep: "six",
	},
	{
		numString: "7",
		stringRep: "seven",
	},
	{
		numString: "8",
		stringRep: "eight",
	},
	{
		numString: "9",
		stringRep: "nine",
	},
}

func main() {
	filepath := "./actualdata.txt"
	bs, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}

	data := string(bs)
	rows := strings.Split(data, "\n")
	sum := 0
	for _, row := range rows {
		fmt.Println(row)
		v1 := findFirstDigitsPt2FromLeft(row)
		v2 := findFirstDigitsPt2FromRight(row)
		fmt.Println(v1 + v2)

		n, err := strconv.Atoi(v1 + v2)
		if err != nil {
			fmt.Println("WTF: ", err)
			os.Exit(1)
		}
		sum += n
	}

	fmt.Println("sum: ", sum)
}

// For part 2
func findFirstDigitsPt2FromLeft(s string) string {
	firstIndex := len(s) + 1

	var foundItem value

	for _, pv := range possibleValues {
		i := strings.Index(s, pv.numString)
		if (i >= 0 && i < firstIndex) {
			fmt.Printf("Found new furthest to the left for '%s' on index %d\n", pv.numString, i)
			foundItem = pv
			firstIndex = i
		}

		k := strings.Index(s, pv.stringRep)
		if (k >= 0 && k < firstIndex) {
			fmt.Printf("Found new furthest to the left for '%s' on index %d\n", pv.stringRep, k)
			foundItem = pv
			firstIndex = k
		}
	}

	return foundItem.numString
}

// For part 2
func findFirstDigitsPt2FromRight(s string) string {
	firstIndex := -1

	var foundItem value

	for _, pv := range possibleValues {
		i := strings.LastIndex(s, pv.numString)
		if (i > firstIndex) {
			fmt.Printf("Found new furthest to the right for '%s' on index %d\n", pv.numString, i)
			foundItem = pv
			firstIndex = i
		}

		k := strings.LastIndex(s, pv.stringRep)
		if (k > firstIndex) {
			fmt.Printf("Found new furthest to the right for '%s' on index %d\n", pv.stringRep, k)
			foundItem = pv
			firstIndex = k
		}
	}

	return foundItem.numString
}

// For part 1
func findFirstDigit(s string, fromLeft bool) string {
	var i int
	n := -1

	if (fromLeft) {
		i = 0
		for n == -1 && i < len(s) {
			c := s[i:i+1]
			num, err := strconv.Atoi(c)
			if err == nil {
				n = num
			}
			i++
		}
	} else {
		i = len(s) - 1
		for n == -1 && i >= 0 {
			c := s[i:i+1]
			num, err := strconv.Atoi(c)
			if err == nil {
				n = num
			}
			i--
		}
	}

	return strconv.Itoa(n)
}