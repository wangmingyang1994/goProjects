package main
import "fmt"

func getYear(year int) [12][31]int{
	var yearDetail [12][31]int
	for i:=1;i<13;i++{
		month := [31]int{}
		length:=0

		if i==4|| i==6||i==9||i==11{
			length =30

		}else if i ==2{
			if year%4 ==0{
				length=29
			}else{
				length=28
			}
		}else{
			length=31
		}
		for j:=0;j<length;j++{
			month[j]=j+1
		}
		yearDetail[i-1]=month

	}
	return yearDetail

}

func main(){
	var year int
	fmt.Print("请输入年份：")
	fmt.Scanln(&year)
	yearDetail:=getYear(year)
	fmt.Println(yearDetail)
}
