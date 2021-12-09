package problem09

import (
	"sort"
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem09")))

func CountRiskLevels() int {
	heights := getHeights()
	riskLevel := 0
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if isLowPoint(heights, j, i) {
				riskLevel += heights[i][j] + 1
			}
		}
	}
	return riskLevel
}

func FindLargestBasins(useBfs bool) int {
	heights := getHeights()
	var basinSizes []int
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if isLowPoint(heights, j, i) {
				if useBfs {
					basinSizes = append(basinSizes, bfs(heights, i, j, len(heights[0]), len(heights)))
				} else {
					basinSizes = append(basinSizes, dfs(heights, i, j, len(heights[0]), len(heights)))
				}
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func getHeights() [][]int {
	var heights [][]int
	for _, line := range input {
		var intLine []int
		for _, char := range strings.Split(line, "") {
			num, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			intLine = append(intLine, num)
		}
		heights = append(heights, intLine)
	}
	return heights
}

func isLowPoint(floorPlan [][]int, x int, y int) bool {
	for _, neighbor := range [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		if isVaildSquare(floorPlan, y+neighbor[1], x+neighbor[0]) {
			if floorPlan[y+neighbor[1]][x+neighbor[0]] <= floorPlan[y][x] {
				return false
			}
		}
	}
	return true
}

func isVaildSquare(grid [][]int, row int, column int) bool {
	width := len(grid[0])
	height := len(grid)
	return row >= 0 && row < height && column >= 0 && column < width
}

func dfs(grid [][]int, i int, j int, width int, height int) int {
	if !isVaildSquare(grid, i, j) || grid[i][j] == -1 || grid[i][j] == 9 {
		return 0
	}
	grid[i][j] = -1
	return 1 + dfs(grid, i+1, j, width, height) + dfs(grid, i-1, j, width, height) + dfs(grid, i, j+1, width, height) + dfs(grid, i, j-1, width, height)
}

func bfs(grid [][]int, i int, j int, width int, height int) int {
	var queue [][]int
	queue = append(queue, []int{j, i})
	count := 0
	for ok := true; ok; ok = len(queue) != 0 {
		// Pop off the queue
		row := queue[0][1]
		column := queue[0][0]
		queue = queue[1:]

		if grid[row][column] != -1 && grid[row][column] != 9 {
			// Add eligible neighbors to queue
			for _, neighbor := range [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				if isVaildSquare(grid, row+neighbor[1], column+neighbor[0]) {
					queue = append(queue, []int{column + neighbor[0], row + neighbor[1]})
				}
			}
			// Mark current spot as visited
			grid[row][column] = -1
			count++
		}
	}
	return count
}
