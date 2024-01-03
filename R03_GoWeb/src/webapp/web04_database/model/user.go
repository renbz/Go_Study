package model

import (
	"R03_GoWeb/src/webapp/web04_database/util"
	"fmt"
)

// User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// AddUser 添加user的方法
func (user *User) AddUser() error {
	// 1. 编写sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 2. 预编译
	inStmt, err := util.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("预编译出现异常")

	}
	// 3. 执行
	_, err2 := inStmt.Exec("admin", "123456", "admin@163.com")
	if err2 != nil {
		fmt.Println("执行出现异常:", err2)
		return err2
	}
	return nil
}

// AddUser 添加user的方法
func (user *User) AddUser2() error {
	// 1. 编写sql
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	// 3. 执行
	_, err2 := util.Db.Exec(sqlStr, "admin2", "123456", "admin2@163.com")
	if err2 != nil {
		fmt.Println("执行出现异常:", err2)
		return err2
	}
	return nil
}

// 获取一条数据库记录
func (user *User) GetUserByID() (*User, error) {
	// 写sql语句
	sqlStr := "select id,username,password,email from users where id = ?"
	// 执行
	row := util.Db.QueryRow(sqlStr, user.ID)
	// 声明
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

// 获取所有用户记录
func (user *User) GetUserAll() ([]*User, error) {
	// 写sql语句
	sqlStr := "select id,username,password,email from users "
	// 执行
	rows, err := util.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		// 声明
		var id int
		var username string
		var password string
		var email string
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, nil
}
