package main

// 564267 - too high

// func main() {
// 	filepath := "./testdata.txt"
// 	bs, err := os.ReadFile(filepath)

// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	data := string(bs)

// 	// fmt.Println(data)
// 	// nums := pickNumsWithAdjacentSymbols(data)
// 	// fmt.Println(nums)
// 	// sum := 0
// 	// for _, num := range nums {
// 	// 	sum += num
// 	// }
// 	// fmt.Println(sum)
// }

// func pickNumsWithAdjacentSymbols(s string) []int {
// 	output := []int{}

// 	rows := strings.Split(s, "\n")
// 	for y, row := range rows {
// 		rownums := []int{}
// 		startOfNum := -1
// 		// fmt.Println(y)
// 		fmt.Println(row)

// 		for x := 0; x < len(row); x++ {
// 			c := row[x:x+1]
// 			// fmt.Println(c)
// 			n, err :=strconv.Atoi(c)
// 			// Found a number
// 			fmt.Println("num: ", n)
// 			if err == nil && startOfNum == -1  {
// 				startOfNum = x
// 			}
// 			// Found end of number by character
// 			endOfNumber := err != nil && startOfNum != -1
// 			// Found end of number by line end
// 			endOfLineAndNumber := x == len(row) - 1 && err == nil
// 			if endOfNumber || endOfLineAndNumber {
// 				var endIndex int
// 				if endOfNumber {
// 					endIndex = x
// 				} else {
// 					endIndex = x + 1
// 				}
// 				fmt.Println("error parsing num: ", err)
// 				// Run logic for checking adjacent coordinates
// 				fmt.Printf("startOfNum: %d, end: %d\n\n", startOfNum, endIndex)
// 				validNumber := isValidNumber(y, len(rows), startOfNum, endIndex, len(row), rows)
// 				n, err := strconv.Atoi(row[startOfNum:endIndex])
// 				fmt.Println("Number found: ", n)
// 				if err != nil {
// 					fmt.Println("Error parsing the number??", err)
// 					os.Exit(1)
// 				}
// 				if validNumber {
// 					output = append(output, n)
// 					rownums = append(rownums, n)
// 				}
// 				// Reset start of num
// 				startOfNum = -1
// 			}
// 		}
// 		fmt.Println(rownums)
// 	}

// 	return output
// }

// func isValidNumber(y int, totalRows int, startX int, endX int, rowLength int, rows []string) bool {
// 	foundSymbol := false
// 	for i := y - 1; i <= y + 1 && !foundSymbol; i++ {
// 		if (i >= 0 && i < totalRows) {
// 			for k := startX - 1; k <= endX + 1 && !foundSymbol; k++ {
// 				if (k >= 0 && k < rowLength) {
// 					c := rows[i][k:k+1]
// 					fmt.Printf("Checking %s  ", c)
// 					numericOrDot := strings.ContainsAny(c, "1234567890.")
// 					if !numericOrDot {
// 						foundSymbol = !numericOrDot
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return foundSymbol
// }

