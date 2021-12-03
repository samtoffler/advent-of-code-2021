package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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