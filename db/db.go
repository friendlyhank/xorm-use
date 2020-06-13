package db

import (
	mysql "github.com/go-sql-driver/mysql"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
)

var (
	// 的数据库
	dbEngine *xorm.Engine
)

// Init - 自动初始化数据库的链接
func Init() {

	var (
		err error
	)
	// 把这个依赖关系放到这里, 这样这里其实是初始化了驱动为Mysql
	if mysql.ErrInvalidConn != nil {
	}

	//
	dbSource := "root:123456@tcp(127.0.0.1:3306)/sakila?charset=utf8mb4&loc=Local&interpolateParams=true"
	if dbEngine, err = xorm.NewEngine("mysql", dbSource); err != nil {
		panic(err)
	}
	dbEngine.ShowSQL(true)
	dbEngine.ShowExecTime(true)
	dbEngine.Logger().SetLevel(core.LOG_INFO)
}

// Engine - 主数据库
func Engine() *xorm.Engine {
	return dbEngine
}
