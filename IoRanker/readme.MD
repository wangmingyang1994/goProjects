作业：\
1.使用冒泡排序、快速排序分别实现体脂排序功能\
2.注册时，保存所有注册信息到文件中，使用json格式保存，每行一条记录\
作业说明：\
1.rankOfFateRate.rank中为原排行榜程序\
2.rankOfFateRate.sorter文件中抽象出一个排序接口Sorter。
切换冒泡排序或快速排序只需在rank文件的第25/26行切换即可。\
如下：\
//sort := Bubble{}\
//sort := Quick{}\
3.writer.writer中实现了文件写入的方法，默认在上一层级的文件夹中生成initLog文件，将注册信息保存到此文件中\
4.快速测试：\
cd IoRanker\
go run main.go

