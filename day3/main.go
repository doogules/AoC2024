package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {

	file, _ := os.Open("day3/day3_input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	sum_part1 := 0
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		pattern := "mul(X,Y)"
		Xdigits := make([]string, 3)
		Ydigits := make([]string, 3)
		patternIndex := 0

		for _, j := range line {
			jrune := rune(j)
			jstring := string(j)
			if jrune == 'm' {
				Xdigits = make([]string, 3)
				Ydigits = make([]string, 3)
				patternIndex = 1
				continue
			}
			if patternIndex > 0 {
				if j == pattern[patternIndex] {
					patternIndex += 1
					continue
				} else if jrune == ',' && Xdigits[0] != "" {
					patternIndex = 6
					continue
				} else if Xdigits[0] != "" && Ydigits[0] != "" && !unicode.IsDigit(jrune) && jrune != ')' {
					Xdigits = make([]string, 3)
					Ydigits = make([]string, 3)
					patternIndex = 0
				} else if jrune == ')' && Xdigits[0] != "" && Ydigits[0] != "" {
					Xint, _ := strconv.Atoi(Xdigits[0] + Xdigits[1] + Xdigits[2])
					Yint, _ := strconv.Atoi(Ydigits[0] + Ydigits[1] + Ydigits[2])
					product := Xint * Yint
					sum_part1 += product
					println(fmt.Sprintf("mul(%d,%d)", Xint, Yint))
					Xdigits = make([]string, 3)
					Ydigits = make([]string, 3)
					patternIndex = 0
				} else if j != pattern[patternIndex] && !unicode.IsDigit(jrune) {
					Xdigits = make([]string, 3)
					Ydigits = make([]string, 3)
					patternIndex = 0
					continue
				} else if (patternIndex == 4 || patternIndex == 6) && !unicode.IsDigit(jrune) {
					Xdigits = make([]string, 3)
					Ydigits = make([]string, 3)
					patternIndex = 0
					continue
				} else if unicode.IsDigit(jrune) {
					if patternIndex == 4 {
						if Xdigits[0] != "" && Xdigits[1] != "" && Xdigits[2] != "" { // X longer than 3 digits
							Xdigits = make([]string, 3)
							Ydigits = make([]string, 3)
							patternIndex = 0
							continue
						} else if Xdigits[0] != "" && Xdigits[1] != "" {
							Xdigits[2] = jstring
							continue
						} else if Xdigits[0] != "" {
							Xdigits[1] = jstring
							continue
						} else {
							Xdigits[0] = jstring
							continue
						}
					} else if patternIndex == 6 {
						if Ydigits[0] != "" && Ydigits[1] != "" && Ydigits[2] != "" { // Y longer than 3 digits
							Xdigits = make([]string, 3)
							Ydigits = make([]string, 3)
							patternIndex = 0
							continue
						} else if Ydigits[0] != "" && Ydigits[1] != "" {
							Ydigits[2] = jstring
							continue
						} else if Ydigits[0] != "" {
							Ydigits[1] = jstring
							continue
						} else {
							Ydigits[0] = jstring
							continue
						}
					}
				}
			}
		}
	}
	fmt.Printf("part 1 answer: %d\n", sum_part1)
	// 2369451, 32939762
}
