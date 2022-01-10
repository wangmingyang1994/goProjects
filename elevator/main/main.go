package main

import (
	"fmt"
	"goProjects/elevator"
	"time"
)

func main() {
	e1 := elevator.Elevator{
		3,
		[]int{},
		[]int{},
		time.Duration(time.Second),
		false,
		false,
	}
	passagers := []elevator.Passenger{
		{
			3,
			4,
		},
		{

			3,
			5,
		},
		{
			3,
			2,
		},
	}
	fmt.Println("初始化：电梯在三楼，三楼乘客分别去4楼、5楼和2楼")
	e1.Get(passagers)
	s := e1.Cost()
	fmt.Println()
	fmt.Println("经停的楼层有：")
	for _, v := range s {
		fmt.Printf("%d楼\n", v)
	}
}
