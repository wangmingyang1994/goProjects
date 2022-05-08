package v1

import "github.com/gin-gonic/gin"

type BookManage struct{}

func NewBookManage()*BookManage{
	return &BookManage{}
}

func(u *BookManage)AddBook(c *gin.Context){

}

func(u *BookManage)EditBook(c *gin.Context){

}

func(u *BookManage)BorrowBook(c *gin.Context){

}

func(u *BookManage)ReturnBook(c *gin.Context){

}
func(u *BookManage)KindOfBooks(c *gin.Context){

}
func(u *BookManage)BooksDetail(c *gin.Context){

}
