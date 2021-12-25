package main

import (
	"fmt"
	"math/rand"
)

func sorter(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s
}

func main() {
	var slice = make([]int, 0)
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.Intn(100))
	}
	fmt.Println(slice)
	result := sorter(slice)
	fmt.Println(result)
}
