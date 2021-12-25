package main

import "fmt"

func main(){
	var nums = [10]int{0,1,2,3,4,5,6,7,8,9}
	var nums_new =make([]int,0)
	for i:=0;i<len(nums);i++{
		nums_new=append(nums_new,nums[len(nums)-i-1])
	}
	fmt.Println(nums_new)
}
