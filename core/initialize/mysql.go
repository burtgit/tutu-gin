package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"tutu-gin/core/global"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func InitMysql() {
	dbConfig := global.SERVICE_CONFIG.DataBase
	var err error
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Dbname
	global.DB, err = xorm.NewEngine("mysql", dsn)

	tbMapper := names.NewPrefixMapper(names.SnakeMapper{}, dbConfig.TablePrefix)

	global.DB.SetTableMapper(tbMapper)
	global.DB.ShowSQL(true)
	if err != nil {
		panic("数据库连接失败")
	}
	_, err = global.DB.Query("SELECT 1 +1 ")
	if err != nil {
		panic("数据库连接失败")
	}

}
