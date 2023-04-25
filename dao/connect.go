package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectMySQL() (err error) {
	connInfo := MySQLConfig
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		connInfo.User,
		connInfo.Password,
		connInfo.Host,
		connInfo.Port,
		connInfo.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(DB.Error)

	return err
}
