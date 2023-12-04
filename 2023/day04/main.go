package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/erin-doyle/advent-of-code-2023/util"
)

//go:embed input.txt
var input string

type card struct {
	winningNumbers string
	myNumbers      string
}

var pointsCache = map[int]int{}

func getCountMyWinningNumbers(card card) int {
	var count int = 0

	// find the winning numbers (numbers on the right that are in the left)
	// TODO

	return count
}

func getPoints(countWinningNumbers int) int {
	// count the points
	// The first match makes the card worth one point
	// each match after the first doubles the point value of that card.
	if _, ok := pointsCache[countWinningNumbers]; !ok {
		var points int = 1

		// TODO: this can be further optimized by starting with the nearest number in the cache if one exists

		for idx := 1; idx <= countWinningNumbers; idx++ {
			points = points * 2
		}

		pointsCache[countWinningNumbers] = points
	}

	return pointsCache[countWinningNumbers]
}

func part1(input string) int {
	var sumPoints int = 0

	cards := parseInput(input)

	for _, card := range cards {
		fmt.Println(card.winningNumbers) // TODO: remove
		fmt.Println(card.myNumbers)      // TODO: remove
		countMyWinningNumbers := getCountMyWinningNumbers(card)

		if countMyWinningNumbers == 0 {
			continue
		}

		points := getPoints(countMyWinningNumbers)

		sumPoints += points
	}

	return sumPoints
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []card) {
	for _, line := range strings.Split(input, "\n") {
		var nextCard card

		// example:
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		numbersRegex := regexp.MustCompile(`Card \d+: (?P<winning>[\d\s]+) \| (?P<mine>[\d\s]+)`)
		numbers := numbersRegex.FindStringSubmatch(line)

		for i, name := range numbersRegex.SubexpNames() {
			if name == "winning" {
				nextCard.winningNumbers = numbers[i]
			} else if name == "mine" {
				nextCard.myNumbers = numbers[i]
			}
		}

		ans = append(ans, nextCard)
	}
	return ans
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
