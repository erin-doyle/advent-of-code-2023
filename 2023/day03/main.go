package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/mpvl/unique"

	"github.com/erin-doyle/advent-of-code-2023/util"
)

//go:embed input.txt
var input string

func isPartNumber(line string, start int, end int) bool {
	characterRegex := regexp.MustCompile(`[^0-9.]`)
	var isPartNumber bool = false

	var checkStart int
	var checkEnd int

	if start > 0 {
		checkStart = start - 1
	} else {
		checkStart = start
	}

	if end+1 < len(line) {
		checkEnd = end + 1
	} else {
		checkEnd = end
	}

	for index := checkStart; index <= checkEnd; index++ {
		if characterRegex.MatchString(string(line[index])) {
			isPartNumber = true
			break
		}
	}

	return isPartNumber
}

func part1(input string) int {
	var sumOfPartNumbers int = 0

	schematicLines := parseInput(input)

	// example:
	// 467..114..
	// ...*......
	// ..35..633.
	// ......#...
	// 617*......
	// .....+.58.
	// ..592.....
	// ......755.
	// ...$.*....
	// .664.598..
	for index := 0; index < len(schematicLines); index++ {
		var prevLine, curLine, nextLine string

		curLine = schematicLines[index]

		if index-1 > -1 {
			prevLine = schematicLines[index-1]
		}

		if index+1 < len(schematicLines) {
			nextLine = schematicLines[index+1]
		}

		numbersRegex := regexp.MustCompile(`\b*(\d+)\b*`)
		var numberIndices [][]int = numbersRegex.FindAllStringIndex(curLine, -1)

		for _, numberIndex := range numberIndices {
			var start int = numberIndex[0]
			var end int = numberIndex[1] - 1
			partNumber := util.ToInt(curLine[start : end+1])

			// check curLine
			if isPartNumber(curLine, start, end) {
				sumOfPartNumbers += partNumber
				continue
			}

			// check prevLine
			if prevLine != "" {
				if isPartNumber(prevLine, start, end) {
					sumOfPartNumbers += partNumber
					continue
				}
			}

			// check nextLine
			if nextLine != "" {
				if isPartNumber(nextLine, start, end) {
					sumOfPartNumbers += partNumber
					continue
				}
			}
		}
	}

	return sumOfPartNumbers
}

func findAdjacentNumbers(line string, location int) []int {
	numberRegex := regexp.MustCompile(`\d`)
	var foundNumbers []int

	var checkStart int
	var checkEnd int

	if location > 0 {
		checkStart = location - 1
	} else {
		checkStart = location
	}

	if location+1 < len(line) {
		checkEnd = location + 1
	} else {
		checkEnd = location
	}

	for index := checkStart; index <= checkEnd; index++ {
		if numberRegex.MatchString(string(line[index])) {
			numStart := index
			numEnd := index

			for numStart > 0 && numberRegex.MatchString(string(line[numStart-1])) {
				numStart--
			}

			for numEnd < len(line)-1 && numberRegex.MatchString(string(line[numEnd+1])) {
				numEnd++
			}

			foundNumbers = append(foundNumbers, util.ToInt(line[numStart:numEnd+1]))
		}
	}

	// dedupe foundNumbers
	unique.Ints(&foundNumbers)

	return foundNumbers
}

func part2(input string) int {
	var sumOfGearRatios int = 0

	schematicLines := parseInput(input)

	// example:
	// 467..114..
	// ...*......
	// ..35..633.
	// ......#...
	// 617*......
	// .....+.58.
	// ..592.....
	// ......755.
	// ...$.*....
	// .664.598..
	for index := 0; index < len(schematicLines); index++ {
		curLine := schematicLines[index]

		gearRegex := regexp.MustCompile(`\*`)
		var gearIndices [][]int = gearRegex.FindAllStringIndex(curLine, -1)

		if len(gearIndices) == 0 {
			continue
		}

		var prevLine, nextLine string

		if index-1 > -1 {
			prevLine = schematicLines[index-1]
		}

		if index+1 < len(schematicLines) {
			nextLine = schematicLines[index+1]
		}

		for _, gearIndex := range gearIndices {
			var location int = gearIndex[0]
			var adjacentNumbers []int

			// check curLine
			adjacentNumbers = append(adjacentNumbers, findAdjacentNumbers(curLine, location)...)

			// check prevLine
			if prevLine != "" {
				adjacentNumbers = append(adjacentNumbers, findAdjacentNumbers(prevLine, location)...)
			}

			// check nextLine
			if nextLine != "" {
				adjacentNumbers = append(adjacentNumbers, findAdjacentNumbers(nextLine, location)...)
			}

			// A gear is any * symbol that is adjacent to exactly two part numbers.
			if len(adjacentNumbers) != 2 {
				continue
			}

			// Its gear ratio is the result of multiplying those two numbers together.
			sumOfGearRatios += adjacentNumbers[0] * adjacentNumbers[1]
		}
	}

	return sumOfGearRatios
}

func parseInput(input string) (ans []string) {
	return strings.Split(input, "\n")
}

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
