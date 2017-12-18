package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
)

func main() {
	var testcases = loadTestCases()
	for k, v := range testcases {
		var result string

		var score, garbage = calculateScoreAndGarbage(k)
		if score == v {
			result = "PASS"
		} else {
			result = "FAIL"
		}

		fmt.Printf("%s : %v,%v (%s)\n", k, score, garbage, result)
	}

	fmt.Println(

	)
	var score, garbage = calculateScoreAndGarbage(loadInputFile("./input.txt"))
	fmt.Printf("score = %v, garbage = %v\n", score, garbage)
}

func calculateScoreAndGarbage(input string) (score int, garbage int) {
	var depth int
	var ignoreNext, inGarbage bool
	for _, r := range input {
		if ignoreNext {
			ignoreNext = false
			continue
		}

		if inGarbage {
			switch r {
			case '>':
				inGarbage = false
			case '!':
				ignoreNext = true
			default:
				garbage++
			}
			continue
		}

		switch r {
		case '{':
			depth++
		case '}':
			score += depth
			depth--
		case '<':
			inGarbage = true
		case '!':
			ignoreNext = true
		}
	}

	return score, garbage
}

func loadInputFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return line
}

func loadTestCases() map[string]int {
	return map[string]int {
		"{}": 1,
		"{{{}}}": 6,
		"{{},{}}": 5,
		"{{{},{},{{}}}}": 16,
		"{<a>,<a>,<a>,<a>}": 1,
		"{{<ab>},{<ab>},{<ab>},{<ab>}}": 9,
		"{{<!!>},{<!!>},{<!!>},{<!!>}}": 9,
		"{{<a!>},{<a!>},{<a!>},{<ab>}}": 3,
	}
}