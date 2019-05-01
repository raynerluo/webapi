package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DB *xorm.Engine
)

//初始化数据库
func init() {
	var err error
	DB, err = xorm.NewEngine("mysql", "root:haha54xxx@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
}
