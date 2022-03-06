package writer

import (
	"log"
	"os"
)


func Write_to_file(b []byte) {
	//打开文件
	path,err:=os.Getwd()
	if err!=nil{
		log.Fatalln("获取文件目录失败：", err)
	}
	file, err := os.OpenFile(path+"/initLog", os.O_CREATE| os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		log.Fatalln("打开文件失败：", err)
	}
	_, err = file.Write(b)
	if err != nil {
		log.Fatalln("写入文件失败：", err)
	}
}
