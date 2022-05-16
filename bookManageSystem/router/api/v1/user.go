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
	CreateAt time.Time //创建时间
	UserType string //1管理员 0普通用户
}

func(u *User)Sign(c *gin.Context){
	// 获取请求参数
	user_name, ok:=c.GetPostForm("user_name")
	password, ok:=c.GetPostForm("password")
	user_type, ok:=c.GetPostForm("user_type")
	if !ok{
		utils.NewResponse(c).ToErrorResponse(404,"parse user message error")
		c.Abort()
		return
	}
	nowTime:=time.Now()
	user := User{ UserName: user_name, Password: password, CreateAt: nowTime,UserType: user_type }
	token,err := utils.GenerateToken(user_name,password)
	if err!=nil{
		utils.NewResponse(c).ToErrorResponse(500,"GenerateToken error")
		c.Abort()
		return
	}
	// 插入库表
	(utils.DB).Create(&user)
	utils.NewResponse(c).ToResponse("sign user success",gin.H{"token":token})
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