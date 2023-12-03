package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func readFile(fileName string) []string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)

	strArr := []string{}

	for scanner.Scan() {
		strArr = append(strArr, scanner.Text())
	}

	return strArr
}

func findFirstAndLast() []string {
	text := readFile("day1.txt")
	digitStrings := []string{}
	for _, str := range text {
		var digitString string
		for i := 0; i < len(str); i++ {
			char := rune(str[i])
			if unicode.IsDigit(char) {
				digitString += string(char)
				break
			}
		}
		for i := len(str) - 1; i >= 0; i-- {
			char := rune(str[i])
			if unicode.IsDigit(char) {
				digitString += string(char)
				break
			}
		}

		digitStrings = append(digitStrings, digitString)
	}

	return digitStrings
}

func sum(digits []string) int {
	sum := 0
	for _, str := range digits {
		num, _ := strconv.Atoi(str)
		sum += num
	}

	return sum
}

func findNumbers() {
	var digits [][]int
	text := readFile("day1_pt2.txt")

	for _, str := range text {
		digArray := searchString(str)
		digits = append(digits, digArray)
	}

	sum := 0
	for _, d := range digits {
		strVal := ""
		if len(d) == 1 {
			strVal = fmt.Sprint(d[0]) + fmt.Sprint(d[0])
		}

		strVal = fmt.Sprint(d[0]) + fmt.Sprint(d[len(d)-1])
		intVal, _ := strconv.Atoi(strVal)
		sum += intVal
	}

	fmt.Println(fmt.Sprint(sum))
}

func searchString(str string) []int {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var positions []int

	for i := 0; i < len(str); i++ {
		r := rune(str[i])
		if unicode.IsDigit(r) {
			intVal, _ := strconv.Atoi(string(str[i]))
			positions = append(positions, intVal)
		}

		if r == 'o' || r == 't' || r == 'f' || r == 's' || r == 'e' || r == 'n' {
			for key, value := range digits {
				if len(key) > len(str)-i {
					continue
				}

				if key == str[i:i+len(key)] {
					positions = append(positions, value)
				}
			}
		}
	}

	return positions
}
