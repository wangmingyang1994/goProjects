package v1

import (
	"github.com/gin-gonic/gin"
	"goProjects/bookManageSystem/utils"
	"strconv"
)

type Book struct{
	BookId int `gorm:"primaryKey"`
	BookName string
	BookType string
	BookAuthor string
	BookStock int
}
type BookManage struct{}

func NewBookManage()*BookManage{
	return &BookManage{}
}

func(b *BookManage)AddBook(c *gin.Context){
	// 解析token，获取用户角色，非管理员无法新增书籍
	token:=c.Request.Header["Token"][0]
	parseToken, _ := utils.ParseToken(token)
	if parseToken.UserType!="1"{
		utils.NewResponse(c).ToErrorResponse(400,"No permissions")
		return
	}
	// 获取新增书籍的参数，并插入库表
	bookName:=c.PostForm("book_name")
	bookType:=c.PostForm("book_type")
	bookAuthor:=c.PostForm("book_author")
	bookStocks:=c.PostForm("book_stock")
	bookStock,_:=strconv.Atoi(bookStocks)
	book:=Book{
		BookName: bookName,
		BookAuthor: bookAuthor,
		BookType: bookType,
		BookStock: bookStock}
	if utils.DB.Create(&book).Error != nil{
		utils.NewResponse(c).ToErrorResponse(500,"add book error")
		return
	}
	utils.NewResponse(c).ToResponse("add book success",
		gin.H{"bookId":book.BookId,"bookName":bookName,"bookType":bookType,"bookAuthor":bookAuthor,"bookStock":bookStock})
	return
}

func(b *BookManage)EditBook(c *gin.Context){
	// 解析token，获取用户角色，非管理员无法编辑书籍
	token:=c.Request.Header["Token"][0]
	parseToken, _ := utils.ParseToken(token)
	if parseToken.UserType!="1"{
		utils.NewResponse(c).ToErrorResponse(400,"No permissions")
		return
	}
	// 获取新增书籍的参数，并插入库表
	bookid:=c.PostForm("book_id")
	bookName:=c.PostForm("book_name")
	bookType:=c.PostForm("book_type")
	bookAuthor:=c.PostForm("book_author")
	bookStocks:=c.PostForm("book_stock")
	bookId,_:=strconv.Atoi(bookid)
	bookStock,_:=strconv.Atoi(bookStocks)
	book:= Book{
		BookId: bookId,
		BookName: bookName,
		BookAuthor: bookAuthor,
		BookType: bookType,
		BookStock: bookStock}
	if utils.DB.Model(Book{}).Where("book_id=?",bookId).Updates(&book).Error != nil{
		utils.NewResponse(c).ToErrorResponse(500,"add book error")
		return
	}
	utils.NewResponse(c).ToResponse("update book success",
		gin.H{"bookName":bookName,"bookType":bookType,"bookAuthor":bookAuthor,"bookStock":bookStock})
	return
}

func(b *BookManage)BorrowBook(c *gin.Context){

}

func(b *BookManage)ReturnBook(c *gin.Context){

}
func(b *BookManage)KindOfBooks(c *gin.Context){
	bookType:=c.Query("book_type")
	booklist:= []Book{}
	if utils.DB.Where("book_type=?",bookType).Find(&booklist).Error != nil{
		utils.NewResponse(c).ToErrorResponse(500,"select book error")
		return
	}
	// 返回书籍详情
	utils.NewResponse(c).ToResponse("select book success",
		gin.H{"books":booklist})
	return
}
func(b *BookManage)BooksDetail(c *gin.Context){
	// 获取请求参数，根据bookId查询库表
	bookid:=c.Query("book_id")
	bookId,_:=strconv.Atoi(bookid)
	book:=&Book{BookId: bookId}
	if utils.DB.First(&book).Error != nil{
		utils.NewResponse(c).ToErrorResponse(500,"select book error")
		return
	}
	// 返回书籍详情
	utils.NewResponse(c).ToResponse("select book success",
		gin.H{"bookId":book.BookId,"bookName":book.BookName,
			"bookType":book.BookType,"bookAuthor":book.BookType,"bookStock":book.BookStock})
	return
}
