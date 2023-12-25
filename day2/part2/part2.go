package main

import (
	"fmt"
	"os"
	"strings"
)

const InputFilePath = "input"

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

func main() {
	dat, err := os.ReadFile(InputFilePath)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	count := 0

	for _, game := range lines {
		gameParts := strings.Split(game, ":")
		if len(gameParts) != 2 {
			continue
		}
		cubesInfo := gameParts[1]

		sets := strings.Split(cubesInfo, ";")

		var maxCubes = [3]int{0, 0, 0}
		for _, set := range sets {
			rgb := UnwrapSet(set)

			for i := range rgb {
				maxCubes[i] = max(maxCubes[i], rgb[i])
			}
		}

		cubesProduct := 1
		for _, cubes := range maxCubes {
			cubesProduct *= cubes
		}
		count += cubesProduct
	}
	print(count)
}
