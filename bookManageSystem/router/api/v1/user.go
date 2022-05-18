package v1

import (
	"github.com/gin-gonic/gin"
	"goProjects/bookManageSystem/utils"
	"strconv"
	"time"
)

type User struct{
	UserId int `gorm:"primaryKey"` //用户ID
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
		utils.NewResponse(c).ToErrorResponse(400,"parse user message error")
		return
	}
	nowTime:=time.Now()
	user := User{ UserName: user_name, Password: password, CreateAt: nowTime,UserType: user_type }
	// 插入库表
	if err1:=(utils.DB).Create(&user).Error;err1!=nil{
		utils.NewResponse(c).ToErrorResponse(500,"create user error")
		return
	}
	// 生成token
	token,err := utils.GenerateToken(user_name,user.UserId,user.UserType)
	if err!=nil{
		utils.NewResponse(c).ToErrorResponse(500,"GenerateToken error")
		return
	}
	// 处理返回信息
	utils.NewResponse(c).ToResponse("sign user success",
		gin.H{"token":token, "user_id":user.UserId, "user_name":user.UserName,"user_type":user.UserType})
}

func(u *User)Login(c *gin.Context){
	// 获取请求参数
	userid, ok:=c.GetPostForm("user_id")
	user_id,err:= strconv.Atoi(userid)
	if err!=nil{
		utils.NewResponse(c).ToErrorResponse(400,"parse user message error")
		return
	}
	password, ok:=c.GetPostForm("password")
	if !ok{
		utils.NewResponse(c).ToErrorResponse(400,"parse user message error")
		return
	}
	// 查询user是否在库表存在
	user := User{UserId: user_id}
	if err1:=utils.DB.First(&user).Error;err1!=nil{
		utils.NewResponse(c).ToErrorResponse(400,"not found user")
		return
	}
	// 判断用户ID与密码是否匹配
	if password !=user.Password{
		utils.NewResponse(c).ToErrorResponse(400,"userId or password error")
		return
	}
	// 生成token
	token,err := utils.GenerateToken(user.UserName,user.UserId,user.UserType)
	if err!=nil{
		utils.NewResponse(c).ToErrorResponse(500,"GenerateToken error")
		return
	}
	// 处理返回数据
	utils.NewResponse(c).ToResponse("login success",
		gin.H{"token":token, "user_id":user.UserId, "user_name":user.UserName,"user_type":user.UserType})
}

func(u *User)Action(c *gin.Context){

}

func(u *User)Detail(c *gin.Context){

}

func NewUser()*User{
	return &User{}
}