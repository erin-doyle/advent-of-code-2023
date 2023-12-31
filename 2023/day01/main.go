package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/erin-doyle/advent-of-code-2023/util"
)

//go:embed input.txt
var input string

var numVals = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
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

func part1(input string) int {
	calibrationDocument := parseInput(input)

	var sumCalibrationValues int
	sumCalibrationValues = 0

	r, _ := regexp.Compile(`\b*(\d){1}\b*`)

	for _, calibrationString := range calibrationDocument {
		var calibrationValues []string
		calibrationValues = r.FindAllString(calibrationString, -1)

		if len(calibrationValues) == 0 {
			continue
		}

		var calibrationValueStr string
		calibrationValueStr = calibrationValues[0]

		if len(calibrationValues) > 1 {
			calibrationValueStr += calibrationValues[len(calibrationValues)-1]
		} else {
			calibrationValueStr += calibrationValues[0]
		}

		calibrationValue, _ := strconv.Atoi(calibrationValueStr)
		sumCalibrationValues += calibrationValue
	}

	return sumCalibrationValues
}

func part2(input string) int {

	calibrationDocument := parseInput(input)

	var sumCalibrationValues int
	sumCalibrationValues = 0

	numSearchStrings := strings.Join(util.Keys(numVals), "|")
	firstMatch, _ := regexp.Compile(fmt.Sprintf(`(%s|\d).*`, numSearchStrings))
	lastMatch, _ := regexp.Compile(fmt.Sprintf(`.*(%s|\d)`, numSearchStrings))

	for _, calibrationString := range calibrationDocument {
		var valueOne, valueTwo string
		valueOne = getNumMatch(calibrationString, firstMatch)
		valueTwo = getNumMatch(calibrationString, lastMatch)

		calibrationValue, err := strconv.Atoi((valueOne + valueTwo))

		if err != nil {
			log.Fatal(err)
		}

		sumCalibrationValues += calibrationValue
	}

	return sumCalibrationValues
}

func getNumMatch(calibrationString string, regex *regexp.Regexp) string {
	var calibrationValues [][]string
	calibrationValues = regex.FindAllStringSubmatch(calibrationString, -1)

	if len(calibrationValues) == 0 {
		log.Fatal("No calibration values found!")
	}

	var calibrationValueStr string
	calibrationValueStr = calibrationValues[0][1]

	if util.CanCastAtoi(calibrationValueStr) {
		return calibrationValueStr
	} else {
		return numVals[calibrationValueStr]
	}
}

func parseInput(input string) (ans []string) {
	ans = strings.Split(input, "\n")
	return ans
}
