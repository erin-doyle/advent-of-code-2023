package cmd

import (
	"flag"
	"time"

	"github.com/erin-doyle/advent-of-code-2023/util/skeleton"
)

func main() {
	today := time.Now()
	day := flag.Int("day", today.Day(), "day number to fetch, 1-25")
	year := flag.Int("year", today.Year(), "AOC year")
	flag.Parse()
	skeleton.Run(*day, *year)
}
