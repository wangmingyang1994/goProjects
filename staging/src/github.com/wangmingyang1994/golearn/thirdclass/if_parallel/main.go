package main

import "fmt"

func getKofLine(p []float64) float64 {
	return (p[3] - p[1]) / (p[2] - p[0])
}
func main() {
	for {
		var point1 []float64
		//输入第一条直线上的两个点（x1,y1） （x2,y2）
		for i := 0; i < 2; i++ {
			var x, y float64
			fmt.Printf("请输入第一条直线的第%d个点坐标(格式 =>如x y)：", i+1)
			fmt.Scanln(&x, &y)
			point1 = append(point1, x, y)
		}
		//计算出此直线的斜率
		line1 := getKofLine(point1)
		//继续输入第二条直线上的两个点（x1,y1） （x2,y2）
		var point2 []float64
		//输入第一条直线上的两个点（x1,y1） （x2,y2）
		for i := 0; i < 2; i++ {
			var x, y float64
			fmt.Printf("请输入第二条直线的第%d个点坐标(格式 =>如x y)：", i+1)
			fmt.Scanln(&x, &y)
			point2 = append(point2, x, y)
		}
		//计算出第二条直线的斜率
		line2 := getKofLine(point2)
		//判断两条直线的斜率是否相等，输出结论
		if line1 == line2 {
			fmt.Println("两条直线平行")
			return
		}
		fmt.Println("两条直线不平行")
		return
	}

}
