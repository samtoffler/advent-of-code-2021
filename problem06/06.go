package problem06

import (
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem06")))

func CountLanternfish(numDays int) int {
	fishCounts := make(map[int]int)
	for i := 0; i < 8; i++ {
		fishCounts[i] = 0
	}

	for _, f := range strings.Split(input[0], ",") {
		num, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		fishCounts[num] = fishCounts[num] + 1
	}

	for i := 0; i < numDays; i++ {
		newFishCounts := make(map[int]int)
		for j := 0; j < 9; j++ {
			if j == 0 {
				newFishCounts[8] = fishCounts[j]
			} else if j == 7 {
				newFishCounts[6] = fishCounts[7] + fishCounts[0]
			} else {
				newFishCounts[j-1] = fishCounts[j]
			}
		}
		fishCounts = newFishCounts
	}

	fishCount := 0
	for i := range fishCounts {
		fishCount += fishCounts[i]
	}
	return fishCount
}
