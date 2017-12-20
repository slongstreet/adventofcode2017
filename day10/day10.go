package main

import "fmt"

func main() {
	var testSlice = makeIncrementingSlice(5)
	knotHash(testSlice, []int{3, 4, 1, 5})
	var checksum = testSlice[0] * testSlice[1]
	fmt.Printf("Test checksum = %v\n\n", checksum)

	var slice = makeIncrementingSlice(256)
	knotHash(slice, []int{130, 126, 1, 11, 140, 2, 255, 207, 18, 254, 246, 164, 29, 104, 0, 224}) // lengths from day10 input
	checksum = slice[0] * slice[1]
	fmt.Printf("checksum = %v\n\n", checksum)
}

func knotHash(list []int, lengths []int) {
	var currentPos, skipSize int
	for _, length := range lengths {
		swapSegment(list, currentPos, length)
		currentPos = safeIndexAdvance(len(list), currentPos, length + skipSize)
		skipSize++
	}
}

func swapSegment(list []int, currentPos int, length int) {
	if length == 1 {
		return
	}

	var listLength = len(list)
	var target = currentPos + length - 1
	var swap int

	for currentPos < target {
		swap = list[currentPos % listLength]
		list[currentPos % listLength] = list[target % listLength]
		list[target % listLength] = swap

		currentPos++
		target--
	}
}

func safeIndexAdvance(sliceLength int, currentPos int, steps int) int {
	return (currentPos + steps) % sliceLength
}

func makeIncrementingSlice(length int) []int {
	var slice = make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = i
	}

	return slice
}