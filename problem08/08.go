package problem08

import (
	"sort"
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem08")))

func CountUniqueNumbers() int {
	count := 0
	for _, line := range input {
		outputValuesString := strings.Split(line, " | ")[1]
		outputValues := strings.Fields(outputValuesString)
		for _, outputValue := range outputValues {
			if len(outputValue) == 2 || len(outputValue) == 4 || len(outputValue) == 3 || len(outputValue) == 7 {
				count++
			}
		}
	}

	return count
}

func DecodePatterns() int {
	total := 0
	for _, line := range input {
		splitLine := strings.Split(line, " | ")
		patternString := splitLine[0]
		outputValuesString := splitLine[1]
		patternValues := strings.Fields(patternString)
		outputValues := strings.Fields(outputValuesString)
		mappings := make(map[int]string)
		for i, patternValue := range patternValues {
			patternValues[i] = sortString(patternValue)
			orderedString := sortString(patternValue)
			if len(orderedString) == 2 {
				mappings[1] = orderedString
			} else if len(orderedString) == 3 {
				mappings[7] = orderedString
			} else if len(orderedString) == 4 {
				mappings[4] = orderedString
			} else if len(orderedString) == 7 {
				mappings[8] = orderedString
			}
		}

		// Determine 5 - 5 is the only digit that has 5 segments and has the top left and middle segments
		var topLeftAndMiddle []string
		for _, s := range strings.Split(mappings[4], "") {
			if strings.Contains(mappings[4], s) && !strings.Contains(mappings[1], s) {
				topLeftAndMiddle = append(topLeftAndMiddle, s)
			}
		}
		for _, patternValue := range patternValues {
			if len(patternValue) == 5 {
				if strings.Contains(patternValue, topLeftAndMiddle[0]) && strings.Contains(patternValue, topLeftAndMiddle[1]) {
					mappings[5] = sortString(patternValue)
				}
			}
		}

		// Determine 9 - 9 can be created by combining 4 and 5
		var nine []string
		nine = append(nine, strings.Split(mappings[4], "")...)
		for _, letter := range strings.Split(mappings[5], "") {
			if !contains(nine, letter) {
				nine = append(nine, letter)
			}
		}
		mappings[9] = sortString(strings.Join(nine, ""))

		// Determine 6 - 6 is the only digit that's 6 segments and has all the same segments as 5, and isn't 9
		for _, patternValue := range patternValues {
			split5 := strings.Split(mappings[5], "")
			if mappings[9] != patternValue && len(patternValue) == 6 && strings.Contains(patternValue, split5[0]) && strings.Contains(patternValue, split5[1]) && strings.Contains(patternValue, split5[2]) && strings.Contains(patternValue, split5[3]) && strings.Contains(patternValue, split5[4]) {
				mappings[6] = sortString(patternValue)
			}
		}

		// Determine 0 - 0 is the only digit that's six segments if 6 and 9 have already been mapped
		for _, patternValue := range patternValues {
			if len(patternValue) == 6 && mappings[6] != patternValue && mappings[9] != patternValue {
				mappings[0] = sortString(patternValue)
			}
		}

		// Determine 3 - 3 is 5 digits long and has the same segments ar 9
		for _, patternValue := range patternValues {
			splitNumber := strings.Split(patternValue, "")
			if len(patternValue) == 5 && mappings[8] != patternValue && mappings[5] != patternValue {
				if strings.Contains(mappings[9], splitNumber[0]) && strings.Contains(mappings[9], splitNumber[1]) && strings.Contains(mappings[9], splitNumber[2]) && strings.Contains(mappings[9], splitNumber[3]) && strings.Contains(mappings[9], splitNumber[4]) {
					mappings[3] = sortString(patternValue)
				}
			}
		}

		// Determine 2 - 2 is the only remaining digit
		var mapValues []string
		for _, v := range mappings {
			mapValues = append(mapValues, v)
		}
		for _, patternValue := range patternValues {
			if !contains(mapValues, patternValue) {
				mappings[2] = sortString(patternValue)
			}
		}

		output := ""
		for _, outputValue := range outputValues {
			for k, v := range mappings {
				if v == sortString(outputValue) {
					output = output + strconv.Itoa(k)
				}
			}
		}

		num, err := strconv.Atoi(output)
		if err != nil {
			panic(err)
		}

		total += num
	}

	return total
}

func sortString(input string) string {
	s := strings.Split(input, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
