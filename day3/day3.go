package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readFile(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	strArr := []string{}

	for scanner.Scan() {
		strArr = append(strArr, scanner.Text())
	}

	return strArr
}

func addNumAndSpecialCharsToMap() map[string]string {
	input := readFile("day3.txt")
	validNumCharMap := make(map[string]string)

	for i := 0; i < len(input); i++ {
		str := input[i]

		xIndex := -1
		chars := ""
		for j := 0; j < len(str); j++ {
			if unicode.IsDigit(rune(str[j])) {
				if xIndex == -1 {
					xIndex = j
				}
				chars += string(str[j])
			} else {
				if len(chars) > 0 {
					validNumCharMap[fmt.Sprint(i)+","+fmt.Sprint(xIndex)] = chars
					xIndex = -1
				}

				if rune(str[j]) != '.' {
					validNumCharMap[fmt.Sprint(i)+","+fmt.Sprint(j)] = string(str[j])
				}
				chars = ""
			}

			if j == len(str)-1 && len(chars) > 0 {
				validNumCharMap[fmt.Sprint(i)+","+fmt.Sprint(xIndex)] = chars
			}
		}

	}

	return validNumCharMap
}

func findValidNumbers() {
	strMap := addNumAndSpecialCharsToMap()
	sum := 0
	sum2 := 0

	for key, value := range strMap {
		if unicode.IsDigit(rune(value[0])) && isValidNum(value, key, strMap) {
			intVal, _ := strconv.Atoi(value)

			sum += intVal
		}
	}

	fmt.Println(sum)

	for key, value := range strMap {
		if unicode.IsDigit(rune(value[0])) {
			sum2 += getValidNum(value, key, strMap)
		}
	}

	fmt.Println(sum2)
}

func isValidNum(num string, key string, strMap map[string]string) bool {
	yIndex, _ := strconv.Atoi(strings.Split(key, ",")[0])
	xIndex, _ := strconv.Atoi(strings.Split(key, ",")[1])
	strIndex := ""
	val := ""

	for i := 0; i < len(num); i++ {
		// check left
		strIndex = strconv.Itoa(yIndex) + "," + strconv.Itoa(xIndex+i-1)
		val = strMap[strIndex]
		if checkForSpecialChar(val) {
			return true
		}

		// check right
		strIndex = strconv.Itoa(yIndex) + "," + strconv.Itoa(xIndex+i+1)
		val = strMap[strIndex]
		if checkForSpecialChar(val) {
			return true
		}

		// check bottom
		strIndex = strconv.Itoa(yIndex+1) + "," + strconv.Itoa(xIndex+i)
		val = strMap[strIndex]
		if checkForSpecialChar(val) {
			return true
		}

		// check top
		strIndex = strconv.Itoa(yIndex-1) + "," + strconv.Itoa(xIndex+i)
		val = strMap[strIndex]
		if checkForSpecialChar(val) {
			return true
		}

		// check diagonal top left
		strIndex = strconv.Itoa(yIndex-1) + "," + strconv.Itoa(xIndex+i-1)
		val = strMap[strIndex]

		if checkForSpecialChar(val) {
			return true
		}

		// check diagonal top right
		strIndex = strconv.Itoa(yIndex-1) + "," + strconv.Itoa(xIndex+i+1)
		val = strMap[strIndex]
		if checkForSpecialChar(val) {
			return true
		}

		// check diagonal bottom left
		strIndex = strconv.Itoa(yIndex+1) + "," + strconv.Itoa(xIndex+i-1)
		val = strMap[strIndex]
		if checkForSpecialChar(val) {
			return true
		}

		// check diagonal bottom right
		strIndex = strconv.Itoa(yIndex+1) + "," + strconv.Itoa(xIndex+i+1)
		val = strMap[strIndex]
		if checkForSpecialChar(val) {
			return true
		}
	}

	return false
}

