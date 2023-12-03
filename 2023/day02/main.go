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

type play struct {
	blue  int
	green int
	red   int
}

type game struct {
	id    int
	plays []map[string]int
}

var colors = []string{"blue", "green", "red"}

func part1(input string) int {
	parsed := parseInput(input)
	_ = parsed

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans []game) {
	// example line:
	// Game 1: 7 green, 14 red, 5 blue; 8 red, 4 green; 6 green, 18 red, 9 blue
	for _, line := range strings.Split(input, "\n") {
		var nextGame game

		gameIdRegex := regexp.MustCompile(`Game (\d+):`)
		nextGame.id = util.ToInt(gameIdRegex.FindStringSubmatch(line)[1])

		playsRegex := regexp.MustCompile(`Game \d+: ([a-z1-9,;\s]+)`)
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
