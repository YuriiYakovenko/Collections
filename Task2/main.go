package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string
	splitedInput := strings.Split(input, " ")
	max := splitedInput[0]
	min := splitedInput[0]
	for i := 0; i < len(splitedInput); i++ {
		if splitedInput[i] > max {
			max = splitedInput[i]
		}
		if splitedInput[i] < min {
			min = splitedInput[i]
		}
	}
	result = max + " " + min
	fmt.Println(result)
}
