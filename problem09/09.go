package problem09

import (
	"sort"
	"strconv"
	"strings"

	"github.com/samtoffler/advent-of-code-2021/helpers"
)

var input []string = helpers.ParseFileAsStrings(helpers.GetInputPath(("problem09")))

func CountRiskLevels() int{
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
	
	
	riskLevel := 0
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if (isLowPoint(heights, j, i)) {
				riskLevel += heights[i][j] + 1
			}
		} 
	}
	return riskLevel
}

func isLowPoint(floorPlan [][]int, x int, y int) bool {
	width := len(floorPlan[0])
	height := len(floorPlan)

	// Corners
	if (x == 0 && y == 0) {
		return (floorPlan[y][x+1] > floorPlan[y][x]) && (floorPlan[y+1][x] > floorPlan[y][x])
	}

	if (x == width - 1 && y == 0) {
		return (floorPlan[y][x-1] > floorPlan[y][x]) && (floorPlan[y+1][x] > floorPlan[y][x])
	}

	if (x == 0 && y == height - 1) {
		return (floorPlan[y-1][x] > floorPlan[y][x]) && (floorPlan[y][x+1] > floorPlan[y][x])
	}

	if (x == width - 1 && y == height - 1) {
		return (floorPlan[y-1][x] > floorPlan[y][x]) && (floorPlan[y][x-1] > floorPlan[y][x])
	}

	// Edges
	if (x == 0) {
		return (floorPlan[y-1][x] > floorPlan[y][x]) && (floorPlan[y][x+1] > floorPlan[y][x]) && (floorPlan[y+1][x] > floorPlan[y][x])
	}

	if (x == width - 1) {
		return (floorPlan[y-1][x] > floorPlan[y][x]) && (floorPlan[y][x-1] > floorPlan[y][x]) && (floorPlan[y+1][x] > floorPlan[y][x])
	}

	if (y == 0) {
		return (floorPlan[y][x-1] > floorPlan[y][x]) && (floorPlan[y+1][x] > floorPlan[y][x]) && (floorPlan[y][x+1] > floorPlan[y][x])
	}

	if (y == height - 1) {
		return (floorPlan[y][x-1] > floorPlan[y][x]) && (floorPlan[y-1][x] > floorPlan[y][x]) && (floorPlan[y][x+1] > floorPlan[y][x])
	}

	// Centers
	return (floorPlan[y-1][x] > floorPlan[y][x]) && (floorPlan[y][x+1] > floorPlan[y][x]) && (floorPlan[y+1][x] > floorPlan[y][x]) && (floorPlan[y][x-1] > floorPlan[y][x])
}

func FindLargestBasins() int{
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
	
	var basinSizes []int 
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if (isLowPoint(heights, j, i)) {
				basinSizes = append(basinSizes, dfs(heights, i, j, len(heights[0]), len(heights)))
			}
		} 
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func dfs(grid [][]int, i int, j int, width int, height int) int {
	if (i < 0 || j < 0 || i > (height - 1) || j > (width - 1) || grid[i][j] == -1 || grid[i][j] == 9) {
		return 0
	}
	grid[i][j] = -1
	return 1 + dfs(grid, i+1, j, width, height) + dfs(grid, i-1, j, width, height) + dfs(grid, i, j+1, width, height) + dfs(grid, i, j-1, width, height)
}