package go_bmi


import (
"fmt"
"testing"
)

//• 为体脂计算器编写单元测试并完善体脂计算器的验证逻辑
//• BMI计算
//1. 录入正常身高、体重，确保计算结果符合预期
//2. 录入0或负数身高，返回错误
//3. 录入0或负数体重，返回错误
func TestGetBMI(t *testing.T) {
	tests := []struct {
		tall, weight, bmi float64
	}{
		// 录入正常身高、体重，确保计算结果符合预期
		{1.70, 70, 24.221453287197235},
		//录入0或负数身高，返回错误
		{0, 80, 0},
		{-1, 80, 0},
		{3.1, 80, 0},
		// 录入0或负数体重，返回错误
		{1.7, 0, 0},
		{1.8, -2, 0},
	}
	for _, tt := range tests {
		actual, err := GetBMI(tt.tall, tt.weight)
		if actual != tt.bmi {
			fmt.Println(err)
			t.Fail()
		}
	}
}

//为体脂计算器编写单元测试并完善体脂计算器的验证逻辑
//• 体脂率计算
//1. 录入正常BMI、年龄、性别，确保计算结果符合预期
//2. 录入非法BMI、年龄、性别（0、负数、超过150的年龄、非男女的性别输入），返回错误
//3. 录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期

func TestGetFatRate(t *testing.T) {
	tests := []struct {
		bmi     float64
		age     int
		sex     string
		fatRate float64
	}{
		// 录入正常BMI、年龄、性别，确保计算结果符合预期
		{20, 20, "男", 0.12400000000000003},
		{20, 20, "女", 0.23200000000000004},
		//录入非法BMI，返回错误
		{0, 50, "男", 0},
		{-1, 50, "男", 0},
		//录入非法年龄（0、负数、超过150的年龄，返回错误）
		{0.3, 0, "男", 0},
		{0.3, -1, "女", 0},
		{0.3, 150, "女", 0},
		// 录入非法的性别，返回错误
		{0.3, 20, "0", 0},
	}
	for _, tt := range tests {
		actual, err := GetFatRate(tt.bmi, tt.age, tt.sex)
		if actual != tt.fatRate {
			fmt.Println(actual, err)
			t.Fail()
		}
	}
}

func TestGetSuggestion(t *testing.T) {
	//录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期
	tests := []struct {
		tall float64
		weight float64
		age int
		sex string
		suggest string
	}{
		// 录入正常BMI、年龄、性别，确保计算结果符合预期
		{1.8, 50, 20, "男","您太瘦啦，要多吃多锻炼哦"},
		{1.8, 60, 20, "男","您太棒啦，要保持哦"},
		{1.8, 70, 30, "男","您有点偏重啦，吃完饭多散散步，消化消化哦"},
		{1.8, 90, 20, "男","您有点肥胖啦，平时要开始加强锻炼啦"},
		{1.8, 120, 20, "男","您严重肥胖啦，管住嘴，迈开腿，开始减肥吧！！！"},
		{1.7, 40, 20, "女","您太瘦啦，要多吃多锻炼哦"},
		{1.7, 60, 20, "女","您太棒啦，要保持哦"},
		{1.7, 70, 20, "女","您有点偏重啦，吃完饭多散散步，消化消化哦"},
		{1.7, 90, 20, "女","您有点肥胖啦，平时要开始加强锻炼啦"},
		{1.7, 130, 20, "女","您严重肥胖啦，管住嘴，迈开腿，开始减肥吧！！！"},

	}
	for _,tt := range tests{
		bmi,_ := GetBMI(tt.tall,tt.weight)
		fatRate,_ := GetFatRate(bmi,tt.age,tt.sex)
		suggest1 := GetSuggestion(tt.sex,tt.age,fatRate)
		if suggest1!=tt.suggest{
			fmt.Println(fatRate,tt.age,suggest1)
			t.Fail()
		}
	}
}