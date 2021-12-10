package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInputPath(folder string) string {
	return fmt.Sprintf("/Users/samtoffler/code/personal/advent-of-code-2021/%s/input.txt", folder)
}

func ParseFileAsStrings(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var res []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}

	return res
}

func ParseFileAsInts(path string) []int {
	fileAsStrings := ParseFileAsStrings(path)
	
	var res []int
	for _, elem := range fileAsStrings {
		number, err := strconv.Atoi(elem)
		if (err != nil) {
			panic(err)
		}
		res = append(res, number)
	}

	return res
}

func ParseFileAsArray(path string) []int {
	fileAsStrings := ParseFileAsStrings(path)
	var res []int
	for _, elem := range strings.Split(fileAsStrings[0], ",") {
		number, err := strconv.Atoi(elem)
		if (err != nil) {
			panic(err)
		}
		res = append(res, number)
	}
	return res
}

func FindMax(input []int) int{
	max := input[0]
	for _, num := range input {
		if num > max {
			max = num
		}
	}
	return max
}

func FindMin(input []int) int {
	min := input[0]
	for _, num := range input {
		if num < min {
			min = num
		}
	}
	return min
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}