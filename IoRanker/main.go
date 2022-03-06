package main

import (
	"fmt"
	"goProjects/IoRanker/rankOfFateRate"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//初始化rank
	rankSlice := make([]rankOfFateRate.Person, 0, 1000)
	rankCh := make(chan rankOfFateRate.Person, 1)
	rank := rankOfFateRate.Rank{
		RankDetail: rankSlice,
		QueueCh:    rankCh,
	}
	//初始化person
	var wg sync.WaitGroup
	rand.Seed(time.Now().Unix())
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(num int) {
			for {
				fRate := rand.Float64()
				if fRate < 0 || fRate > 0.4 {
					continue
				}
				name := fmt.Sprintf("person_%d", i)
				p := rankOfFateRate.Person{
					name,
					fRate,
					0,
				}
				rank.Insert(p)
				break
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
