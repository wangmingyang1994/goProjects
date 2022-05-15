package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB
func NewDB() (*gorm.DB,error){
	dsn:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",DatabaseMsg.UserName,
		DatabaseMsg.Password,DatabaseMsg.Host,DatabaseMsg.DBName)
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil,err
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(DatabaseMsg.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(DatabaseMsg.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second*60)
	return db,nil
}

func SetupDB()error{
	var err error
	DB,err = NewDB()
	if err!= nil{
		return err
	}
	return nil
}


	// 连接表
	//dsn := "root:11111111@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//create table
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	// log.Println(err)
	// return
	//}

	//insert
	//user := User{Id: 9, Name: "Jinzhu", Age: 18}
	//db.Create(&user) // 通过数据的指针来创建

	//for i := 0; i < 9; i++ {
	// user := User{i, fmt.Sprintf("name_%d", i), int32(i + 10)}
	// db.Create(&user)
	//}

	//select
	//user := User{Id: 1}
	//db.First(&user)
	//fmt.Println(user)

	//result := map[string]interface{}{}
	//db.Model(&User{}).First(&result)
	//fmt.Println(result)

	//条件查询
	//users := []User{}
	//db.Find(&users, []int{1, 2, 3})
	//fmt.Println(users)
	//db.Where("id IN ?", []int{1, 2}).Find(&users)
	//fmt.Println(users)

	//关联查询
	//type p1 struct {
	// Name    string
	// Age     int
	// SubjectName string
	// Score    int
	//}
	//var result []p1
	//db.Table("users").Select("users.name, users.age, user_scoree.subject_name, user_scoree.score").Joins("left join user_scoree on users.id = user_scoree.id").Scan(&result)
	//fmt.Println(result)

	//update
	//db.Model(&User{}).Where("id = ?", 1).Update("name", "hello")
	//db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "hello1", Age: 19})

	//delete
	//db.Where("id = ?", "8").Delete(&User{})

	// 开始事务
	//tx := db.Begin()
	//tx.Create()
	//tx.Rollback()
	//tx.Commit()
