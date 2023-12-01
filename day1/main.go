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
	filepath := "./testdata2.txt"
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
	firstIndex := 1000

	var foundItem value

	for _, pv := range possibleValues {
		i := strings.Index(s, pv.numString)
		k := strings.Index(s, pv.stringRep)
		if ((i >= 0 && i < firstIndex) || (k >= 0 && k < firstIndex)) {
			foundItem = pv
			if ((i < k && i >= 0) || k < 0) {
				firstIndex = i
			} else {
				firstIndex = k
			} 
		}
	}

	return foundItem.numString
}

// For part 2
func findFirstDigitsPt2FromRight(s string) string {
	firstIndex := 0

	var foundItem value

	for _, pv := range possibleValues {
		i := strings.LastIndex(s, pv.numString)
		k := strings.LastIndex(s, pv.stringRep)
		if ((i < len(s) && i > firstIndex) || (k < len(s) && k > firstIndex)) {
			foundItem = pv
			if (i > k) {
				firstIndex = i
			} else  {
				firstIndex = k
			}
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