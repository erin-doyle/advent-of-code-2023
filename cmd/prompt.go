package cmd

import "github.com/erin-doyle/advent-of-code-2023/aoc"

func main() {
	day, year, cookie := aoc.ParseFlags()
	aoc.GetPrompt(day, year, cookie)
}
