package problem05

import (
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem05")))


func CountIntersections(countDiagonals bool) int{
	board := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		board[i] = make([]int, 1000)
		for j := 0; j < 1000; j++ {
			board[i][j] = 0
		}
	}

	for _, line := range input {
		x1, x2, y1, y2 := getCoordinates(line)

		if (x1 == x2) {
			if (y1 <= y2) {
				for y := y1; y <= y2; y++ {
					board[y][x1] = board[y][x1] + 1
				}
			} else {
				for y := y1; y >= y2; y-- {
					board[y][x1] = board[y][x1] + 1
				}
			}
		} else if (y1 == y2) {
			if (x1 <= x2) {
				for x := x1; x <= x2; x++ {
					board[y1][x] = board[y1][x] + 1
				}
			} else {
				for x := x1; x >= x2; x-- {
					board[y1][x] = board[y1][x] + 1
				}
			}
		} else if (countDiagonals) {
			if (x1 <= x2 && y1 <= y2) { // Right and down
				for i := 0; i <= (x2-x1); i++ {
					board[y1+i][x1+i] = board[y1+i][x1+i] + 1
				}
			} else if (x1 <= x2 && y1 > y2) { // Right and up
				for i := 0; i <= (x2-x1); i++ {
					board[y1-i][x1+i] = board[y1-i][x1+i] + 1
				}
			} else if (x1 > x2 && y1 <= y2) { // Left and down
				for i := 0; i <= (x1-x2); i++ {
					board[y2-i][x2+i] = board[y2-i][x2+i] + 1
				}
			} else if (x1 > x2 && y1 > y2) { // Left and up
				for i := 0; i <= (x1-x2); i++ {
					board[y2+i][x2+i] = board[y2+i][x2+i] + 1
				}
			}
		}
	}
	return countSpots(board)
}

func getCoordinates(line string) (int, int, int, int) {
	pairs := strings.Split(line, " -> ")
		point1, point2 := pairs[0], pairs[1]
		x1Str, y1Str := strings.Split(point1, ",")[0], strings.Split(point1, ",")[1]
		x2Str, y2Str := strings.Split(point2, ",")[0], strings.Split(point2, ",")[1]
		x1, err := strconv.Atoi(x1Str)
		if err != nil {
			panic(err)
		}
		x2, err := strconv.Atoi(x2Str)
		if err != nil {
			panic(err)
		}
		y1, err := strconv.Atoi(y1Str)
		if err != nil {
			panic(err)
		}
		y2, err := strconv.Atoi(y2Str)
		if err != nil {
			panic(err)
		}
		return x1, x2, y1, y2
}

func countSpots(board [][]int) int {
	sum := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if board[i][j] > 1 {
				sum++
			}
		}
	}
	return sum
}