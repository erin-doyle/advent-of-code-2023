package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/erin-doyle/advent-of-code-2023/util"
)

//go:embed input.txt
var input string

type destinationMapping struct {
	destination int
	mapRange    int
}

type almanac struct {
	seeds    []string
	ordering map[string][]int
	maps     map[string]map[string]destinationMapping
}

func part1(input string) int {
	var lowestLocation int = 0

	almanac := parseInput(input)

	fmt.Println(almanac) // TODO: remove

	// a list of maps which describe how to convert numbers from a source category into numbers in a destination category
	// Each line within a map contains three numbers:
	// the destination range start, the source range start, and the range length.

	// What is the lowest location number that corresponds to any of the initial seed numbers?

	return lowestLocation
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (ans almanac) {
	seedsRegex := regexp.MustCompile(`seeds: ([\d\s]+)`)

	var section string
	for index, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if index == 0 {
			seedsMatch := seedsRegex.FindStringSubmatch(line)
			ans.seeds = strings.Split(seedsMatch[1], " ")
			ans.maps = map[string]map[string]destinationMapping{}
			ans.ordering = map[string][]int{}
			continue
		}

		if line == "" {
			if section != "" {
				sectionMap := ans.maps[section]
				sectionKeys := make([]int, 0, len(sectionMap))

				for k := range sectionMap {
					sectionKeys = append(sectionKeys, util.ToInt(k))
				}

				slices.Sort(sectionKeys)

				ans.ordering[section] = sectionKeys
			}

			section = ""
			continue
		}

		if section == "" {
			switch line {
			case "seed-to-soil map:":
				section = "seedToSoil"
			case "soil-to-fertilizer map:":
				section = "soilToFertilizer"
			case "fertilizer-to-water map:":
				section = "fertilizerToWater"
			case "water-to-light map:":
				section = "waterToLight"
			case "light-to-temperature map:":
				section = "lightToTemp"
			case "temperature-to-humidity map:":
				section = "tempToHumidity"
			case "humidity-to-location map:":
				section = "humidityToLocation"
			default:
				panic("map type not found!")
			}

			continue
		}

		if section != "" && line != "" {
			mapping := strings.Split(line, " ")

			if _, ok := ans.maps[section]; !ok {
				ans.maps[section] = map[string]destinationMapping{}
			}

			ans.maps[section][mapping[1]] = destinationMapping{
				destination: util.ToInt(mapping[0]),
				mapRange:    util.ToInt(mapping[1]),
			}

			continue
		}

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
