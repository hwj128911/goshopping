package models

import (
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thewinds/goshopping/server/com"
)

type User struct {
	Id          int64
	Username    string
	Password    string
	Phone       string
	Iconpath    string
	Email       string
	Bonuspoints int
}

//AddUser 根据用户结构添加用户
func AddUser(u User) error {
	if has, err := x.Where("phone=?", u.Phone).Get(new(User)); err == nil {
		if has {
			return com.NewErr("添加用户失败,用户已存在", nil)
		}
	} else {
		return com.NewErr("添加用户失败,查询失败", err)
	}

	u.Id = 0
	u.Bonuspoints = 10
	u.Username = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	_, err := x.Insert(u)
	if err != nil {
		return com.NewErr("添加用户失败,插入记录失败", err)
	}
	return nil
}

//DeleteUser 根据用户Id删除用户
func DeleteUser(uid int64) error {
	n, err := x.Id(uid).Delete(new(User))
	if err != nil {
		return com.NewErr("删除用户失败", err)
	}
	if n == 0 {
		return com.NewErr("删除用户失败,用户不存在", nil)
	}
	return nil
}

//GetUser 根据用户Id获取用户
func GetUser(uid int64) (u *User, err error) {
	user := new(User)
	ok, err := x.Id(uid).Get(user)
	if err != nil {
		return nil, com.NewErr("获取用户失败", err)
	}
	if ok == false {
		return nil, com.NewErr("用户不存在", err)
	}
	return user, nil
}

//GetAllUsers 获取所有用户
func GetAllUsers() []User {
	users := make([]User, 0)
	x.Find(&users)
	return users
}

//UpdateUser 更新用户
func UpdateUser(u User) error {
	user, err := GetUser(u.Id)
	if err != nil {
		return com.NewErr("更新用户失败", err)
	}
	if u.Bonuspoints != 0 {
		user.Bonuspoints = u.Bonuspoints
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Iconpath != "" {
		user.Iconpath = u.Iconpath
	}
	if u.Password != "" {
		user.Password = u.Password
	}
	if u.Phone != "" {
		user.Phone = u.Phone
	}
	if u.Username != "" {
		user.Username = u.Username
	}
	_, err = x.Id(u.Id).Update(user)
	if err != nil {
		return com.NewErr("更新数据失败", err)
	}
	return nil
}

type Deliveryadress struct {
	Id         int64
	Uid        int64
	Default    int //0非默认 1默认
	Provinceid int64
	Cityid     int64
	Districtid int64
	Adress     string
	Phone      string
}
