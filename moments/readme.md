模块介绍\
1.utils 提供公共方法（SQL连接，计算BMI及体脂率）\
2.states 方法实现\
3.server 服务实现\
4.sql 建表语句

功能介绍\
1.注册用户\
2.发送动态\
3.删除动态\
4.获取所有动态(feed中所有动态)\
5.获取单一用户的所有动态(个人主页所有动态)\
6.获取一条指定动态(进入动态详情)

接口调用\
1.注册用户接口：使用raw json传递参数\
curl -X POST 'http://127.0.0.1:8081/initPerson' -H 'Content-Type: application/json' -d '{"name":"xiaoying","sex":"女","age":22,"tall":1.66,"weight":50}'\
2.发布动态接口：使用form/data传递参数\
curl -X POST  "http://127.0.0.1:8081/addMoments" -d "personId=1&content=i am xiaoying"\
3.删除动态接口：使用form/data传递参数\
curl -X POST  "http://127.0.0.1:8081/deleteMoments" -d "personId=1&statesId=1"\
4.获取所有动态接口\
curl -X GET  "http://127.0.0.1:8081/getAllMoments" \
5.获取单一用户的所有动态接口：使用params传递参数\
curl -X GET  "http://127.0.0.1:8081/getMyMoments?personId=1" \
6.获取一条指定动态接口：使用params传递参数\
curl -X GET  "http://127.0.0.1:8081/getMoment?statesId=4"