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
	depthIncreases := 0

	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		check(err)
		str := ""
		if previousDepth < 0 {
			str = "(N/A - no previous measurement)"
		} else {
			if previousDepth < depth {
				str = "Increases"
				depthIncreases++
			} else {
				str = "Decreases"
			}
		}
		fmt.Printf("%d (%s)\n", depth, str)
		previousDepth = depth
	}
	fmt.Printf("Depth increased %d times", depthIncreases)
}
