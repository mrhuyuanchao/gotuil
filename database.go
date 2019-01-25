package gotuil

import (
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

// InitDB 初始化数据库配置
func InitDB(config *DBSetting) *xorm.Engine {
	if config == nil {
		return nil
	}
	engine, err := xorm.NewEngine(config.Type, config.DataSource)
	if err != nil {
		LogClient.Error(err.Error())
	}
	return engine
}
