package elevator

import (
	"testing"
	"time"
)

func TestCost1(t *testing.T) {
	//案例1：楼层有5层，所有电梯楼层没有人请求电梯，电梯不动
	e1 := Elevator{
		1,
		[]int{},
		[]int{},
		time.Duration(time.Second),
		false,
		false,
	}
	passagers := []Passenger{
	}
	e1.Get(passagers)
	result := e1.Cost()
	if len(result) != 0{
		t.Fatalf("expect 0,but got%d",result)
	}

}
func TestCost2(t *testing.T) {
	//案例2：楼层有5层，电梯在1层。三楼按电梯。电梯向三楼行进，并停在三楼。
	e1 := Elevator{
		1,
		[]int{},
		[]int{},
		time.Duration(time.Second),
		false,
		false,
	}
	passagers := []Passenger{
		{3,3},
	}
	e1.Get(passagers)
	result := e1.Cost()
	if result[0] != 3{
		t.Fatalf("expect 0,but got%d",result)
	}

}

func TestCost3(t *testing.T) {
	//案例3-1：楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、2楼。电梯先向上到4楼，然后转头到2楼，最后停在2楼。
	e1 := Elevator{
		3,
		[]int{},
		[]int{},
		time.Duration(time.Second),
		false,
		false,
	}
	passagers := []Passenger{
		{3,4},
		{3,2},
	}
	e1.Get(passagers)
	result := e1.Cost()
	expect := []int{
		3,4,2,
	}
	for i,v :=  range result{
		if v!=expect[i]{
			t.Fail()
		}
	}

}

func TestCost4(t *testing.T) {
	//案例3-2：楼层有5层，电梯在3层。上来一些人后，目标楼层： 4楼、5楼、2楼。电梯先向上到4楼，然后到5楼，之后转头到2楼，最后停在2楼。
	e1 := Elevator{
		3,
		[]int{},
		[]int{},
		time.Duration(time.Second),
		false,
		false,
	}
	passagers := []Passenger{
		{3,4},
		{3,5},
		{3,2},
	}
	e1.Get(passagers)
	result := e1.Cost()
	expect := []int{
		3,4,5,2,
	}
	for i,v :=  range result{
		if v!=expect[i]{
			t.Fail()
		}
	}

}

