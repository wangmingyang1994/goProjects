package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestRankOfFateRate(t *testing.T) {
	data := make([]Person, 0)
	rank := &RankOfFateRate{data, sync.Mutex{}}
	//设置随机种子
	rand.Seed(time.Now().Unix())
	//设置一个队列
	var wg sync.WaitGroup
	//初始化1000个用户，插入排行榜
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go MockInsert(i, rank, &wg)
		wg.Wait()
	}

	for {
		//更新1000个用户的体脂率
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go MockUpdate(rank, i, &wg)
			wg.Wait()
		}
	}

}
