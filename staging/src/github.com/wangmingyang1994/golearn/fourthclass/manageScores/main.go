package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	//定义map容器，分配内存空间
	var Scores map[string]int
	Scores = make(map[string]int, 20)
	//初始化值
	for i := 0; i < 20; i++ {
		name := fmt.Sprintf("name%d", i)
		score := rand.Intn(100)
		Scores[name] = score
	}
	fmt.Println("20名学生分数的初始值是：",Scores)
	fmt.Println("==========")

	var score_all, student_num int
	for name, _ := range Scores {
		score, ok := Scores[name]
		if ok {
			score_all += score
			student_num++
		}
		continue
	}
	score_avg := score_all / student_num
	fmt.Printf("%d个学生的平均分是%d\n", student_num, score_avg)
	fmt.Println("==========")
	var rankOfScores [][]string
	for k, v := range Scores {
		var s = make([]string, 2)
		s[0] = k
		s[1] = strconv.Itoa(v)
		rankOfScores = append(rankOfScores, s)
	}
	for i := 0; i < len(rankOfScores)-1; i++ {
		for j := i + 1; j < len(rankOfScores); j++ {
			if rankOfScores[i][1] < rankOfScores[j][1] {
				rankOfScores[i], rankOfScores[j] = rankOfScores[j], rankOfScores[i]
			}
		}
	}
	fmt.Println("20名学生分数排序后：",rankOfScores)


}
