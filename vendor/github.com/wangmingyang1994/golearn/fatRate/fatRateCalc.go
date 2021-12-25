package fatRate

import "fmt"

func GetBMI(tall, weight float64) (Bmi float64, err error) {
	if weight <= 0 {
		return 0, fmt.Errorf("体重输入有误：%f", weight)
	}
	if tall <= 0 || tall > 3 {
		return 0, fmt.Errorf("身高输入有误：%f", tall)
	}
	Bmi = weight / (tall * tall)
	return Bmi, nil
}

func GetFatRate(bmi float64, age int, sex string) (FatRate float64, err error) {
	if bmi <= 0 {
		err = fmt.Errorf("bmi输入有误：%f", bmi)
		return 0, err
	}
	if age <= 0 || age >= 150 {
		err = fmt.Errorf("年龄输入有误：%d", age)
		return 0, err
	}
	switch sex {
	case "男":
		FatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*1) / 100
		return FatRate, nil
	case "女":
		FatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*0) / 100
		return FatRate, nil
	default:
		err = fmt.Errorf("性别输入有误：%s", sex)
		return 0, err
	}

}

func GetSuggestion(sex string, age int, fatRate float64) (suggest string) {
	suggestsMap := make(map[string]string, 5)
	suggestsMap["偏瘦"] = "您太瘦啦，要多吃多锻炼哦"
	suggestsMap["标准"] = "您太棒啦，要保持哦"
	suggestsMap["偏重"] = "您有点偏重啦，吃完饭多散散步，消化消化哦"
	suggestsMap["肥胖"] = "您有点肥胖啦，平时要开始加强锻炼啦"
	suggestsMap["严重肥胖"] = "您严重肥胖啦，管住嘴，迈开腿，开始减肥吧！！！"
	if sex == "男" {
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				suggest = suggestsMap["偏瘦"]
				return
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				suggest = suggestsMap["标准"]
				return
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				suggest = suggestsMap["偏重"]
				return
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				suggest = suggestsMap["肥胖"]
				return
			} else {
				suggest = suggestsMap["严重肥胖"]
				return
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.11 {
				suggest = suggestsMap["偏瘦"]
				return
			} else if fatRate > 0.11 && fatRate <= 0.17 {
				suggest = suggestsMap["标准"]
				return
			} else if fatRate > 0.17 && fatRate <= 0.22 {
				suggest = suggestsMap["偏重"]
				return
			} else if fatRate > 0.22 && fatRate <= 0.27 {
				suggest = suggestsMap["肥胖"]
				return
			} else {
				suggest = suggestsMap["严重肥胖"]
				return
			}
		} else if age >= 60 {
			if fatRate <= 0.13 {
				suggest = suggestsMap["偏瘦"]
				return
			} else if fatRate > 0.13 && fatRate <= 0.19 {
				suggest = suggestsMap["标准"]
				return
			} else if fatRate > 0.19 && fatRate <= 0.24 {
				suggest = suggestsMap["偏重"]
				return
			} else if fatRate > 0.24 && fatRate <= 0.29 {
				suggest = suggestsMap["肥胖"]
				return
			} else {
				suggest = suggestsMap["严重肥胖"]
				return
			}
		} else {
			suggest = "暂时不支持计算未成年人的体脂率～"
			return
		}
	} else {
		if age >= 18 && age <= 39 {
			if fatRate <= 0.2 {
				suggest = suggestsMap["偏瘦"]
				return
			} else if fatRate > 0.2 && fatRate <= 0.27 {
				suggest = suggestsMap["标准"]
				return
			} else if fatRate > 0.27 && fatRate <= 0.34 {
				suggest = suggestsMap["偏重"]
				return
			} else if fatRate > 0.34 && fatRate <= 0.39 {
				suggest = suggestsMap["肥胖"]
				return
			} else {
				suggest = suggestsMap["严重肥胖"]
				return
			}
		} else if age >= 40 && age <= 59 {
			if fatRate <= 0.21 {
				suggest = suggestsMap["偏瘦"]
				return
			} else if fatRate > 0.21 && fatRate <= 0.28 {
				suggest = suggestsMap["标准"]
				return
			} else if fatRate > 0.28 && fatRate <= 0.35 {
				suggest = suggestsMap["偏重"]
				return
			} else if fatRate > 0.35 && fatRate <= 0.41 {
				suggest = suggestsMap["肥胖"]
				return
			} else {
				suggest = suggestsMap["严重肥胖"]
				return
			}
		} else if age >= 60 {
			if fatRate <= 0.22 {
				suggest = suggestsMap["偏瘦"]
				return
			} else if fatRate > 0.22 && fatRate <= 0.29 {
				suggest = suggestsMap["标准"]
				return
			} else if fatRate > 0.29 && fatRate <= 0.36 {
				suggest = suggestsMap["偏重"]
				return
			} else if fatRate > 0.36 && fatRate <= 0.41 {
				suggest = suggestsMap["肥胖"]
				return
			} else {
				suggest = suggestsMap["严重肥胖"]
				return
			}
		} else {
			suggest = "暂时不支持计算未成年人的体脂率～"
			return
		}
	}
}
