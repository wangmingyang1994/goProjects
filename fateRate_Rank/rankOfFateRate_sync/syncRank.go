package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
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
	ranks.Insert(p)
}

type RankOfFateRate struct {
	Data []Person
	sync.Mutex
}

func (r *RankOfFateRate) Insert(p Person) (rankNum int) {
	//判断插入用户的体脂率是否符合规范
	if p.FateRate <= 0 || p.FateRate > 0.4 {
		fmt.Printf("%s的体脂率%f不符合规则！请重试\n", p.Name, p.FateRate)
		return
	}
	r.Lock()
	defer r.Unlock()
	//将用户加入排行榜
	r.Data = append(r.Data, p)
	//重新排序
	r.SortRank()
	//返回插入用户的排行
	rankNum, _ = r.Search(p.Id)
	fmt.Println(p.Name, "插入成功！")
	//fmt.Printf("插入成功！%s的排行是第%d，体脂率是%f\n", p.Name, rankNum, p.FateRate)
	return rankNum
}

func (r *RankOfFateRate) SortRank() {
	//排行榜按用户的体脂率从小到大排序
	sort.Slice(r.Data, func(i, j int) bool {
		return r.Data[i].FateRate < r.Data[j].FateRate
	})
}

func (r *RankOfFateRate) Search(id int) (rankNum int, person Person) {
	//遍历排行榜，按ID查询到用户，获取其排行和体脂率并返回
	for i, v := range r.Data {
		if v.Id != id {
			continue
		}
		rankNum = i + 1
		person = v
		break
	}
	//fmt.Printf("查询到用户%s的排名是%d\n",person.Name,rankNum)
	return rankNum, person

}

func (r *RankOfFateRate) Update(id int, fateRate float64) (rankNum int) {
	//判断体脂率是否符合规则
	if fateRate <= 0 || fateRate > 0.4 {
		fmt.Printf("体脂率%f不符合规则！请重试\n", fateRate)
		return 0
	}
	//遍历,查询到要更改的ID，并更改其体脂率
	r.Lock()
	defer r.Unlock()
	for i := 0; i < len(r.Data); i++ {
		if r.Data[i].Id != id {
			continue
		}
		r.Data[i].FateRate = fateRate
		break
	}
	//更改完成，重新排序，并返回当前用户的排行
	r.SortRank()
	rankNum, p := r.Search(id)
	//fmt.Printf("更新成功！%s排行是第%d，体脂率是%f\n", v.Name, rankNum, fateRate)
	fmt.Println(p.Name, "更新成功！")
	return rankNum
}

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
			newFateRate = v.FateRate + add
		} else {
			newFateRate = v.FateRate - add
		}
		//判断新生成的体脂率，若符合规范则直接更新，否则重新生成
		if newFateRate < 0 || newFateRate > 0.4 {
			continue
		}
		rank.Update(id, newFateRate)
		wg.Done()
		break
	}
}

func main() {
	//初始化排行榜
	data := make([]Person, 0)
	rank := &RankOfFateRate{data, sync.Mutex{}}
	//初始化两个用户
	name := "user1"
	fateRate := 0.312
	InitPerson(name, fateRate, rank)
	name1 := "user2"
	fateRate1 := 0.344
	InitPerson(name1, fateRate1, rank)

	fmt.Println("排行榜：", rank.Data)
	//将第二个用户体脂率更新
	rank.Update(1, 0.211)
	fmt.Println("排行榜：", rank.Data)

}
