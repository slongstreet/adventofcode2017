package main

import (
	"strings"
	"strconv"
	"fmt"
	"os"
	"log"
	"bufio"
)

type condition struct {
	register string
	operator string
	value int
}

type command struct {
	register string
	operator string
	value int
}

func main() {
	var registers = make(map[string]int)

	// perform test using sample data
	var testData = loadTestData()
	fmt.Printf("loaded %v lines of input...\n", len(testData))
	for _, line := range testData {
		com, con := parseInputLine(line)
		if isConditionTrue(registers, con) {
			processCommand(registers, com)
		}
	}
	fmt.Printf("max value = %v\n\n", calculateLargestRegisterValue(registers))

	// clear registers and use real input data
	registers = make(map[string]int)
	var maxRegisterValue = 0
	var inputData = loadInputFile("./input.txt")
	fmt.Printf("loaded %v lines of input...\n", len(inputData))
	for _, line := range inputData {
		com, con := parseInputLine(line)
		if isConditionTrue(registers, con) {
			processCommand(registers, com)

			var val = calculateLargestRegisterValue(registers)
			if val > maxRegisterValue {
				maxRegisterValue = val
			}
		}
	}
	fmt.Printf("max value = %v\n", calculateLargestRegisterValue(registers))
	fmt.Printf("highest value in registers = %v\n", maxRegisterValue)
}

func calculateLargestRegisterValue(registers map[string]int) int {
	var maxValue = 0
	for _, v := range registers {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

func getOrAddRegister(registers map[string]int, newRegister string) int {
	if _,ok := registers[newRegister]; !ok {
		registers[newRegister] = 0
	}

	return registers[newRegister]
}

func isConditionTrue(registers map[string]int, con condition) bool {
	var currValue = getOrAddRegister(registers, con.register)
	switch con.operator {
	case "<":
		return currValue < con.value
	case "<=":
		return currValue <= con.value
	case "==":
		return currValue == con.value
	case "!=":
		return currValue != con.value
	case ">":
		return currValue > con.value
	case ">=":
		return currValue >= con.value
	default:
		return false
	}
}

func processCommand(registers map[string]int, com command) {
	if com.operator == "inc" {
		registers[com.register] += com.value
	} else if com.operator == "dec" {
		registers[com.register] -= com.value
	}
}

func parseInputLine(input string) (command, condition) {
	var parts = strings.Split(input, " if ")

	// parse command from LHS
	var lhsParts = strings.Split(parts[0], " ")
	commandValue, _ := strconv.Atoi(lhsParts[2])

	// parse condition from RHS
	var rhsParts = strings.Split(parts[1], " ")
	conditionValue, _ := strconv.Atoi(rhsParts[2])

	return command { lhsParts[0], lhsParts[1], commandValue },
		condition { rhsParts[0], rhsParts[1], conditionValue }
}

func loadTestData() []string {
	return []string {
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}
}

func loadInputFile(filepath string) []string {
	var lines []string

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}