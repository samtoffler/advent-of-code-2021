package problem02

import (
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem02")))

func CalculatePosition(useAim bool) int {
	xPos, yPos, aim := 0, 0, 0
	for _, command := range input {
		splitCommand := strings.Split(command, " ")
		direction := splitCommand[0]
		distance, err := strconv.Atoi(splitCommand[1])
		if err != nil {
			panic(err)
		}
		switch direction {
		case "forward":
			if useAim {
				yPos += aim * distance
			}
			xPos += distance
		case "up":
			if useAim {
				aim -= distance
			} else {
				yPos -= distance
			}
		case "down":
			if useAim {
				aim += distance
			} else {
				yPos += distance
			}

		}
	}
	return xPos * yPos
}
