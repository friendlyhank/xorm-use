package db

import (
	"git.biezao.com/ant/xmiss/foundation/vars"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	logs.Debug("|foundation|init|db|Init")

	var (
		err error
	)
	// 把这个依赖关系放到这里, 这样这里其实是初始化了驱动为Mysql
	if mysql.ErrInvalidConn != nil {
	}

	//
	dbSource := vars.StoreSource.MySQL
	if dbEngine, err = xorm.NewEngine("mysql", dbSource); err != nil {
		logs.Error("Engine Init Err:%v", err)
		panic(err)
	}
	dbEngine.ShowSQL(true)
	dbEngine.ShowExecTime(true)
	dbEngine.SetLogger(xorm.NewSimpleLogger(beego.BeeLogger))
	dbEngine.Logger().SetLevel(core.LOG_INFO)
}

// Engine - 主数据库
func Engine() *xorm.Engine {
	return dbEngine
}
