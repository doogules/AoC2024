package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {

	file, _ := os.Open("day3/day3_input.txt")
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	sum_part2 := 0

	pattern := "mul(X,Y)"
	enabled := true
	doPattern := "do()"
	dontPattern := "don't()"
	Xdigits := make([]string, 3)
	Ydigits := make([]string, 3)
	patternIndex := 0
	doPatternIndex := 0

	for _, j := range data {
		jrune := rune(j)
		jstring := string(j)
		if !enabled {
			if jrune == 'd' {
				doPatternIndex = 1
				continue
			} else if j == doPattern[doPatternIndex] {
				if j == ')' {
					enabled = true
					doPatternIndex = 0
					println("do()")
					continue
				} else {
					doPatternIndex += 1
					continue
				}
			} else {
				doPatternIndex = 0
				continue
			}
		} else if enabled {
			if jrune == 'd' {
				doPatternIndex = 1
			} else if j == dontPattern[doPatternIndex] {
				if j == ')' {
					enabled = false
					doPatternIndex = 0
					println("don't()")
					continue
				} else {
					doPatternIndex += 1
				}
			} else {
				doPatternIndex = 0
			}
		}
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
				sum_part2 += product
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
	fmt.Printf("part 2 answer: %d\n", sum_part2)
	// 2369451, 32939762
}
