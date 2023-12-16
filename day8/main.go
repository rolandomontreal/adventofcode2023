package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type node struct {
	left string
	right string
}

func main() {
	bs, err := os.ReadFile("./actualdata.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := string(bs)
	splitted := strings.Split(data, "\n\n")
	instructions, rawnodes := splitted[0], splitted[1]
	// fmt.Println(repeatedInstr)

	fmt.Println(instructions)
	// fmt.Println("Raw nodes: ", rawnodes)

	nm, currentNode := parseNodes(rawnodes)
	fmt.Println(nm)
	fmt.Println(currentNode)
	
	zzzfound := false
	maxIterations := 100
	k := 0
	linstr := "L"
	finalvalue := "ZZZ"
	start := time.Now()
	for !zzzfound && k < maxIterations {
		if k % 100000 == 0 {
			fmt.Println(float64(k) / float64(maxIterations))
		}
		for i := 0; i < len(instructions) && !zzzfound; i++ {
			var newkey string
			if instructions[i:i+1] == linstr {
				newkey = currentNode.left
			} else {
				newkey = currentNode.right
			}
			if newkey == finalvalue {
				zzzfound = true
				steps := k * len(instructions) + i + 1
				fmt.Println("\n\nFound zzz after", steps, "steps :)")
			} else {
				currentNode = nm[newkey]
			}
		}
		k++
	}
	duration := time.Since(start)
	fmt.Println("duration: ", duration)
}

func parseNodes(nodesString string) (map[string]node, node) {
	nodemap := map[string]node{}
	var initialNode node

	rows := strings.Split(nodesString, "\n")

	for _, row := range rows {
		split := strings.Split(row, " = ")
		key, values := split[0], split[1]
		v1, v2 := values[1:4], values[6:9]
		nodemap[key] = node{
			left: v1,
			right: v2,
		}
		if key == "AAA" {
			initialNode = node{
				left: v1,
				right: v2,
			}
		}
	}

	return nodemap, initialNode
}