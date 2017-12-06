package main

import (
	"bytes"
	"strconv"
	"fmt"
)

func main() {
	var testCase = []int { 0, 2, 7, 0 }
	var uniquenessMap = make(map[string]bool)
	var rebalances = 0

	for {
		rebalances++
		state := rebalance(testCase)
		if _, ok := uniquenessMap[state]; ok {
			break // if we've seen this state before, we're done
		}

		uniquenessMap[state] = true // insert state into uniqueness map
	}

	fmt.Printf("test case => %v\n", rebalances)

	var input = []int { 4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3 }
	var state string
	uniquenessMap = make(map[string]bool)
	rebalances = 0

	for {
		rebalances++
		state = rebalance(input)
		if _, ok := uniquenessMap[state]; ok {
			break // if we've seen this state before, we're done
		}

		uniquenessMap[state] = true // insert state into uniqueness map
	}

	fmt.Printf("part 1 => %v\n", rebalances)

	// to calculate the size of the loop (part 2), keep the bank state, but reset the uniqueness map and rebalance count
	uniquenessMap = make(map[string]bool)
	uniquenessMap[state] = true
	rebalances = 0

	for {
		rebalances++
		state := rebalance(input)
		if _, ok := uniquenessMap[state]; ok {
			break // if we've seen this state before, we're done
		}

		uniquenessMap[state] = true // insert state into uniqueness map
	}

	fmt.Printf("part 2 => %v\n", rebalances)
}

// rebalance banks and return a string representation to enable uniqueness checking
func rebalance(banks []int) string {
	// find the bank with the most blocks
	var topBank = 0
	var maxBlocks = 0
	for i := 0; i < len(banks); i++ {
		if banks[i] > maxBlocks {
			maxBlocks = banks[i]
			topBank = i
		}
	}

	// rebalance
	var currentBankIndex = topBank
	banks[topBank] = 0
	for blocks := maxBlocks; blocks > 0; blocks-- {
		currentBankIndex = (currentBankIndex + 1) % len(banks)
		banks[currentBankIndex]++
	}

	// build string representation of bank states
	var buffer bytes.Buffer
	for i := 0; i < len(banks); i++ {
		buffer.WriteString(strconv.Itoa(banks[i]))
		buffer.WriteRune('.')
	}

	return buffer.String()
}
