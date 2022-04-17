package user

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"math/rand"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "11111111"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "mydb"
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

type User struct{}

func (User) SignUp(ctx context.Context, r *UserRegister) (*UserDetail, error) {
	//注册服务
	rand.Seed(time.Now().Unix())
	//生成userId
	userId := rand.Int31n(999999)
	//连接数据库
	db := SqlConn()
	defer db.Close()
	var num int
	//查询用户ID是否存在
	row := db.QueryRow("select count(*) from mydb.User where id=?", userId)
	if err := row.Scan(&num); err != nil || err != io.EOF {
		e := fmt.Errorf("signup读取数据库失败:%v\n", err.Error())
		return &UserDetail{}, e
	}
	//若该Id已存在，需要重试
	if num > 0 {
		e := fmt.Errorf("账号已存在，请登录!\n")
		return &UserDetail{}, e
	}
	//插入用户表
	_, errs := db.Exec("insert into mydb.User(id,name,password) values (?,?,?)",
		userId, r.UserName, r.Password)
	if errs != nil {
		e := fmt.Errorf("注册用户失败:%v\n", errs.Error())
		return &UserDetail{}, e
	}
	u1 := UserDetail{}
	u1.UserId = userId
	u1.UserName = r.UserName
	u1.Status = 0

	return &u1, nil
}

func (User) SignIn(ctx context.Context, r *UserLogin) (*UserDetail, error) {
	//连接库表
	db := SqlConn()
	defer db.Close()
	var num int
	//查询用户ID是否存在
	row := db.QueryRow("select count(*) from mydb.User where id=?", r.UserId)
	if err := row.Scan(&num); err != nil {
		e := fmt.Errorf("signin,读取数据库失败:%v\n", err.Error())
		return &UserDetail{}, e
	}
	//若不存在，需要去注册
	if num == 0 {
		e := fmt.Errorf("账号不存在，请注册!\n")
		return &UserDetail{}, e
	}
	//若存在，判断用户ID和密码是否匹配，若匹配则登录成功
	var name, password string
	row1 := db.QueryRow("select name,password from mydb.User where id=?", r.UserId)
	if err := row1.Scan(&name, &password); err != nil {
		e := fmt.Errorf("signin2,数据库读取失败:%v\n", err.Error())
		return &UserDetail{}, e
	}
	if password != r.Password {
		e := fmt.Errorf("账号或密码错误!\n")
		return &UserDetail{}, e
	}
	u1 := UserDetail{}
	u1.UserId = r.UserId
	u1.UserName = name
	u1.Status = 1
	_, errs := db.Exec("update mydb.User set status=1 where id=?",
		r.UserId)
	if errs != nil {
		e := fmt.Errorf("更新登录信息失败:%v\n", errs.Error())
		return &UserDetail{}, e
	}
	return &u1, nil
}

func (User) Logout(ctx context.Context, r *UserLogout) (*UserDetail, error) {
	db := SqlConn()
	defer db.Close()
	//将用户状态置为退出状态
	_, errs := db.Exec("update mydb.User set status=0 where id=?",
		r.UserId)
	if errs != nil {
		e := fmt.Errorf("更新登录信息失败:%v\n", errs.Error())
		return &UserDetail{}, e
	}
	u1 := UserDetail{UserId: r.UserId, Status: 0}
	return &u1, nil
}

func (User) GetUserInfos(ctx context.Context, n *Needusers) (*UserInfos, error) {
	db := SqlConn()
	defer db.Close()
	var count int32
	//查询所有用户总数
	row := db.QueryRow("select count(*) from mydb.User ")
	if err := row.Scan(&count); err != nil {
		e := fmt.Errorf("getuser,数据库读取失败:%v\n", err.Error())
		return &UserInfos{}, e
	}
	//根据limit，分页取用户，并返回
	rows, err := db.Query("select id,name,status from mydb.User where status=1 limit ?,?;", (*n).Page*(*n).Pagenums, (*n).Pagenums)
	if err != nil {
		return nil, err
	}
	userLists := make([]*UserDetail, 0, (*n).Pagenums)
	for rows.Next() {
		user := &UserDetail{}
		err := rows.Scan(&user.UserId, &user.UserName, &user.Status)
		if err != nil {
			return nil, err
		}
		userLists = append(userLists, user)
	}
	u := &UserInfos{
		Total: count,
		Users: userLists,
	}
	return u, err
}
