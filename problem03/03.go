package problem03

import (
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem03")))

func GetPowerConsumption() int64 {
	zeroCounts, oneCounts := countOnesAnsZeros(input)

	// Create new binary numbers
	var gammaRateBinaryArray []string
	var epsilonRateBinaryArray []string
	for i := range zeroCounts {
		if zeroCounts[i] > oneCounts[i] {
			gammaRateBinaryArray = append(gammaRateBinaryArray, "0")
			epsilonRateBinaryArray = append(epsilonRateBinaryArray, "1")
		} else {
			gammaRateBinaryArray = append(gammaRateBinaryArray, "1")
			epsilonRateBinaryArray = append(epsilonRateBinaryArray, "0")
		}
	}

	gammaRateString := strings.Join(gammaRateBinaryArray, "")
	epsilonRateString := strings.Join(epsilonRateBinaryArray, "")

	gammaRate, err := strconv.ParseInt(gammaRateString, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilonRate, err := strconv.ParseInt(epsilonRateString, 2, 64)
	if err != nil {
		panic(err)
	}

	return gammaRate * epsilonRate
}

func countOnesAnsZeros(binaries []string) ([]int, []int) {
	// Initialize arrays that hold the counts of each bit
	binarySize := len(input[0])
	var zeroCounts []int
	var oneCounts []int
	for i := 0; i < binarySize; i++ {
		zeroCounts = append(zeroCounts, 0)
		oneCounts = append(oneCounts, 0)
	}

	// Count ones and zeros
	for _, elem := range binaries {
		for j, digit := range strings.Split(elem, "") {
			if digit == "0" {
				zeroCounts[j] += 1
			} else {
				oneCounts[j] += 1
			}
		}
	}

	return zeroCounts, oneCounts
}

func GetLifeSupportRating() int64 {
	return getOxygenGeneratorRating() * getCO2ScrubberRating()
}

func getOxygenGeneratorRating() int64 {
	return filterBinaries(input, 0, true)
}

func getCO2ScrubberRating() int64 {
	return filterBinaries(input, 0, false)
}

func filterBinaries(remainingBinaries []string, bitToConsider int, useMostCommonBit bool) int64 {
	if len(remainingBinaries) == 1 {
		rate, err := strconv.ParseInt(remainingBinaries[0], 2, 64)
		if err != nil {
			panic(err)
		}
		return rate
	}

	zeroCounts, oneCounts := countOnesAnsZeros(remainingBinaries)

	// Count most frequent bit
	var mostFrequent string
	if zeroCounts[bitToConsider] <= oneCounts[bitToConsider] {
		mostFrequent = "1"
	} else {
		mostFrequent = "0"
	}

	// Filter input
	var importantBit string
	var filteredBinaries []string
	for _, binary := range remainingBinaries {
		if (useMostCommonBit) {
			importantBit = mostFrequent
		} else {
			importantBit = switchBit(mostFrequent)

		}
		if binary[bitToConsider] == []byte(importantBit)[0] {
			filteredBinaries = append(filteredBinaries, binary)
		}
	}

	return filterBinaries(filteredBinaries, bitToConsider+1, useMostCommonBit)
}

func switchBit(bit string) string {
	if bit == "0" {
		return "1"
	} else {
		return "0"
	}
}
