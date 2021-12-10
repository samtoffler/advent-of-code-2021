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
	"github.com/samtoffler/advent-of-code-2021/problem07"
	"github.com/samtoffler/advent-of-code-2021/problem08"
	"github.com/samtoffler/advent-of-code-2021/problem09"
	"github.com/samtoffler/advent-of-code-2021/problem10"
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
	case "7a":
		fmt.Println(problem07.GetFuelCost(true))
	case "7b":
		fmt.Println(problem07.GetFuelCost(false))
	case "8a":
		fmt.Println(problem08.CountUniqueNumbers())
	case "8b":
		fmt.Println(problem08.DecodePatterns())
	case "9a":
		fmt.Println(problem09.CountRiskLevels())
	case "9b":
		fmt.Println(problem09.FindLargestBasins(true))
	case "9c":
		fmt.Println(problem09.FindLargestBasins(false))
	case "10a":
		fmt.Println(problem10.ScoreIllegalCharacters(true))
	case "10b":
		fmt.Println(problem10.ScoreIllegalCharacters(false))
	default:
		fmt.Println("Unsolved problem")
	}
}