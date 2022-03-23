package states

import (
	"fmt"
	"goProjects/moments/utils"
)

type Person struct {
	PersonId int     `form:"personId"`
	Name     string  `form:"name"`
	Sex      string  `form:"sex"`
	Age      int     `form:"age"`
	Tall     float64 `form:"tall"`
	Weight   float64 `form:"weight"`
	FatRate  float64 `form:"fatRate"`
}

func (person *Person) NewPerson() (int, error) {
	// 初始化用户到数据库
	// 计算出体脂率
	bmi, err := utils.GetBMI(person.Tall, person.Weight)
	if err != nil {
		return 0, err
	}
	fatRate, err := utils.GetFatRate(bmi, person.Age, person.Sex)
	if err != nil {
		return 0, err
	}
	// 连接数据库
	DB := utils.SqlConn()
	defer DB.Close()
	// 数据库插入用户
	_, errs := DB.Exec("insert into persons(name, sex, age,tall, weight, fatRate) values(?,?,?,?,?,?);",
		person.Name, person.Sex, person.Age, person.Tall, person.Weight, fatRate)
	if errs != nil {
		e := fmt.Errorf("插入用户失败:%v\n", errs)
		return 0, e
	}
	// 获取用户ID，并返回
	var personId int
	row1 := DB.QueryRow("select personId from persons where name=?", person.Name)
	if err := row1.Scan(&personId); err != nil {
		e := fmt.Errorf("写入peronId失败:%v\n", err)
		return 0, e
	}
	return personId, nil
}

type States struct {
	StatesId      int     `form:"statesId" json:"statesId"`
	Content       string  `form:"content" json:"content"`
	CreateTime    string  `form:"createTime" json:"createTime"`
	AuthorId      int     `form:"authorId" json:"authorId"`
	AuthorName    string  `form:"authorName" json:"authorName"`
	AuthorAge     int     `form:"authorAge" json:"authorAge"`
	AuthorTall    float64 `form:"authorTall" json:"authorTall"`
	AuthorWeight  float64 `form:"authorWeight" json:"authorWeight"`
	AuthorFatRate float64 `form:"authorFatRate" json:"authorFatRate"`
}

func NewStates(personId int, content string) (int, error) {
	// 连接数据库，插入消息到数据库
	DB := utils.SqlConn()
	defer DB.Close()
	//插入动态
	_, errs := DB.Exec("insert into test.states(personId, content) values(?,?);", personId, content)
	if errs != nil {
		e := fmt.Errorf("插入动态失败:%v\n", errs)
		return 0, e
	}
	//获取statesId，并返回
	var statesId int
	row1 := DB.QueryRow("select statesId from states where personId=? and content=?", personId, content)
	if err := row1.Scan(&statesId); err != nil {
		e := fmt.Errorf("写入statesId失败:%v\n", err)
		return 0, e
	}
	return statesId, nil
}

func DeleteStates(personId string, contentId string) error {
	// 删除用户，根据用户ID和消息ID，将用户可见置为否
	DB := utils.SqlConn()
	defer DB.Close()
	_, errs := DB.Exec("update test.states set visable=false where statesId=? and personId=?",
		contentId, personId)
	if errs != nil {
		e := fmt.Errorf("删除动态失败:%v\n", errs)
		return e
	}
	return nil
}

func GetAllStates() ([]States, error) {
	//连接数据库，获取所有的动态
	DB := utils.SqlConn()
	defer DB.Close()
	rows, err := DB.Query("select b.statesId,b.content,b.createTime,a.personId,a.name,a.age," +
		"a.tall,a.weight,a.fatRate from test.persons a join test.states b on a.personId=b.personId " +
		"where b.visable=1 order by b.createTime desc;")
	if err != nil {
		return nil, err
	}
	// 初始化动态结构体slice，将查询结果集写入slice，并返回
	moments := make([]States, 0, 1000)
	for rows.Next() {
		moment := States{}
		err := rows.Scan(&moment.StatesId, &moment.Content, &moment.CreateTime,
			&moment.AuthorId, &moment.AuthorName, &moment.AuthorAge, &moment.AuthorTall,
			&moment.AuthorWeight, &moment.AuthorFatRate)
		if err != nil {
			return nil, err
		}
		moments = append(moments, moment)
	}
	return moments, nil
}

func GetMyStates(personId int) ([]States, error) {
	//连接数据库，获取某个用户所有的动态
	DB := utils.SqlConn()
	defer DB.Close()
	rows, err := DB.Query("select b.statesId,b.content,b.createTime,a.personId,a.name,a.age,"+
		"a.tall,a.weight,a.fatRate from test.persons a join test.states b on a.personId=b.personId where b.visable=1 and a.personId=? order by b.createTime desc;", personId)
	if err != nil {
		return nil, err
	}
	// 初始化动态结构体slice，将查询结果集写入slice，并返回
	moments := make([]States, 0, 1000)
	for rows.Next() {
		moment := States{}
		err := rows.Scan(&moment.StatesId, &moment.Content, &moment.CreateTime,
			&moment.AuthorId, &moment.AuthorName, &moment.AuthorAge, &moment.AuthorTall,
			&moment.AuthorWeight, &moment.AuthorFatRate)
		if err != nil {
			return nil, err
		}
		moments = append(moments, moment)
	}
	return moments, nil
}

func GetStates(statesId int) ([]States, error) {
	//连接数据库，根据动态ID获取某条动态
	DB := utils.SqlConn()
	defer DB.Close()
	rows, err := DB.Query("select b.statesId,b.content,b.createTime,a.personId,a.name,a.age,"+
		"a.tall,a.weight,a.fatRate from test.persons a join test.states b "+
		"on a.personId=b.personId where b.visable=1 and b.statesId=? order by b.createTime desc;", statesId)
	if err != nil {
		return nil, err
	}
	// 初始化动态结构体slice，将查询结果集写入slice，并返回
	moments := make([]States, 0, 1000)
	for rows.Next() {
		moment := States{}
		err := rows.Scan(&moment.StatesId, &moment.Content, &moment.CreateTime,
			&moment.AuthorId, &moment.AuthorName, &moment.AuthorAge, &moment.AuthorTall,
			&moment.AuthorWeight, &moment.AuthorFatRate)
		if err != nil {
			return nil, err
		}
		moments = append(moments, moment)
	}
	return moments, nil
}
