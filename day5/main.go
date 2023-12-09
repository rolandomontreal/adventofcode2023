package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// seed-to-soil map:
// Destination range start, source range start, range length
// 50 98 2
// 52 50 48


type plantingMap struct {
	name string
	destinationRangeStart int
	sourceRangeStart int
	maprange int
}

func main() {
	fpath := "./actualdata.txt"
	bs, err := os.ReadFile(fpath)

	if err != nil {
		fmt.Println("Could not read input file: ", err)
		os.Exit(1)
	}

	data := string(bs)
	sections := strings.Split(data, "\n\n")
	seeds := parseSeeds(sections[0])
	results := [][]int{seeds}

	sections = sections[1:]
	// fmt.Println(sections)

	// Each section is a mapping
	for _, section := range sections {
		// Parse to proper format
		parsedSection := parseRanges(section)
		// mappedNums := []int{}
		// Use latest results from mapping from previous step
		latestResults := results[len(results)-1]
		newResults := []int{}
		for _, num := range latestResults {
			// fmt.Println("Mapping num: ", num)
			result := mapNum(parsedSection, num)
			// fmt.Println("Result: ", result, "\n")
			newResults = append(newResults, result)
		}
		results = append(results, newResults)
		fmt.Println("All results after iteration: ", results)
	}

	lastResult := results[len(results) - 1]
	fmt.Println("\n" ,lastResult)
	smallest := lastResult[0]
	for _, n := range lastResult {
		if n < smallest {
			smallest = n
		}
	}
	fmt.Println("Smallest num found: ", smallest)
}

func mapNum(pms []plantingMap, n int) int {
	output := n
	for i := 0; i < len(pms) && output == n; i++ {
		// fmt.Printf("Running map for %d for mapping %s\n", n, pms[i].name)
		lowerbound := pms[i].sourceRangeStart
		upperbound := pms[i].sourceRangeStart + pms[i].maprange
		// fmt.Printf("lowerbound: %d, upperbound: %d\n", lowerbound, upperbound)
		if (lowerbound <= n && n < upperbound) {
			diff := pms[i].sourceRangeStart - pms[i].destinationRangeStart
			output = n - diff
			// fmt.Printf("Found match for num %d, to output: %d\n", n, output)
		}
	}
	return output
}

func parseSeeds(seedsRow string) []int {
	output := []int{}

	numsStr := strings.Fields(strings.Split(seedsRow, ":")[1])
	for _, sn := range numsStr {
		n, err := strconv.Atoi(sn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output = append(output, n)
	}

	return output
}

func parseRanges(stringMap string) []plantingMap {
	output := []plantingMap{}
	name := strings.Split(strings.Split(stringMap, "\n")[0], " ")[0]

	rows := strings.Split(stringMap, "\n")[1:]
	for _, row := range rows {
		stringNums := strings.Fields(row)
		pm := plantingMap{
			name: name,
		}
		for i, sn := range stringNums {
			n, err := strconv.Atoi(sn)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if i == 0 {
				pm.destinationRangeStart = n
			} else if (i == 1) {
				pm.sourceRangeStart = n
			} else {
				pm.maprange = n
			}
		}
		output = append(output, pm)
	}

	return output
}