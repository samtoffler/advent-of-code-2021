package problem10

import (
	"sort"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem10")))


func ScoreIllegalCharacters(getCorruptedScore bool) int {
	pairs := make(map[string]string)
	pairs[")"] = "("
	pairs["]"] = "["
	pairs["}"] = "{"
	pairs[">"] = "<"

	
	corruptedScore := 0
	var incompleteScores []int
	for _, line := range input {
		var stack []string
		isCorrupted := false

		for _, char := range strings.Split(line, "") {
			if helpers.Contains([]string{"(", "[", "{", "<"}, char) {
				stack = append(stack, char)
			} else {
				if (pairs[char] == stack[len(stack)-1]) {
					stack = stack[:len(stack)-1]
				} else {
					switch char {
					case ")":
						corruptedScore += 3
					case "]":
						corruptedScore += 57
					case "}":
						corruptedScore += 1197
					case ">":
						corruptedScore += 25137
					}
					isCorrupted = true
					break
				}
			}
		}

		if (!isCorrupted) {
			incompleteScore := 0
			for i := len(stack) - 1; i >= 0; i-- {
				switch stack[i] {
				case "(":
					incompleteScore = incompleteScore * 5 + 1
				case "[":
					incompleteScore = incompleteScore * 5 + 2
				case "{":
					incompleteScore = incompleteScore * 5 + 3
				case "<":
					incompleteScore = incompleteScore * 5 + 4
				}
			}
			incompleteScores = append(incompleteScores, incompleteScore)
		}

	}
	if (getCorruptedScore) {
		return corruptedScore
	} else {
		sort.Ints(incompleteScores)
		return incompleteScores[len(incompleteScores)/2]
	}
	
	
}