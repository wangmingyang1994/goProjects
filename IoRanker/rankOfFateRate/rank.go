package rankOfFateRate

import (
	"encoding/json"
	"goProjects/IoRanker/writer"
	"log"
)

type Person struct {
	Name    string  `json:"name"`
	FatRate float64 `json:"fatRate"`
	Order   int     `json:"order"`
}

type Rank struct {
	RankDetail []Person
	QueueCh    chan Person
}

func (rank *Rank) Insert(person Person) int {
	//插入用户排行表
	rank.QueueCh <- person
	rank.RankDetail = append(rank.RankDetail, <-rank.QueueCh)
	//引入sorter包中的不同排序算法，进行排序
	//sort := Bubble{}
	sort := Quick{}
	sort.SortRank(&rank.RankDetail, 0, 0)
	//写入到文件
	order := rank.GetMyRank(person.Name)
	person.Order = order
	d, err := json.Marshal(person)
	if err != nil {
		log.Fatalln("json序列化失败：", err)
	}
	data := append([]byte(d), []byte("\n")...)
	writer.Write_to_file(data)
	return order
}

func (rank *Rank) GetMyRank(name string) int {
	//获取我的名次，并更新到排行榜中对应的用户
	count := -1
	for i, value := range rank.RankDetail {
		if value.Name == name {
			count = i + 1
			rank.RankDetail[i].Order = count
		}
	}
	return count
}
