package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func check(e error) {
	if e!= nil {
		panic(e)
	}
}

func main() {
	podatafile, err := os.Open("input.txt")
	check(err)
	defer podatafile.Close()

	scanner := bufio.NewScanner(podatafile)

	previousDepth := -1
	previousDepth2 := -1
	previousDepth3 := -1
	previousSum := -1
	depthIncreases := 0

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)
		previousDepth3 = previousDepth2
		previousDepth2 = previousDepth
		previousDepth = depth
		if previousDepth3 > 0 {
			sum := previousDepth + previousDepth2 + previousDepth3
			str := ""
			if previousSum < 0 {
				str = "N/A - no previous sum"
			} else if previousSum < sum {
				str = "Increases"
				depthIncreases++
			} else {
				str = "Decreases"
			}
			fmt.Printf("%d (%s)\n", sum, str)
			previousSum = sum
		}
	}
	fmt.Printf("Depth increased %d times", depthIncreases)
}
