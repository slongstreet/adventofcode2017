package main

import (
	"math"
	"fmt"
)

type Cell struct {
	Value int
	X int
	Y int
}

func main() {
	var testCases = []int{1, 12, 23, 1024}
	var resultCases = []int{0, 3, 2, 31}

	for i, val := range testCases {
		grid := populateGrid(val)
		dist := calculateDistance(grid[val-1])

		var result string
		if dist == resultCases[i] {
			result = "PASS"
		} else {
			result = "FAIL"
		}

		fmt.Printf("test: %v => %v (%s)\n", val, dist, result)
	}

	fmt.Println()

	var input = 289326
	grid := populateGrid(input)
	dist := calculateDistance(grid[input-1])
	fmt.Printf("dist: %v => %v\n", input, dist)
}

func populateGrid(maxValue int) map[int]Cell {
	var cellMap = make(map[int]Cell, maxValue)
	var maxPosX = 0
	var maxNegX = 0
	var maxPosY = 0
	var maxNegY = 0
	var currentDirection = 0  // 0 = +x, 1 = +y, 2 = -x, 3 = -y

	cellMap[0] = Cell{ 1, 0, 0 }
	var prevCell = cellMap[0]

	for i := 1; i < maxValue; i++ {
		switch currentDirection {
		case 0: // +x
			cellMap[i] = Cell{i + 1, prevCell.X + 1, prevCell.Y}
			if cellMap[i].X > maxPosX {
				maxPosX = cellMap[i].X
				currentDirection = 1 // +y
			}
		case 1: // +y
			cellMap[i] = Cell{i + 1, prevCell.X, prevCell.Y + 1}
			if cellMap[i].Y > maxPosY {
				maxPosY = cellMap[i].Y
				currentDirection = 2 // -x
			}
		case 2: // -x
			cellMap[i] = Cell{i + 1, prevCell.X - 1, prevCell.Y}
			if cellMap[i].X < maxNegX {
				maxNegX = cellMap[i].X
				currentDirection = 3 // -y
			}
		case 3: // -y
			cellMap[i] = Cell{i + 1, prevCell.X, prevCell.Y - 1}
			if cellMap[i].Y < maxNegY {
				maxNegY = cellMap[i].Y
				currentDirection = 0 // +x
			}
		}

		prevCell = cellMap[i]
	}

	return cellMap
}

func calculateDistance(cell Cell) int {
	return int(math.Abs(float64(cell.X)) + math.Abs(float64(cell.Y)))
}