package problem04

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

func FindFirstBingoWinner() int {
	bingoNumbers, boards := parseInput()

	for _, bingoNumber := range bingoNumbers {
		for _, board := range boards {
			updateBoard(&board, bingoNumber)
			if isWinningBoard(board) {
				bingoNumberInt, err := strconv.Atoi(bingoNumber)
				if err != nil {
					panic(err)
				}
				return calculateSum(board) * bingoNumberInt
			}
		}
	}

	return 0
}

func FindLastBingoWinner() int {
	bingoNumbers, boards := parseInput()

	winnersMap := make(map[int]bool)
	for i := range boards {
		winnersMap[i] = false
	}

	numWinners := 0
	for _, bingoNumber := range bingoNumbers {
		for boardNum, board := range boards {
			updateBoard(&board, bingoNumber)
			if isWinningBoard(board) && !winnersMap[boardNum] {
				numWinners += 1
				if numWinners == len(boards) {
					bingoNumberInt, err := strconv.Atoi(bingoNumber)
					if err != nil {
						panic(err)
					}
					return calculateSum(board) * bingoNumberInt
				}
				winnersMap[boardNum] = true
			}
		}
	}

	return 0
}

func calculateSum(board [][]string) int {
	sum := 0
	for row := range board {
		for column := 0; column < len(board[0]); column++ {
			number := board[row][column]
			if number != "Y" {
				numberNumber, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				sum += numberNumber
			}
		}
	}
	return sum
}

func isWinningBoard(board [][]string) bool {
	if isHorizontalWinner(board) || isVerticalWinner(board) {
		return true
	}
	return false
}

func updateBoard(board *[][]string, bingoNumber string) {
	for row := range *board {
		for column := 0; column < len((*board)[0]); column++ {
			if (*board)[row][column] == bingoNumber {
				(*board)[row][column] = "Y"
			}
		}
	}
}

func isWinningLine(line []string) bool {
	for _, number := range line {
		if number != "Y" {
			return false
		}
	}
	return true
}

func isHorizontalWinner(board [][]string) bool {
	for _, row := range board {
		if isWinningLine(row) {
			return true
		}
	}
	return false
}

func isVerticalWinner(board [][]string) bool {
	var line []string
	for column := 0; column < len(board[0]); column++ {
		for row := range board {
			line = append(line, board[row][column])
		}
		if isWinningLine(line) {
			return true
		}
		line = make([]string, 0)
	}
	return false
}

func parseInput() ([]string, [][][]string) {
	file, err := os.Open(helpers.GetInputPath(("problem04")))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	var bingoNumbers []string
	var boards [][][]string
	var board [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			bingoNumbers = append(bingoNumbers, strings.Split(line, ",")...)
		} else if lineNum > 1 {
			if line == "" {
				boards = append(boards, board)
				board = make([][]string, 0)
				continue
			}
			var row []string
			row = append(row, strings.Fields(line)...)
			board = append(board, row)
		}
		lineNum++
	}
	boards = append(boards, board)

	return bingoNumbers, boards
}
