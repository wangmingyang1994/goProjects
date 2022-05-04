package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goProjects/blog_service/pkg/setting"
	"time"
)

type Model struct{
	ID uint32 `grom:"primary_key" json:"id"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings)(*gorm.DB,error){
	dsn:=fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,)
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		return nil,err
	}
	//if global.ServerSetting.RunMode == "debug"{
	//	db.LogMode(true)
	//}
	//db.SingularTable(true)
	//db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	//db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db,nil

}

func CreateCallback(db *gorm.DB)*gorm.DB{
	nowTime:=time.Now().Unix()
	return db.Update("CreatedOn", nowTime)
}

func UpdateCallback(db *gorm.DB)*gorm.DB{
	nowTime:=time.Now().Unix()
	return db.Update("ModifiedOn", nowTime)
}

func DeleteCallback(db *gorm.DB)*gorm.DB{
	nowTime:=time.Now().Unix()
	return db.Update("DeletedOn", nowTime)

}

func addExtraSpaceIfExist(str string) string{
	if str!=""{
		return " "+str
	}
	return ""
}