package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filepath := "./actualdata.txt"
	bs, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := string(bs)
	// fmt.Println(data)
	rows := strings.Split(data, "\n")
	fmt.Println(rows)
	columns := len(rows[0])
	fmt.Println("columns: ", columns)

	var numberRegex = regexp.MustCompile(`\d+`)

	tests := "...123123%\n"
	otherSymbolRegex := regexp.MustCompile(`[^0-9.\n]`)
	if (otherSymbolRegex.MatchString(tests)) {
		fmt.Println("MATCH")
	} else {
		fmt.Println("NO MATCH")
	}

	sum := 0
	for y, row := range rows {
		fmt.Println("Row: ", row)
		bs := []byte(row)
		numsIndices := numberRegex.FindAllIndex(bs, -1)
		// fmt.Println("numsIndices: ", numsIndices)
		for _, match := range numsIndices {
			// fmt.Println(match)
			startX := max(0, match[0] - 1)
			// fmt.Println("Startx: ", startX)
			endX := min(columns, match[1] + 1)
			// fmt.Println("endX: ", endX)
			var upperrow string
			if (y > 0) {
				upperrow = rows[y-1][startX:endX]
				// fmt.Println("upperrow: ", upperrow)
			}
			rowsection := row[startX:endX]
			// fmt.Println("rowSection: ", rowsection)
			var lowerrow string
			if (y < len(rows) - 1) {
				lowerrow = rows[y+1][startX:endX]
				// fmt.Println("lowerrow: ", lowerrow)
			}
			section := upperrow + "\n" + rowsection + "\n" + lowerrow
			fmt.Println("\n" + section)
			otherSymbolRegex := regexp.MustCompile(`[^0-9.\n]`)
			if (otherSymbolRegex.MatchString(section)) {
				fmt.Println("VALID NUMBER")
				nstr := row[match[0]:match[1]]
				fmt.Println(nstr)
				n, err := strconv.Atoi(nstr)
				if err != nil {
					fmt.Println("Error parsing shit: ", err)
					os.Exit(1)
				}
				sum += n
			} else {
				fmt.Println("INVALID NUMBER")
			}
		}
	}

	fmt.Println("sum: ", sum)
}

func max(a int, b int) int {
	if (a > b) {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if (a < b) {
		return a
	} else {
		return b
	}
}