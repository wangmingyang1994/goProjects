package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Person struct {
	Name     string
	FateRate float64
	Id       int
}

func InitPerson(name string, faterate float64, ranks *RankOfFateRate) {
	// 获取排行榜当前的人数，并将ID（人数+1）赋给最新的用户
	index := len(ranks.Data)
	// 初始化用户
	p := Person{Name: name, FateRate: faterate, Id: index}
	// 插入排行，获取到当前排名，输出到控制台
	r := ranks.Insert(p)
	fmt.Printf("id:%d,%s的排行是第%d，体脂率是%f\n", index, p.Name, r, p.FateRate)

}

type RankOfFateRate struct {
	Data []Person
	ch   chan Person
}

func (r *RankOfFateRate) Insert(p Person) (rankNum int) {
	//判断插入用户的体脂率是否符合规范
	if p.FateRate <= 0 || p.FateRate > 0.4 {
		fmt.Printf("%s的体脂率%f不符合规则！请重试\n", p.Name, p.FateRate)
		return
	}
	r.ch <- p
	//将用户加入排行榜
	r.Data = append(r.Data, <-r.ch)
	//重新排序
	r.SortRank()
	//返回插入用户的排行
	rankNum, _ = r.Search(p.Id)
	return rankNum
}

func (r *RankOfFateRate) SortRank() {
	//排行榜按用户的体脂率从小到大排序
	sort.Slice(r.Data, func(i, j int) bool {
		return r.Data[i].FateRate < r.Data[j].FateRate
	})
}

func (r *RankOfFateRate) Search(id int) (num int, fateRate float64) {
	//遍历排行榜，按ID查询到用户，获取其排行和体脂率并返回
	for i, v := range r.Data {
		if v.Id != id {
			continue
		}
		num = i + 1
		fateRate = v.FateRate
		break
	}
	return num, fateRate

}

func (r *RankOfFateRate) Update(id int, fateRate float64) (rankNum int, err error) {
	//判断体脂率是否符合规则
	if fateRate <= 0 || fateRate > 0.4 {
		message := fmt.Sprintf("体脂率%f不符合规则！请重试\n", fateRate)
		return 0, errors.New(message)
	}
	//遍历,查询到要更改的ID，并更改其体脂率
	for i := 0; i < len(r.Data); i++ {
		if r.Data[i].Id != id {
			continue
		}
		name := r.Data[i].Name
		updatePerson := Person{name, fateRate, id}
		r.ch <- updatePerson
		r.Data[i] = <-r.ch

		break
	}
	//更改完成，重新排序，并返回当前用户的排行
	r.SortRank()
	rankNum, _ = r.Search(id)
	fmt.Printf("%d排行是%d,体脂率是%f\n", id, rankNum, fateRate)
	return rankNum, nil
}

//func (r *RankOfFateRate) Delete(id int){
//	//遍历排行榜，按ID找到用户，并将其在slice中删除
//	for i,v := range r.Data{
//		if v.Id !=id{
//			continue
//		}
//		r.Data = append(r.Data[i:],r.Data[i+1:]...)
//	}
//}

func MockInsert(n int, rank *RankOfFateRate, wg *sync.WaitGroup) {
	//初始化用户名
	name := fmt.Sprintf("name%d", n)
	for {
		//获取一个随机数，若符合体脂率的标准，则有效。
		fateRate := rand.Float64()
		if fateRate < 0 || fateRate > 0.4 {
			continue
		}
		//插入到排行榜中
		InitPerson(name, fateRate, rank)
		wg.Done()
		break
	}
}

func MockUpdate(rank *RankOfFateRate, id int, wg *sync.WaitGroup) {
	for {
		//获取一个0-0.2之间的随机数，作为体脂率要加减的数
		add := rand.Float64()
		if add < 0 || add > 0.2 {
			continue
		}
		//根据ID获取要更改用户的体脂率
		_, v := rank.Search(id)
		//获取一个0-1之间的随机数，若是偶数，则为加，否则为减
		x := rand.Intn(10)
		var newFateRate float64
		if x%2 == 0 {
			newFateRate = v + add
		} else {
			newFateRate = v - add
		}
		//判断新生成的体脂率，若符合规范则直接更新，否则重新生成
		if newFateRate < 0 || newFateRate > 0.4 {
			continue
		}
		_, err := rank.Update(id, newFateRate)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
		break
	}
}

func main() {
	//初始化排行榜
	data := make([]Person, 0)
	ch := make(chan Person, 1)
	rank := &RankOfFateRate{data, ch}
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
	//将1000个用户，体脂率更新
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go MockUpdate(rank, i, &wg)
		wg.Wait()
	}

}
