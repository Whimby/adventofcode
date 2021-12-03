package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
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
	distance := 0
	depth := 0

	for scanner.Scan() {
		command := scanner.Text()
		reg := regexp.MustCompile(`(\w+) (\d+)`)
		regArray := reg.FindStringSubmatch(command)
		direction := regArray[1]
		value, err := strconv.Atoi(regArray[2])
		check(err)

		switch direction {
		case "forward":
			distance += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}
	fmt.Printf("%d moved forward, %d moved down\n", distance, depth)
	fmt.Printf("multiplied together you get %d", distance * depth)
}
