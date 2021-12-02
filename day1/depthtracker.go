package main

import (
	"fmt"
	"os"
)

func check(e error) {
	fmt.Print("fail")
	if e!= nil {
		panic(e)
	}
}

func main() {
	podata, err := os.ReadFile("input.txt")
	check(err)
	fmt.Print(string(podata))
}
