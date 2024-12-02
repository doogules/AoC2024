package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("day1/day1_input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	leftColumn := []int{}
	rightColumn := []int{}

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
		line_split := strings.Split(line_string, "   ")
		left_string := line_split[0]
		right_string := line_split[1]
		left_int, err := strconv.Atoi(left_string)
		if err == nil {
			leftColumn = append(leftColumn, left_int)
		}
		right_int, err := strconv.Atoi(right_string)
		if err == nil {
			rightColumn = append(rightColumn, right_int)
		}
	}
	slices.Sort(leftColumn)
	slices.Sort(rightColumn)
	total_distance := 0
	for i, l := range leftColumn {
		r := rightColumn[i]
		if l == r {
			continue
		} else if l > r {
			total_distance += (l - r)
		} else if r > l {
			total_distance += (r - l)
		}
	}
	fmt.Printf("part 1 answer: %d", total_distance)

	similarity_score := 0
	for _, l := range leftColumn {
		multiplier := 0
		for _, r := range rightColumn {
			if r < l {
				continue
			} else if r == l {
				multiplier += 1
				continue
			} else if r > l {
				break
			}
		}
		similarity_score += (l * multiplier)
	}
	fmt.Printf("part 2 answer: %d", similarity_score)
}
