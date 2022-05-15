package v1

import (
	"github.com/gin-gonic/gin"
	"goProjects/bookManageSystem/utils"
	"time"
)

type User struct{
	UserId int //用户ID
	UserName string //用户名
	Password string //密码
	CreateAt int64 //创建时间
	UserType uint8 //1管理员 0普通用户
}

func(u *User)Sign(c *gin.Context){
	nowTime:=time.Now().Unix()
	user := User{ UserName: "root", Password: "abc", CreateAt: nowTime,UserType: 0}
	(utils.DB).Create(&user)

}

func(u *User)Login(c *gin.Context){

}

func(u *User)Action(c *gin.Context){

}

func(u *User)Detail(c *gin.Context){

}

func NewUser()*User{
	return &User{}
}