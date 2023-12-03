package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() []string {
	file, _ := os.Open("day2.txt")
	scanner := bufio.NewScanner(file)

	strArr := []string{}

	for scanner.Scan() {
		strArr = append(strArr, scanner.Text())
	}

	return strArr
}

func parseGame() {
	games := readFile()
	sum := 0
	sum2 := 0

	for _, g := range games {
		strId := strings.Split(g, " ")[1]
		strId = strId[0 : len(strId)-1]
		id, _ := strconv.Atoi(strId)

		// Part One
		if gamePossible(g[8:]) {
			sum += id
		}

		// Part two
		sum2 += minSetCubes(g)
	}

	fmt.Printf("Part1: %d\n", sum)
	fmt.Printf("Part2: %d\n", sum2)
}

func minSetCubes(game string) int {
	semiIndex := strings.Index(game, ":")
	game = game[semiIndex+2:]
	games := strings.Split(game, "; ")
	minBlue := 0
	minRed := 0
	minGreen := 0

	for _, g := range games {
		colors := strings.Split(g, ", ")

		for _, c := range colors {
			strArr := strings.Split(c, " ")
			num, _ := strconv.Atoi(strArr[0])

			color := strArr[1]

			switch color {
			case "blue":
				if num > minBlue {
					minBlue = num
				}
			case "green":
				if num > minGreen {
					minGreen = num
				}
			case "red":
				if num > minRed {
					minRed = num
				}
			}
		}
	}

	return minBlue * minRed * minGreen
}

func gamePossible(game string) bool {
	semiIndex := strings.Index(game, ":")
	game = game[semiIndex+2:]
	games := strings.Split(game, "; ")

	for _, g := range games {
		colors := strings.Split(g, ", ")

		for _, c := range colors {
			strArr := strings.Split(c, " ")
			num, _ := strconv.Atoi(strArr[0])

			color := strArr[1]

			if color == "blue" && num > 14 {
				return false
			} else if color == "green" && num > 13 {
				return false
			} else if color == "red" && num > 12 {
				return false
			}
		}
	}

	return true
}
