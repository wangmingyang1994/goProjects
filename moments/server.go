package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goProjects/moments/states"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello moments!")
	})
	r.POST("/initPerson", InitUser)
	r.POST("/addMoments", AddMoment)
	r.POST("/deleteMoments", DeleteMoments)
	r.GET("/getAllMoments", GetAllMoments)
	r.GET("/getMyMoments", GetMyMoments)
	r.GET("/getMoment", GetMoment)

	r.Run(":8081")
}



func InitUser(c *gin.Context) {
	//初始化用户，绑定到obj
	person := states.Person{}
	if c.ShouldBind(&person) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    "用户注册信息有误!",
		})
		return
	}
	//注册用户，获取用户ID
	personId, err := person.NewPerson()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"msg":    "注册失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"msg":      "注册成功!",
			"personId": personId,
		})
	}
}


func AddMoment(c *gin.Context) {
	//解析personId和动态内容
	pId,err:= c.GetPostForm("personId")
	personId,err1:= strconv.Atoi(pId)
	content,err2:= c.GetPostForm("content")
	if !err || err1!=nil || !err2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    "动态或用户ID输入有误!",
		})
		return
	}
	//插入动态到数据库
	statesId, err3 := states.NewStates(personId, content)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"msg":    "发送动态失败！",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"msg":      "动态发送成功!",
			"statesId": statesId,
		})
	}
}

func DeleteMoments(c *gin.Context) {
	personId,b := c.GetPostForm("personId")
	contentId,b1 := c.GetPostForm("statesId")
	if !b || !b1{
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"msg":    "用户ID或动态ID输入有误！",
		})
		return
	}
	fmt.Println(personId,contentId)
	if err:=states.DeleteStates(personId, contentId);err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"msg":    "删除动态失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"msg":    "动态删除成功!",
		})
	}
}

func GetAllMoments(c *gin.Context) {
	var moments []states.States
	moments, err := states.GetAllStates()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"msg":    "获取动态失败！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   moments,
	})
}

func GetMyMoments(c *gin.Context) {
	personId, _ := strconv.Atoi(c.Query("personId"))
	var moments []states.States
	moments, err := states.GetMyStates(personId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"msg":    "获取用户动态失败！",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "success",
			"msg":    "获取用户状态成功！",
			"data":   moments,
		})
	}

}

func GetMoment(c *gin.Context) {
	momentId, _ := strconv.Atoi(c.Query("statesId"))
	var moments []states.States
	moments, err := states.GetStates(momentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"msg":    "获取动态失败！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"msg":    "获取动态成功！",
			"data":   moments,
		})
	}
}
