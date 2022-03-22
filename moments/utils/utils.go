package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "11111111"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "test"
)


func SqlConn() *sql.DB {
	// 连接数据库后返回DB
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open mysql failed,err:", err)
	}
	DB.SetConnMaxLifetime(10 * time.Second)
	return DB
}

func GetBMI(tall, weight float64) (Bmi float64, err error) {
	//计算BMI
	if weight <= 0 {
		return 0, fmt.Errorf("体重输入有误:%f\n", weight)
	}
	if tall <= 0 || tall > 3 {
		return 0, fmt.Errorf("身高输入有误:%d\n", tall)
	}
	Bmi = weight / (tall * tall)
	return Bmi, nil
}

func GetFatRate(bmi float64, age int, sex string) (FatRate float64, err error) {
	// 计算体脂率
	if bmi <= 0 {
		err = fmt.Errorf("bmi输入有误:%f\n", bmi)
		return 0, err
	}
	if age <= 0 || age >= 150 {
		err = fmt.Errorf("年龄输入有误:%d\n", age)
		return 0, err
	}
	switch sex {
	case "男":
		FatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*1) / 100
		return FatRate, nil
	case "女":
		FatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*0) / 100
		return FatRate, nil
	default:
		err = fmt.Errorf("性别输入有误:%s\n", sex)
		return 0, err
	}

}
