package FatRateCalc

import (
	"fmt"
	"github.com/spf13/cobra"
	fatRates "github.com/wangmingyang1994/golearn/fatRate"
)

func InputPrintfateRate() {
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)
	cmd := &cobra.Command{
		Use:  "fateRateCalc",
		Long: "计算您的体脂率！",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("name: ", name)
			fmt.Println("sex: ", sex)
			fmt.Println("tall: ", tall)
			fmt.Println("weight: ", weight)
			fmt.Println("age: ", age)

		},
	}
	//绑定参数
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")
	//执行命令
	cmd.Execute()
	//使用替换后的本地体脂率计算实现
	bmi, err := fatRates.GetBMI(tall, weight)
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err1 := fatRates.GetFatRate(bmi, age, sex)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	suggest := fatRates.GetSuggestion(sex, age, f)
	fmt.Printf("哈喽%s!您的体脂率是：%f,%s", name, f, suggest)

}
