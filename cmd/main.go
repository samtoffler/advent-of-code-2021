package main

import (
	"flag"
	"fmt"
	"github/samtoffler/advent-of-code-2021/problem01"
	"github/samtoffler/advent-of-code-2021/problem02"
)

var (problemNum = flag.String("problem", "1a", "The problem to run"))

func main() {
	flag.Parse()
	switch (*problemNum) {
	case "1a":
		fmt.Println(problem01.CountIncreases())
	case "1b":
		fmt.Println(problem01.CountIncreasesWithAim())
	case "2a":
		fmt.Println(problem02.CalculatePosition(false))
	case "2b":
		fmt.Println(problem02.CalculatePosition(true))
	default:
		fmt.Println("Unsolved problem")
	}
}