package main

import "fmt"

func TestFat(){
	fmt.Println("欢迎使用体脂计算器～")
	var name string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	var weight float64
	fmt.Print("请输入体重（kg）：")
	fmt.Scanln(&weight)
	for weight <5 || weight>1000{
			fmt.Println("体重不合法，请再次输入：体重（公斤）：")
			fmt.Scanln(&weight)
	}

	var tall float64
	fmt.Print("请输入身高（m）：")
	fmt.Scanln(&tall)
	for tall <0.5 || tall>3.0{
		fmt.Println("身高不合法，请再次输入：身高（m）：")
		fmt.Scanln(&tall)
	}
	var bmi float64 = weight / (tall * tall)

	var age int
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	for age <0 || age>200{
		fmt.Println("年龄不合法，请再次输入：年龄（岁）：")
		fmt.Scanln(&age)
	}
	var sex string
	fmt.Print("请输入性别：（男/女）")
	fmt.Scanln(&sex)

	var sexWeight float64
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	var fatRate float64 = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*sexWeight)/100
	fmt.Println("姓名：", name)
	fmt.Printf("体脂率是：%.2f\n",fatRate)

	if sex == "男" {
		if age >=18 && age <=39{
			if fatRate <=0.1{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.1&&fatRate<=0.16{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.16&&fatRate<=0.21 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.21&&fatRate<=0.26{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age >=40 &&age <=59{
			if fatRate <=0.11{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.11&&fatRate<=0.17{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.17&&fatRate<=0.22 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.22&&fatRate<=0.27{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age>=60{
			if fatRate <=0.13{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.13&&fatRate<=0.19{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.19&&fatRate<=0.24 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.24&&fatRate<=0.29{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else{
			fmt.Println("暂时不支持计算未成年人的体脂率～")
		}
	} else {
		if age >=18 && age <=39{
			if fatRate <=0.2{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.2&&fatRate<=0.27{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.27&&fatRate<=0.34 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.34&&fatRate<=0.39{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age >=40 &&age <=59{
			if fatRate <=0.21{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.21&&fatRate<=0.28{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.28&&fatRate<=0.35 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.35&&fatRate<=0.41{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else if age>=60{
			if fatRate <=0.22{
				fmt.Println("目前偏瘦，多吃多锻炼")
			}else if fatRate>0.22&&fatRate<=0.29{
				fmt.Println("目前标准，太棒了，要保持哦")
			}else if fatRate>0.29&&fatRate<=0.36 {
				fmt.Println("目前偏重，吃完饭多散散步，消化消化")
			}else if fatRate>0.36&&fatRate<=0.41{
				fmt.Println("目前肥胖，开始锻炼锻炼吧～胖砸")
			}else{
				fmt.Println("目前严重肥胖！！！管住嘴，迈开腿，开始减肥吧！！！")
			}
		}else{
			fmt.Println("暂时不支持计算未成年人的体脂率～")
		}
	}
}
func main() {
	TestFat()
	for{
		var ifcontinue string
		fmt.Print("是否继续？（y/n）")
		fmt.Scanln(&ifcontinue)
		if ifcontinue == "y"{
			TestFat()
		}else{
			return
		}
	}


}
