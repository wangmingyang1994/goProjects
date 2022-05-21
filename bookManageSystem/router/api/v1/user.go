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

type UserBookRecord struct{
	RecordId int `gorm:"primaryKey"` //记录ID
	UserId int //用户ID
	BookId int //书籍ID
	BookName string //书名
	StartDate time.Time// 借书申请时间
	Days int // 借书时长
	BookStatus string // 0未持有 1已归还 2未归还
	PassStatus string // 0审核中 1已审核 2已拒绝
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


func(u *User)Message(c *gin.Context){
	// 解析token，获取用户角色，非管理员获取审核列表
	token:=c.Request.Header["Token"][0]
	parseToken, _ := utils.ParseToken(token)
	if parseToken.UserType!="1"{
		utils.NewResponse(c).ToErrorResponse(400,"No permissions")
		return
	}
	records := make([]UserBookRecord,0,100)
	// 查询用户的书籍
	if utils.DB.Model(UserBookRecord{PassStatus: "0"}).Scan(&records).Error!= nil{
		utils.NewResponse(c).ToErrorResponse(500,"select user books error")
		return
	}
	// 处理返回数据
	utils.NewResponse(c).ToResponse("select user books success",
		gin.H{"num_list":len(records), "lists":records})

}

func(u *User)Action(c *gin.Context){
	// 解析token，获取用户角色，非管理员无法审核
	token:=c.Request.Header["Token"][0]
	parseToken, _ := utils.ParseToken(token)
	if parseToken.UserType!="1"{
		utils.NewResponse(c).ToErrorResponse(400,"No permissions")
		return
	}
	do, ok:=c.GetPostForm("do")
	record_id1, ok:=c.GetPostForm("record_id")
	record_id, _:=strconv.Atoi(record_id1)
	//0不通过 1通过
	if !ok{
		utils.NewResponse(c).ToErrorResponse(400,"parse user message error")
		return
	}
	// 处理审核
	switch do{
	case "0":
		//更新审核状态
		utils.DB.Model(UserBookRecord{}).Where("record_id=?",record_id).Update("pass_status",2)
		utils.NewResponse(c).ToResponse("已标记为失败",gin.H{})
		return
	case "1":
		tx := utils.DB.Begin()
		//更新审核状态，发放书籍
		if tx.Model(UserBookRecord{}).Where("record_id=?",record_id).Update("pass_status",1).Error!=nil{
			tx.Rollback()
			utils.NewResponse(c).ToResponse("操作失败",gin.H{})
			return
		}
		if tx.Model(UserBookRecord{}).Where("record_id=?",record_id).Update("book_status",2).Error!=nil{
			tx.Rollback()
			utils.NewResponse(c).ToResponse("操作失败",gin.H{})
			return
		}
		//书架减去库存
		record:=UserBookRecord{RecordId: record_id}
		tx.First(&record)
		bookId:=record.BookId
		book:=Book{BookId:bookId }
		tx.First(&book)
		if tx.Model(Book{}).Where("book_id=?",bookId).Update("book_stock",book.BookStock-1).Error!=nil{
			tx.Rollback()
			utils.NewResponse(c).ToResponse("操作失败",gin.H{})
			return
		}
		tx.Commit()
		utils.NewResponse(c).ToResponse("已审核通过",gin.H{})
		return
	}
}

func(u *User)Books(c *gin.Context){
	// 解析token，获取用户UserId
	token:=c.Request.Header["Token"][0]
	parseToken, _ := utils.ParseToken(token)
	userId := parseToken.UserId
	records:= make([]UserBookRecord,0,100)
	// 查询用户的书籍
	if utils.DB.Model(UserBookRecord{UserId: userId}).Scan(&records).Error!= nil{
		utils.NewResponse(c).ToErrorResponse(500,"select user books error")
		return
	}
	// 处理返回数据
	utils.NewResponse(c).ToResponse("select user books success",
		gin.H{"userId":userId, "books":records})
}

func NewUser()*User{
	return &User{}
}