func getValidNum(num string, key string, strMap map[string]string) int {
	yIndex, _ := strconv.Atoi(strings.Split(key, ",")[0])
	xIndex, _ := strconv.Atoi(strings.Split(key, ",")[1])
	strIndex := ""
	val := ""
	strNum := ""
	specialChar := ""

	for i := 0; i < len(num); i++ {
		// check left
		strIndex = strconv.Itoa(yIndex) + "," + strconv.Itoa(xIndex+i-1)
		val = strMap[strIndex]
		specialChar = getSpecialChar(val)
		if specialChar != "" {
			strNum = findValidNumForSpecialChar(false, false, false, false, xIndex+i-1, yIndex, strMap)

			if strNum != "" && unicode.IsDigit(rune(strNum[0])) {
				num1, _ := strconv.Atoi(num)
				num2, _ := strconv.Atoi(strNum)

				fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)

				return num1 * num2
			}
		}

		// check right
		strIndex = strconv.Itoa(yIndex) + "," + strconv.Itoa(xIndex+i+1)
		val = strMap[strIndex]
		specialChar = getSpecialChar(val)
		if specialChar != "" {
			strNum = findValidNumForSpecialChar(false, false, false, true, xIndex+i+1, yIndex, strMap)

			if strNum != "" && unicode.IsDigit(rune(strNum[0])) {
				num1, _ := strconv.Atoi(num)
				num2, _ := strconv.Atoi(strNum)

				fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)

				return num1 * num2
			}
		}

		// check bottom
		strIndex = strconv.Itoa(yIndex+1) + "," + strconv.Itoa(xIndex+i)
		val = strMap[strIndex]
		specialChar = getSpecialChar(val)
		if specialChar != "" {
			strNum = findValidNumForSpecialChar(false, false, true, true, xIndex+i, yIndex+1, strMap)

			if strNum != "" && unicode.IsDigit(rune(strNum[0])) {
				num1, _ := strconv.Atoi(num)
				num2, _ := strconv.Atoi(strNum)

				fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)

				return num1 * num2
			}
		}

		// check diagonal bottom left
		strIndex = strconv.Itoa(yIndex+1) + "," + strconv.Itoa(xIndex+i-1)
		val = strMap[strIndex]
		specialChar = getSpecialChar(val)
		if specialChar != "" {
			strNum = findValidNumForSpecialChar(false, false, true, true, xIndex+i-1, yIndex+1, strMap)

			if strNum != "" && unicode.IsDigit(rune(strNum[0])) {
				num1, _ := strconv.Atoi(num)
				num2, _ := strconv.Atoi(strNum)

				fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)

				return num1 * num2
			}
		}

		// check diagonal bottom right
		strIndex = strconv.Itoa(yIndex+1) + "," + strconv.Itoa(xIndex+i+1)
		val = strMap[strIndex]
		specialChar = getSpecialChar(val)
		if specialChar != "" {
			strNum = findValidNumForSpecialChar(false, true, true, true, xIndex+i+1, yIndex+1, strMap)
			if strNum != "" && unicode.IsDigit(rune(strNum[0])) {
				num1, _ := strconv.Atoi(num)
				num2, _ := strconv.Atoi(strNum)

				fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)

				return num1 * num2
			}
		}
	}

	// check neighbor number
	strIndex = strconv.Itoa(yIndex) + "," + strconv.Itoa(xIndex+len(num)+1)
	strNum = strMap[strIndex]

	if strNum != "" && unicode.IsDigit(rune(strNum[0])) {
		// check top right
		strIndex = strconv.Itoa(yIndex-1) + "," + strconv.Itoa(xIndex+len(num))
		val = strMap[strIndex]
		specialChar = getSpecialChar(val)

		if specialChar != "" {
			num1, _ := strconv.Atoi(num)
			num2, _ := strconv.Atoi(strNum)

			fmt.Printf("Num1: %d, Num2: %d\n", num1, num2)

			return num1 * num2
		}
	}

	return 0
}

func findValidNumForSpecialChar(checkTopLeft bool, checkTopRight bool, checkLeft bool, checkRight bool, xIndex int, yIndex int, strMap map[string]string) string {
	str := ""

	str = checkBottomArr(strMap, yIndex, xIndex)
	if str != "" {
		return str
	}

	if checkLeft {
		str = checkLeftArr(strMap, yIndex, xIndex)
		if str != "" {
			return str
		}
	}

	if checkRight {
		str = checkRightArr(strMap, yIndex, xIndex)
		if str != "" {
			return str
		}
	}

	if checkTopRight {
		str = checkTopRightArr(strMap, yIndex, xIndex)

		if str != "" {
			return str
		}
	}

	return ""
}

func checkLeftArr(strMap map[string]string, yIndex int, xIndex int) string {
	if strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex-1)] != "" {
		return strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex-1)]
	}
	if strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex-2)] != "" {
		return strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex-2)]
	}
	if strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex-3)] != "" {
		if len(strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex-3)]) == 3 {
			return strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex-3)]
		}
	}

	return ""
}

func checkRightArr(strMap map[string]string, yIndex int, xIndex int) string {
	if strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex+1)] != "" {
		return strMap[strconv.Itoa(yIndex)+","+strconv.Itoa(xIndex+1)]
	}

	return ""
}

func checkBottomArr(strMap map[string]string, yIndex int, xIndex int) string {
	str := ""

	str = strMap[strconv.Itoa(yIndex+1)+","+strconv.Itoa(xIndex)]
	if str != "" && unicode.IsDigit(rune(str[0])) {
		return str
	}

	str = strMap[strconv.Itoa(yIndex+1)+","+strconv.Itoa(xIndex-1)]
	if str != "" && unicode.IsDigit(rune(str[0])) {
		return str
	}

	str = strMap[strconv.Itoa(yIndex+1)+","+strconv.Itoa(xIndex-2)]
	if str != "" && unicode.IsDigit(rune(str[0])) {
		return str
	}

	str = strMap[strconv.Itoa(yIndex+1)+","+strconv.Itoa(xIndex-3)]
	if str != "" && unicode.IsDigit(rune(str[0])) && len(str) == 3 {
		return str
	}

	str = strMap[strconv.Itoa(yIndex+1)+","+strconv.Itoa(xIndex+1)]
	if str != "" && unicode.IsDigit(rune(str[0])) {
		return str
	}

	return ""
}

func checkTopRightArr(strMap map[string]string, yIndex int, xIndex int) string {
	if strMap[strconv.Itoa(yIndex-1)+","+strconv.Itoa(xIndex+1)] != "" {
		return strMap[strconv.Itoa(yIndex-1)+","+strconv.Itoa(xIndex+1)]
	}

	return ""
}

func checkForSpecialChar(value string) bool {
	if len(value) == 1 && !unicode.IsDigit(rune(value[0])) && rune(value[0]) != '.' {
		return true
	}

	return false
}

func getSpecialChar(value string) string {
	if len(value) == 1 && !unicode.IsDigit(rune(value[0])) && rune(value[0]) != '.' {
		return string(value[0])
	}

	return ""
}
