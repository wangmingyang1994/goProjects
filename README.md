作业：\
1.使用github的公有函数来实现功能\
2.自己重写BMI计算器、体脂计算器，用module replace替换github上的\
3.使用vendor为项目依赖提供保障

说明：\
1.使用github的公有函数实现录入，在goModTry/completeFatRate/newFateRate/newFateRate.go中，使用cobra命令行工具实现录入功能\
2.在本地（fatRateCalc/fatRate.go）中重写了BMI计算器和体脂计算器\
3.在go.mod(module goProjects/goModTry)使用了replace功能引入了其他go.mod(module goProjects/FatRateCalc)中的功能\
4.在goModTry/completeFatRate/main.go中调用程序，实现体脂计算并输出到控制台\
执行目录： cd goModTry/completeFatRate\
执行命令： go run main.go --name dandy --sex 男 --tall 1.8 --weight 60 --age 20\
5.使用了vendor命令在module goProjects/goModTry中生成了vender目录
