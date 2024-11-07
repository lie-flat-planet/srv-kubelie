package model

import (
	"github.com/lie-flat-planet/srv-kubelie/config"
)

func init() {
	config.Config.Mysql.AppendModel(&User{})
}

type User struct {
	ID
	Name  string
	Age   int
	Phone string
	TimeAt
}

func (User) TableName() string {
	return "user"
}