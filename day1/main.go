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
	rows := strings.Split(data, "\n")
	sum := 0
	for _, row := range rows {
		sn1 := findFirstDigit(row, true)
		sn2 := findFirstDigit(row, false)
		n, err := strconv.Atoi(sn1 + sn2)
		if err != nil {
			fmt.Println("WTF: ", err)
			os.Exit(1)
		}
		sum += n
	}


	fmt.Println("sum: ", sum)
}

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