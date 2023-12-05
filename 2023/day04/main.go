package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/erin-doyle/advent-of-code-2023/util"
)

//go:embed input.txt
var input string

type card struct {
	id             string
	winningNumbers []string
	myNumbers      []string
}

type winningCard struct {
	card                card
	countWinningNumbers int
	countCopies         int
}

var pointsCache = map[int]int{}
var winnersCache = map[string]winningCard{}

func compareNumStrings(a, b string) int {
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)

	if aInt == bInt {
		return 0
	}
	if aInt < bInt {
		return -1
	}
	if aInt > bInt {
		return 1
	}

	return -1
}

func getCountMyWinningNumbers(card card) int {
	var count int = 0

	slices.SortFunc(card.winningNumbers, compareNumStrings)
	slices.SortFunc(card.myNumbers, compareNumStrings)

	for _, myNumber := range card.myNumbers {

		if myNumber == "" {
			continue
		}

		for _, winningNumber := range card.winningNumbers {
			if winningNumber == "" {
				continue
			}

			comparison := compareNumStrings(myNumber, winningNumber)

			// match found
			if comparison == 0 {
				count++
				break
			}

			// no match in winning numbers, exit
			if comparison == -1 {
				break
			}

			// match not yet found, keep looking
			if comparison == 1 {
				continue
			}
		}
	}

	return count
}

func getPoints(countWinningNumbers int) int {
	// count the points
	// The first match makes the card worth one point
	// each match after the first doubles the point value of that card.
	if _, ok := pointsCache[countWinningNumbers]; !ok {
		var points int = 1

		for idx := 1; idx < countWinningNumbers; idx++ {
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
		countMyWinningNumbers := getCountMyWinningNumbers(card)

		if countMyWinningNumbers == 0 {
			continue
		}

		points := getPoints(countMyWinningNumbers)

		sumPoints += points
	}

	return sumPoints
}

func getCountCardCopies(cardId int, allCards *[]card) int {
	winningCard, ok := winnersCache[strconv.Itoa(cardId)]

	if !ok {
		card := (*allCards)[cardId-1]
		winningCard.card = card
		winningCard.countWinningNumbers = getCountMyWinningNumbers(card)

		if winningCard.countWinningNumbers > 0 {
			var index int = 1
			var countCopies int = 0

			// exit if we've reached the table
			for index <= winningCard.countWinningNumbers && (index+cardId) <= len(*allCards) {
				countCopies += getCountCardCopies((index+cardId), allCards) + 1
				index++
			}

			winningCard.countCopies = countCopies

		} else {
			winningCard.countCopies = 0
		}

		winnersCache[winningCard.card.id] = winningCard
	}

	return winningCard.countCopies
}

func part2(input string) int {
	var sumTotalCards int = 0

	cards := parseInput(input)

	for cardNo := len(cards); cardNo > 0; cardNo-- {
		sumTotalCards += getCountCardCopies(cardNo, &cards) + 1
	}

	return sumTotalCards
}

func parseInput(input string) (ans []card) {
	for _, line := range strings.Split(input, "\n") {
		var nextCard card

		// example:
		// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		numbersRegex := regexp.MustCompile(`Card\s+(?P<id>\d+): (?P<winning>[\d\s]+) \| (?P<mine>[\d\s]+)`)
		cardNumbers := numbersRegex.FindStringSubmatch(line)

		for i, name := range numbersRegex.SubexpNames() {
			if i == 0 {
				continue
			}

			if name == "id" {
				nextCard.id = cardNumbers[i]
				continue
			}

			var numbers []string = strings.Split(cardNumbers[i], " ")

			if name == "winning" {
				nextCard.winningNumbers = numbers
			} else if name == "mine" {
				nextCard.myNumbers = numbers
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
