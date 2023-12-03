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

type game struct {
	id    int
	plays []map[string]int
}

var colors = []string{"blue", "green", "red"}
var limits = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func part1(input string) int {
	var sumOfPossibleGames int = 0
	var totalPossibleCubes int = limits["blue"] + limits["green"] + limits["red"]

	games := parseInput(input)

	for _, nextGame := range games {
		var isGamePossible bool = true

		for _, play := range nextGame.plays {
			var totalPlayCubes int = 0

			for _, cubeColor := range colors {
				cubeCount, ok := play[cubeColor]

				if ok {
					if cubeCount > limits[cubeColor] {
						isGamePossible = false
					}

					totalPlayCubes += cubeCount
				}

			}

			if isGamePossible && totalPlayCubes > totalPossibleCubes {
				isGamePossible = false
			}
		}

		if isGamePossible {
			sumOfPossibleGames += nextGame.id
		}

	}

	return sumOfPossibleGames
}

func part2(input string) int {
	var sumOfPowers int = 0

	games := parseInput(input)

	for _, nextGame := range games {
		minimums := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}

		for _, play := range nextGame.plays {

			for _, cubeColor := range colors {
				cubeCount, ok := play[cubeColor]

				if ok && cubeCount > minimums[cubeColor] {
					minimums[cubeColor] = cubeCount
				}

			}
		}

		// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together
		sumOfPowers += minimums["blue"] * minimums["green"] * minimums["red"]
	}

	return sumOfPowers
}

func parseInput(input string) (ans []game) {
	// example line:
	// Game 1: 7 green, 14 red, 5 blue; 8 red, 4 green; 6 green, 18 red, 9 blue
	for _, line := range strings.Split(input, "\n") {
		var nextGame game

		gameIdRegex := regexp.MustCompile(`Game (\d+):`)
		nextGame.id = util.ToInt(gameIdRegex.FindStringSubmatch(line)[1])

		playsRegex := regexp.MustCompile(`Game \d+: ([a-z0-9,;\s]+)`)
		playsString := strings.Split(playsRegex.FindStringSubmatch(line)[1], ";")

		for _, playString := range playsString {
			play := map[string]int{}

			for _, color := range colors {
				colorRegex := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, color))
				colorMatch := colorRegex.FindStringSubmatch(playString)

				if len(colorMatch) > 0 {
					play[color] = util.ToInt(colorMatch[1])
				}
			}

			nextGame.plays = append(nextGame.plays, play)
		}

		ans = append(ans, nextGame)
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
