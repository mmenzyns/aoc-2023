package main

import (
	"fmt"
	"os"
	"strings"
)

const InputFilePath = "input"

var gameRequirements = []int{12, 13, 14} // Red: 12, Green: 13, Blue: 14
var cubeTypes = map[string]int{
	"red":   0,
	"green": 1,
	"blue":  2,
}

func UnwrapSet(set string) (rgb [3]int) {
	var count int
	var color string

	for _, cubes := range strings.Split(set, ",") {
		n, err := fmt.Sscanf(cubes, " %d %s", &count, &color)
		if err != nil || n != 2 {
			panic(err)
		}
		index := cubeTypes[color]
		rgb[index] = count
	}
	return rgb
}

func followsRequirements(rgb [3]int) bool {
	for colorIndex := range gameRequirements {
		if rgb[colorIndex] > gameRequirements[colorIndex] {
			return false
		}
	}
	return true
}

func main() {
	dat, err := os.ReadFile(InputFilePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	viableGameIds := 0
LineLoop:
	for _, line := range lines {
		gameParts := strings.Split(line, ":")
		if len(gameParts) != 2 {
			continue
		}
		gameInfo := gameParts[0]
		cubeInfo := gameParts[1]

		sets := strings.Split(cubeInfo, ";")

		var gameId int
		n, err := fmt.Sscanf(gameInfo, "Game %d", &gameId)
		if err != nil || n != 1 {
			panic("Couldn't parse game info")
		}
		for _, set := range sets {
			rgb := UnwrapSet(set)
			if !followsRequirements(rgb) {
				continue LineLoop
			}
		}
		viableGameIds += gameId
	}
	print(viableGameIds)
}
