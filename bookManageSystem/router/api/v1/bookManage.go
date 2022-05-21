package v1

import (
	"github.com/gin-gonic/gin"
	"goProjects/bookManageSystem/utils"
	"strconv"
	"time"
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
	// 获取userID
	token:=c.Request.Header["Token"][0]
	parseToken, _ := utils.ParseToken(token)
	userId := parseToken.UserId
	// 获取借书参数
	bookId_str:=c.PostForm("book_id")
	bookId,_:=strconv.Atoi(bookId_str)
	bookName:=c.PostForm("book_name")
	day_str:=c.PostForm("day")
	day,_:=strconv.Atoi(day_str)
	now:= time.Now()
	// 初始化借书实例
	record:=UserBookRecord{
		UserId: userId,
		BookId: bookId,
		BookName: bookName,
		Days: day,
		StartDate: now,
	}
	//借书
	if utils.DB.Create(&record).Error !=nil{
		utils.NewResponse(c).ToErrorResponse(500,"borrow book error")
		return
	}
	//处理返回
	utils.NewResponse(c).ToResponse("borrow book success",
		gin.H{"records":record})
	return
}

func(b *BookManage)ReturnBook(c *gin.Context){
	// 获取还书参数
	recordId_str:=c.PostForm("record_id")
	recordId,_:=strconv.Atoi(recordId_str)
	// 初始化借书实例
	record:=UserBookRecord{
		RecordId:recordId,
	}
	//查询bookId
	tx:= utils.DB.Begin()
	tx.First(&record)
	//查询bookstock
	book:= Book{BookId: record.BookId}
	tx.First(&book)
	//还书
	if tx.Model(UserBookRecord{}).Where("book_id=?",recordId).Update("book_status",1).Error !=nil{
		tx.Rollback()
		utils.NewResponse(c).ToErrorResponse(500,"return book error")
		return
	}
	if utils.DB.Model(Book{}).Where("book_id=?",record.BookId).Update("book_stock",book.BookStock+1).Error!=nil{
		tx.Rollback()
		utils.NewResponse(c).ToErrorResponse(500,"return book error")
		return
	}
	tx.Commit()
	//处理返回
	utils.NewResponse(c).ToResponse("return book success",
		gin.H{"records":record})
	return
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

func (b *BookManage) BookKinds(c *gin.Context) {
	//获取所有书籍分类
	kinds := make([]string,0,10)
	if utils.DB.Model(Book{}).Distinct("book_type").Scan(&kinds).Error!=nil{
		utils.NewResponse(c).ToErrorResponse(500,"select book_kinds error")
		return
	}
	utils.NewResponse(c).ToResponse("select book_kinds success",
		gin.H{"book_kinds":kinds,"num":len(kinds)})
	return

}
