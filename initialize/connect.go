package initialize

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() {
	connInfo := MySQLConfig
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		connInfo.User,
		connInfo.Password,
		connInfo.Host,
		connInfo.Port,
		connInfo.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(10)              //设置最大连接数
	sqlDb.SetMaxOpenConns(20)              //设置最大的空闲连接数
	data, _ := json.Marshal(sqlDb.Stats()) //获得当前的SQL配置情况
	fmt.Println(string(data))
}
