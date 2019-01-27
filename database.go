package goutil

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

// InitDB 初始化数据库配置
func InitDB(config *DBSetting) (*xorm.Engine, error) {
	if config == nil {
		return nil, fmt.Errorf("无配置")
	}
	return xorm.NewEngine(config.Type, config.DataSource)
}
