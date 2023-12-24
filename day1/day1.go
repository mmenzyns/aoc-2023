package main

import (
	"maps"
	"os"
	"strconv"
	"strings"
)

const PartMode = 2 // Switch the logic between part 1 and part 2
const InputFilePath = "input"

var Numbers = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"0": 0,
}

var NumbersWords = map[string]int{ // For part 2
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"zero":  0,
}

func findIndex(line string, reverse bool) (finalIndex int, finalValue int) {
	/*
		Find the lowest finalIndex of a number in a "line" string.
		"reverse" will reverse the search, so it will find the highest finalIndex.)
	*/
	var indexFunc func(string, string) int
	var isBetterCandidate func(int, int) bool
	if reverse {
		indexFunc = strings.LastIndex
		isBetterCandidate = func(i1 int, i2 int) bool { return i1 > i2 }
	} else {
		indexFunc = strings.Index
		isBetterCandidate = func(i1 int, i2 int) bool { return i1 < i2 }
	}

	finalIndex = -1 // to initialize first value
	for key, value := range Numbers {
		index := indexFunc(line, key)
		if index == -1 {
			continue
		}
		if finalIndex == -1 { // first iteration, initialize
			finalIndex, finalValue = index, value
			continue
		}
		if isBetterCandidate(index, finalIndex) {
			finalIndex, finalValue = index, value
		}
	}
	return
}

func main() {
	dat, err := os.ReadFile(InputFilePath)
	if err != nil {
		panic(err)
	}

	if PartMode == 2 {
		maps.Copy(Numbers, NumbersWords)
	}

	lines := strings.Split(string(dat), "\n")

	count := 0
	for _, line := range lines {
		i1, v1 := findIndex(line, false)
		i2, v2 := findIndex(line, true)
		if i1 == -1 || i2 == -1 {
			continue
		}

		d, err := strconv.Atoi(strconv.Itoa(v1) + strconv.Itoa(v2))
		if err != nil {
			panic(err)
		}
		count += d
	}
	print(count)
}
