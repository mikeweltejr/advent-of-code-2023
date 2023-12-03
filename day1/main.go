package main

import (
	"fmt"
)

func main() {
	digitStrings := findFirstAndLast()
	sum := sum(digitStrings)

	fmt.Println(sum)

	findNumbers()
}
