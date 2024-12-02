package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("day2/day2_input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	safe_count_part1 := 0
	safe_count_part2 := 0
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			fmt.Printf("ReadLine: %q\n", line)
		}
		if err != nil {
			break
		}

		line_string := string(line)
		line_split := strings.Split(line_string, " ")
		line_int := []int{}
		for _, level := range line_split {
			level_int, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal(err)
			}
			line_int = append(line_int, level_int)
		}

		if isSafe(line_int) {
			safe_count_part1 += 1
			safe_count_part2 += 1
			continue
		} else {
			for i, _ := range line_int {
				report_without_level := slices.Clone(line_int)
				report_without_level = slices.Delete(report_without_level, i, i+1)
				if isSafe(report_without_level) {
					safe_count_part2 += 1
					break
				}
			}
		}
	}
	fmt.Printf("part 1 answer: %d\n", safe_count_part1)
	fmt.Printf("part 2 answer: %d", safe_count_part2)
}

func isSafe(levels []int) bool {
	is_safe := true
	previous_level := -1
	trend := "unknown"

	for _, level := range levels {

		if previous_level == -1 { // means it's the first item in the row
			previous_level = level
			continue
		} else if previous_level == level {
			is_safe = false
			break
		} else if level-previous_level > 3 || previous_level-level > 3 {
			is_safe = false
			break
		} else if trend == "increasing" && level < previous_level {
			is_safe = false
			break
		} else if trend == "decreasing" && level > previous_level {
			is_safe = false
			break
		} else if trend == "unknown" && level > previous_level {
			trend = "increasing"
		} else if trend == "unknown" && level < previous_level {
			trend = "decreasing"
		}
		previous_level = level
		continue
	}
	return is_safe
}
