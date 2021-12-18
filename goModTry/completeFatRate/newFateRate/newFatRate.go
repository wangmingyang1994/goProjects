package newFateRate

import (
	"fmt"
	"github.com/spf13/cobra"
	g "goProjects/FatRateCalc"
)

func InputfateRate() (name string, fatRate float64) {
	var (
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
	//使用本地体脂率计算实现
	bmi := g.GetBMI(weight, tall)
	if sex == "男" {
		fatRate := g.GetmanFatRate(bmi, age)
		return name, fatRate
	}
	return name, g.GetwomanFatRate(bmi, age)

}
