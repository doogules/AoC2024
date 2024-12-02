package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("day2/day2_input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	safe_count := 0
	// Section 2
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

		is_safe := true
		previous_level := -1
		trend := "unknown"
		for _, level := range line_split {
			level_int, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal(err)
			}
			if previous_level == -1 { // means it's the first item in the row
				previous_level = level_int
				continue
			} else if previous_level == level_int {
				is_safe = false
				break
			} else if level_int-previous_level > 3 || previous_level-level_int > 3 {
				is_safe = false
				break
			} else if trend == "increasing" && level_int < previous_level {
				is_safe = false
				break
			} else if trend == "decreasing" && level_int > previous_level {
				is_safe = false
				break
			} else if trend == "unknown" && level_int > previous_level {
				trend = "increasing"
			} else if trend == "unknown" && level_int < previous_level {
				trend = "decreasing"
			}
			previous_level = level_int
			continue
		}
		if is_safe {
			safe_count += 1
		}
		is_safe = true
		previous_level = -1
		trend = "unknown"
	}

	fmt.Printf("part 1 answer: %d", safe_count)
}
