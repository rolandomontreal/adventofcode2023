package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	totalms int
	record int
}

func main() {
	fpath := "./actualdata.txt"
	bs, err := os.ReadFile(fpath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// races := parseRaces(bs)
	race := parseRaces2(bs)
	fmt.Println(race)
	
	nrways := 0
	for i := 0; i <= race.totalms; i++ {
		velocity := i
		timetravelling := race.totalms - i
		distance := velocity * timetravelling
		if distance > race.record {
			nrways++
		}
	}
	fmt.Println(nrways)
	

	// Part 1
	// nrwaystowin := []int{}
	// for _, race := range races {
	// 	nrways := 0
	// 	for i := 0; i <= race.totalms; i++ {
	// 		velocity := i
	// 		timetravelling := race.totalms - i
	// 		distance := velocity * timetravelling
	// 		if distance > race.record {
	// 			nrways++
	// 		}
	// 	}
	// 	nrwaystowin = append(nrwaystowin, nrways)
	// }

	// fmt.Println(races)
	// fmt.Println(nrwaystowin)

	// product := 1
	// for _, nrways := range nrwaystowin {
	// 	product *= nrways
	// }
	// fmt.Println(product)
}


// Part 2
func parseRaces2(bs []byte) race {
	s := string(bs)
	rows := strings.Split(s, "\n")	
	ts := strings.Join(strings.Fields(strings.Split(rows[0], ":")[1]), "")
	t, _ := strconv.Atoi(ts)
	ds := strings.Join(strings.Fields(strings.Split(rows[1], ":")[1]), "")
	d, _ := strconv.Atoi(ds)
	fmt.Println(ts)
	fmt.Println(ds)

	r := race{
		totalms: t,
		record: d,
	}

	return r
}

// Part 1
func parseRaces(bs []byte) []race {
	s := string(bs)
	rows := strings.Split(s, "\n")	
	ts := strings.Fields(strings.Split(rows[0], ":")[1])
	ds := strings.Fields(strings.Split(rows[1], ":")[1])

	output := []race{}
	for i, timestring := range ts {
		t, _ := strconv.Atoi(timestring)
		d, _ := strconv.Atoi(ds[i])
		r := race{
			totalms: t,
			record: d,
		}
		output = append(output, r)
	}

	return output
}