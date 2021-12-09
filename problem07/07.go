package problem07

import (
	"math"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var positions []int = helpers.ParseFileAsArray(helpers.GetInputPath(("problem07")))

func GetFuelCost(constantCost bool) int {
	maxPosition := helpers.FindMax(positions)
	minPosition := helpers.FindMin(positions)

	var minFuelCost int
	if constantCost {
		minFuelCost = maxPosition * len(positions)
	} else {
		minFuelCost = maxPosition * (maxPosition + 1) / 2 * len(positions)
	}
	
	for i := minPosition; i < maxPosition; i++ {
		runningFuelCost := 0
		for _, position := range(positions) {
			distance := int(math.Abs(float64(position-i)))
			if (constantCost) {
				runningFuelCost += distance
			} else {
				runningFuelCost += (distance * (distance + 1) / 2)
			}
		}
		if (runningFuelCost < minFuelCost) {
			minFuelCost = runningFuelCost
		}
	}
	return minFuelCost
}