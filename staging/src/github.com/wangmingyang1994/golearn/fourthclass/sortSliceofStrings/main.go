package main

import (
	"fmt"
	"strings"
)

func sortSliceOfStrings(s []string) []string {
	for i := 0; i < len(s); i++ {
		n := []rune(s[i])
		if n[0] >= 'a' && n[0] <= 'z' {
			s[i] = strings.ToUpper(s[i])
		} else if n[0] >= 'A' && n[0] <= 'Z' {
			s[i] = strings.ToLower(s[i])
		} else {
			fmt.Printf("%v is no need to trans\n", s[i])
		}
	}
	return s
}

func main() {

	var strs = []string{
		"abc", "BGS", "hello", "哈哈哈你好",
	}
	new_strs := sortSliceOfStrings(strs)
	fmt.Println("result is :", new_strs)
}
