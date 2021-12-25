package main
import (
	"fmt"
	"math/rand")

func randPoker(s []string) []string{
	var new = make([]string,0)
	for i :=0;i<54;i++{
		l := len(s)
		rand_int := rand.Intn(l)
		ele := s[rand_int]
		new = append(new, ele)
		s = append(s[:rand_int],s[rand_int+1:]...)
	}
	return new

}
func main(){
	var s = make([]string,0)
	//生成扑克牌
	for i:=1;i<14;i++{
		a := fmt.Sprintf("红桃%d",i)
		b := fmt.Sprintf("黑桃%d",i)
		c := fmt.Sprintf("方片%d",i)
		d := fmt.Sprintf("梅花%d",i)
		s = append(s,a,b,c,d)
	}
	var king1,king2 string = "大王","小王"
	s =append(s,king1,king2)
	fmt.Println(s)
	fmt.Println(len(s))
	//乱序后的扑克牌
	new := randPoker(s)
	fmt.Println(new)
	fmt.Println(len(new))



}
