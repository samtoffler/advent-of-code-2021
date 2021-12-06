package main

import (
	"flag"
	"fmt"

	"github.com/samtoffler/advent-of-code-2021/problem01"
	"github.com/samtoffler/advent-of-code-2021/problem02"
	"github.com/samtoffler/advent-of-code-2021/problem03"
	"github.com/samtoffler/advent-of-code-2021/problem04"
	"github.com/samtoffler/advent-of-code-2021/problem05"
	"github.com/samtoffler/advent-of-code-2021/problem06"
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
	case "3a":
		fmt.Println(problem03.GetPowerConsumption())
	case "3b":
		fmt.Println(problem03.GetLifeSupportRating())
	case "4a":
		fmt.Println(problem04.FindFirstBingoWinner())
	case "4b":
		fmt.Println(problem04.FindLastBingoWinner())
	case "5a":
		fmt.Println(problem05.CountIntersections(false))
	case "5b":
		fmt.Println(problem05.CountIntersections(true))
	case "6a":
		fmt.Println(problem06.CountLanternfish(80))
	case "6b":
		fmt.Println(problem06.CountLanternfish(256))
	default:
		fmt.Println("Unsolved problem")
	}
}