package db

import (
	"time"
)

type User struct {
	ID         int32
	Name       string
	Passwd     string
	Salt       string
	Desc       string
	Auth       int64 //权限
	CreateTime time.Time
	UpdateTime time.Time
}

type DB interface {
	Init() error
	Destroy()
}

type UserDB interface {
	DB
	QueryUserByID(id int32) (User, error)
	QueryUserByName(name string) (User, error)
	InsertUser(user *User) (int32, error)
	UpdateUser(user *User) error
	UpdateUserName(user *User) error
	UpdateUserPasswd(user *User) error
	UpdateUserDesc(user *User) error
	UpdateUserAuth(user *User) error
}
