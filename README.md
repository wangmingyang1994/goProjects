作业：\
1.使用github上的lib：github.com/armstrongli/go-bmi 完成体脂计算器\
本地添加module的replace，并在本地项目扩展 github.com/armstrongli/go-bmi以支持
BMI、FatRate的计算\
使用 vendor 保证代码的完整性与可运行\
2.1）为体脂计算器编写单元测试并完善体脂计算器的验证逻辑\
• BMI计算\
录入正常身高、体重，确保计算结果符合预期\
录入0或负数身高，返回错误\
录入0或负数体重，返回错误\
2）为体脂计算器编写单元测试并完善体脂计算器的验证逻辑\
• 体脂率计算\
录入正常BMI、年龄、性别，确保计算结果符合预期\
录入非法BMI、年龄、性别(0、负数、超过150的年龄、非男女的性别输入)，返回错误\
录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期\

说明：\
作业1：lib库：github.com/armstrongli/go-bmi 中go-bmi.go文件中体脂计算功能缺失，通过使用了replace功能，将本地目录staging/src/github.com/armstrongli/go-bmi 改写，将体脂计算功能完善，并在本工程中（goProjects/FatRateCalc/calcFatRate.go）优化了输入方式，使用cobra包的功能后进行调用\
体脂率功能验证命令：\
执行目录： 主目录 => goProjects\
执行命令： go run main.go --name dandy --sex 男 --tall 1.8 --weight 60 --age 20\
作业2：因体脂计算功能在staging/src/github.com/armstrongli/go-bmi 中实现，便将验证逻辑写在了同一目录下\
其中：\
TestGetBMI方法验证BMI计算逻辑\
TestGetFatRate方法验证体脂的计算逻辑\
TestGetSuggestion方法建议逻辑\
测试函数中使用表格驱动测试的方法，将输入条件及预期结果封装到struct中\
体脂计算测试文件验证命令：\
执行目录：staging/src/github.com/armstrongli/go-bmi\
执行命令：\
cd staging/src/github.com/armstrongli/go-bmi\
go test . 或 go test -coverprofile=c.out

