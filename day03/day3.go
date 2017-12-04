package main

import (
	"fmt"
	"math"
	"strconv"
)

type Cell struct {
	Id int
	Value int
	X int
	Y int
}

func main() {
	var testCases = []int{1, 12, 23, 1024}
	var resultCases = []int{0, 3, 2, 31}

	for i, val := range testCases {
		grid := populateGrid(val, false)
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
	grid := populateGrid(input, false)
	dist := calculateDistance(grid[input-1])
	fmt.Printf("dist: %v => %v\n", input, dist)
	fmt.Println()

	// for part 2, populate the grid again using the "sumGrid" algorithm
	grid = populateGrid(50000, true)
	for i := 0; i < 50000; i++ {
		if grid[i].Value > input {
			fmt.Printf("answer2: %v\n", grid[i].Value)
			break
		}
	}
}

func populateGrid(maxValue int, sumGrid bool) map[int]Cell {
	var cellMap = make(map[int]Cell, maxValue)
	var maxPosX = 0
	var maxNegX = 0
	var maxPosY = 0
	var maxNegY = 0
	var currentDirection = 0  // 0 = +x, 1 = +y, 2 = -x, 3 = -y

	cellMap[0] = Cell{1, 1, 0, 0 }
	var prevCell = cellMap[0]

	var ptrMap = make(map[string]Cell, maxValue)
	ptrMap["0,0"] = cellMap[0]

	for i := 1; i < maxValue; i++ {
		var id = i + 1
		var newX, newY int

		switch currentDirection {
		case 0: // +x
			newX = prevCell.X + 1
			newY = prevCell.Y
			if newX > maxPosX {
				maxPosX = newX
				currentDirection = 1 // +y
			}
		case 1: // +y
			newX = prevCell.X
			newY = prevCell.Y + 1
			if newY > maxPosY {
				maxPosY = newY
				currentDirection = 2 // -x
			}
		case 2: // -x
			newX = prevCell.X - 1
			newY = prevCell.Y
			if newX < maxNegX {
				maxNegX = newX
				currentDirection = 3 // -y
			}
		case 3: // -y
			newX = prevCell.X
			newY = prevCell.Y - 1
			if newY < maxNegY {
				maxNegY = newY
				currentDirection = 0 // +x
			}
		}

		if sumGrid {
			// calculate the sum for the new cell using its neighbors in the pointer map
			var sum = 0
			if cell, ok := ptrMap[strconv.Itoa(newX-1) + "," + strconv.Itoa(newY)]; ok {
				sum += cell.Value
			}
			if cell, ok := ptrMap[strconv.Itoa(newX-1) + "," + strconv.Itoa(newY+1)]; ok {
				sum += cell.Value
			}
			if cell, ok := ptrMap[strconv.Itoa(newX) + "," + strconv.Itoa(newY+1)]; ok {
				sum += cell.Value
			}
			if cell, ok := ptrMap[strconv.Itoa(newX+1) + "," + strconv.Itoa(newY+1)]; ok {
				sum += cell.Value
			}
			if cell, ok := ptrMap[strconv.Itoa(newX+1) + "," + strconv.Itoa(newY)]; ok {
				sum += cell.Value
			}
			if cell, ok := ptrMap[strconv.Itoa(newX+1) + "," + strconv.Itoa(newY-1)]; ok {
				sum += cell.Value
			}
			if cell, ok := ptrMap[strconv.Itoa(newX) + "," + strconv.Itoa(newY-1)]; ok {
				sum += cell.Value
			}
			if cell, ok := ptrMap[strconv.Itoa(newX-1) + "," + strconv.Itoa(newY-1)]; ok {
				sum += cell.Value
			}

			// create the new Cell, setting its value to the sum of its neighbors
			cellMap[i] = Cell{id, sum, newX, newY}

			// add the new cell to the pointer map
			var ptrKey = strconv.Itoa(newX) + "," + strconv.Itoa(newY)
			ptrMap[ptrKey] = cellMap[i]
		} else {
			// create the new Cell using i + 1 as its value
			cellMap[i] = Cell{id, i + 1, newX, newY}
		}

		prevCell = cellMap[i]
	}

	return cellMap
}

func calculateDistance(cell Cell) int {
	return int(math.Abs(float64(cell.X)) + math.Abs(float64(cell.Y)))
}