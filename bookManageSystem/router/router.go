package router

import (
	"github.com/gin-gonic/gin"
	v1 "goProjects/bookManageSystem/router/api/v1"
	"net/http"
)


func NewRouter () (*gin.Engine){
	r:= gin.New()
	r.Use(gin.Logger(),gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"hello~")
	})
	user:= v1.NewUser()
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
