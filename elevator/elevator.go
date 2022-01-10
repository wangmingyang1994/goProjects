package elevator

import (
	"fmt"
	"sort"
	"time"
)

type Passenger struct {
	Now    int
	Target int
}

type Elevator struct {
	Floor         int
	Targets       []int
	NextTargets   []int
	Speed         time.Duration
	TargetsUp     bool
	NextTargetsUp bool
}

func (e *Elevator) Get(p []Passenger) {
	if len(p) == 0 {
		fmt.Println("电梯闲置")
		return
	}
	if p[0].Now-e.Floor > 0 {
		e.Targets = append(e.Targets, p[0].Now)
		e.TargetsUp = true
	} else {
		e.Targets = append(e.Targets, p[0].Now)
		e.TargetsUp = true
	}

	e.Targets = append(e.Targets, p[0].Target)

	//第一个乘客的方向（如果>0，上行，否则下行）
	r := p[0].Target - p[0].Now
	if r > 0 {
		for _, v := range p {
			//如果乘客在上行的方向上，添加到任务
			if v.Target > p[0].Now && v.Target != p[0].Target {
				e.Targets = append(e.Targets, v.Target)
				sort.Ints(e.Targets)
				e.TargetsUp = true
			} else if v.Target != p[0].Target {
				e.NextTargets = append(e.NextTargets, v.Target)
				sort.Sort(sort.Reverse(sort.IntSlice(e.NextTargets)))
			}
		}
	} else {
		for _, v := range p {
			//如果乘客在下行的方向上，添加到任务
			if v.Target < p[0].Now && v.Target != p[0].Target {
				e.Targets = append(e.Targets, v.Target)
				sort.Sort(sort.Reverse(sort.IntSlice(e.Targets)))
			} else if v.Target != p[0].Target {
				e.NextTargets = append(e.NextTargets, v.Target)
				sort.Ints(e.NextTargets)
				e.NextTargetsUp = true
			}
		}
	}

}

func (e *Elevator) Cost() (result []int) {
	if len(e.Targets) == 0 && len(e.NextTargets) == 0 {
		fmt.Println("no to run")
		return
	}
	if e.TargetsUp {
		for i := e.Floor; i <= e.Targets[len(e.Targets)-1]; i++ {
			time.Sleep(e.Speed)
			fmt.Printf("到达%d楼\n", i)
			for _, v := range e.Targets {
				if i == v {
					fmt.Println("开门，···，关门")
					result = append(result, i)
				}
			}
		}

	} else {
		for i := e.Floor; i >= e.Targets[len(e.Targets)-1]; i-- {
			time.Sleep(e.Speed)
			fmt.Printf("到达%d楼\n", i)
			for _, v := range e.Targets {
				if i == v {
					fmt.Println("开门，···，关门")
					result = append(result, i)
				}
			}
		}
	}

	if e.NextTargetsUp && len(e.NextTargets) != 0 {
		for i := e.Targets[len(e.Targets)-1] + 1; i <= e.NextTargets[len(e.NextTargets)-1]; i++ {
			time.Sleep(e.Speed)
			fmt.Printf("到达%d楼\n", i)
			for _, v := range e.NextTargets {
				if i == v {
					fmt.Println("开门，···，关门")
					result = append(result, i)
				}
			}
		}

	} else if len(e.NextTargets) != 0 && e.NextTargetsUp == false {
		for i := e.Targets[len(e.Targets)-1] - 1; i >= e.NextTargets[len(e.NextTargets)-1]; i-- {
			time.Sleep(e.Speed)
			fmt.Printf("到达%d楼\n", i)
			for _, v := range e.NextTargets {
				if i == v {
					fmt.Println("开门，···，关门")
					result = append(result, i)
				}
			}
		}
	}
	return result
}
