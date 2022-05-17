package router

import (
	"github.com/gin-gonic/gin"
	"goProjects/bookManageSystem/router/api/v1"
	"goProjects/bookManageSystem/utils"
	"net/http"
)


func NewRouter () (*gin.Engine){
	// 创建gin实例
	r:= gin.New()
	// 使用中间件
	r.Use(utils.Authvalidate(),utils.GinLogger(),utils.GinRecovery(true))
	// 测试路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello~")
	})
	//创建路由组，注册路由
	user := v1.NewUser()
	userServer := r.Group("book/user")
	{
		userServer.POST("/sign", user.Sign)
		userServer.POST("/login", user.Login)
		userServer.POST("/action", user.Action)
		userServer.GET("/detail/:userId", user.Detail)

	}
	bookManage := v1.NewBookManage()
	bookManageServer := r.Group("book/bookManage")
	{
		bookManageServer.POST("/addBook", bookManage.AddBook)
		bookManageServer.PUT("/editBook", bookManage.EditBook)
		bookManageServer.POST("/borrowBook", bookManage.BorrowBook)
		bookManageServer.POST("/returnBook", bookManage.ReturnBook)
		bookManageServer.POST("/kindOfBooks", bookManage.KindOfBooks)
		bookManageServer.POST("/booksDetail/:bookId", bookManage.BooksDetail)

	}
	return r
}
