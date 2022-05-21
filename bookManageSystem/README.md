项目说明：

1.目录介绍\
config/config.yaml：配置文件\
docs/db.sql：建表语句记录\
logs：存放日志文件\
router/api/v1/bookManage.go：图书管理实现\
router/api/v1/user.go：用户服务实现\
router/router.go：注册路由组\
utils/auth.go：中间件-token验证\
utils/db.go：gorm初始化\
utils/jwt.go：token生成与解析\
utils/logger.go：日志初始化\
utils/response.go：处理接口返回信息\
utils/setting.go：初始化配置\
main.go：程序入口，服务注册与退出\
2.使用说明\
流程：\
管理员：注册->登录->新增书籍/编辑书籍->处理审核（管理员也可借书）\
普通用户：注册->登录->借书->还书\
接口：\
POST("book/user/sign", user.Sign)：注册用户（管理员/普通用户）\
POST("book/user/login", user.Login)：登录用户\
GET("book/user/message", user.Message)：管理员获取待审核列表\
POST("book/user/action", user.Action)：管理员处理审核\
POST("book/bookManage/addBook", bookManage.AddBook)：管理员添加书籍\
PUT("book/bookManage/editBook", bookManage.EditBook)：管理员编辑书籍\
GET("book/user/books", user.Books)：普通用户获取用户自身的借书列表及明细\
GET("book/bookManage/bookKinds", bookManage.BookKinds)：获取所有书籍分类\
GET("book/bookManage/kindOfBooks", bookManage.KindOfBooks)：获取某一类书籍的列表\
GET("book/bookManage/booksDetail", bookManage.BooksDetail)：获取某一本书的明细\
POST("book/bookManage/borrowBook", bookManage.BorrowBook)：普通用户借书\
POST("book/bookManage/returnBook", bookManage.ReturnBook)：普通用户还书\
3.服务启动：\
go run main.go 或 \
go build \
./bookManageSystem\
4.请求示例：\
注册用户：\
curl --location --request POST 'http://127.0.0.1:8088/book/user/sign' \
--form 'user_name="min111"' \
--form 'password="11222"' \
--form 'user_type="0"'\
登录用户：\
curl --location --request POST 'http://127.0.0.1:8088/book/user/login' \
--form 'user_id="7"' \
--form 'password="11222"'\
获取待审核列表：\
curl --location --request GET 'http://127.0.0.1:8088/book/user/message' \
处理审核消息：\
curl --location --request POST 'http://127.0.0.1:8088/book/user/action' \
--form 'book_name="笑话大全"' \
--form 'do="1"' \
--form 'record_id="1"'  \
添加书籍：\
curl --location --request POST 'http://127.0.0.1:8088/book/bookManage/addBook' \
--form 'book_name="笑话大全1"' \
--form 'book_type="娱乐类"' \
--form 'book_author="小朋友31"' \
--form 'book_stock="1"'\
编辑书籍：\
curl --location --request PUT 'http://127.0.0.1:8088/book/bookManage/editBook' \
--form 'book_name="笑话大全"' \
--form 'book_type="娱乐类"' \
--form 'book_author="小朋友2"' \
--form 'book_stock="3"' \
--form 'book_id="4"'\
获取用户借书列表：\
curl --location --request GET 'http://127.0.0.1:8088/book/user/books' \
获取所有书籍分类：\
curl --location --request GET 'http://127.0.0.1:8088/book/bookManage/bookKinds' \
获取某一类书籍列表：\
curl --location --request GET 'http://127.0.0.1:8088/book/bookManage/kindOfBooks?book_type="xxx"' \
获取某一本书明细：\
curl --location --request GET 'http://127.0.0.1:8088/book/bookManage/booksDetail'?book_id="xxx" \
借书：\
curl --location --request POST 'http://127.0.0.1:8088/book/bookManage/borrowBook' \
--form 'book_id="1"' \
--form 'book_name="golang开发"' \
--form 'day="10"'\
还书：\
curl --location --request POST 'http://127.0.0.1:8088/book/bookManage/returnBook' \
--form 'book_name="笑话大全"' \
--form 'record_id="1"'
