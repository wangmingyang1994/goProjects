package main

import "fmt"

func main(){
	const pi = 3.14
	var circle1_bj float64= 3
	circle1 := pi *(circle1_bj*circle1_bj)
	var circle2_bj float64= 5
	circle2 := pi *(circle2_bj*circle2_bj)
	fmt.Printf("圆1的面积是%.3f\n", circle1)
	fmt.Printf("圆2的面积是%.3f\n", circle2)
	fmt.Printf("圆1+圆2的面积和是：%.3f\n", circle1+circle2)
}

