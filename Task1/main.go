package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	mapForCheck := make(map[int]bool)
	for _, value := range arr {
		if _, ok := mapForCheck[value]; ok {
			continue
		}
		mapForCheck[value] = true
		result = append(result, value)
	}
	fmt.Println(result)
}
