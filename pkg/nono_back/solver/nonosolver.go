package solver

import "fmt"

func isSolvable(tips [2][][]int) {
	grid := make([][]int, len(tips[0]))
	for i := range grid {
		grid[i] = make([]int, len(tips[1]))
	}
	grid = mathPrefill(tips[0], len(tips[1]), grid)
	grid = rotateGrid(mathPrefill(tips[1], len(tips[0]), rotateGrid(grid)))
	nonoToJPG(grid, "test_solver.jpg")
}

func rotateGrid(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid[0]))
	for i := range newGrid {
		newGrid[i] = make([]int, len(grid))
		for j := range grid {
			newGrid[i][j] = grid[j][i]
		}
	}
	return newGrid
}

func sum(nbs []int) int {
	sum := 0
	for i := range nbs {
		sum += nbs[i]
	}
	return sum
}

func mathPrefill(tips [][]int, size int, grid [][]int) [][]int {
	for xTips := range tips {
		lineSum := sum(tips[xTips]) + len(tips[xTips]) - 1
		currLineSum := 0
		for tip := range tips[xTips] {
			if currLineSum > 0 {
				currLineSum++
			}
			if tips[xTips][tip] > size-lineSum {
				for toFill := tips[xTips][tip] - (size - lineSum); toFill > 0; toFill-- {
					fmt.Println(tips[xTips][tip], size, lineSum, currLineSum, toFill)
					grid[xTips][currLineSum+toFill+size-lineSum-1] = 1
				}
			}
			currLineSum += tips[xTips][tip]
		}
	}
	return grid
}
