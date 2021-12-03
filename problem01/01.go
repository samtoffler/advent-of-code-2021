package problem01

import (
	"github/samtoffler/advent-of-code-2021/helpers"
)

var input []int = helpers.ParseFileAsInts(helpers.GetInputPath(("problem01")))

func CountIncreases() int {
	return countIncreases(input)
}

func CountIncreasesWithAim() int {
	var depthSums []int
	for i := range input[:len(input)-2] {
		depthSums = append(depthSums, input[i]+input[i+1]+input[i+2])
	}
	return countIncreases(depthSums)
}

func countIncreases(depths []int) int {
	numIncreases := 0
	curr := depths[0]
	for _, depth := range depths[1:] {
		if depth > curr {
			numIncreases += 1
		}
		curr = depth
	}
	return numIncreases
}
