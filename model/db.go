package model

import (
	"babyblog/utils"
	"fmt"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	conn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		fmt.Printf("连接数据库错误。请检查连接参数", err)
	}
	fmt.Printf("db object info: %p, %T\n", db, db)
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	//sqlDb, err := db.DB()
	//// db.SingularTable(true)
	//// // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	//sqlDb.SetMaxIdleConns(10)
	//
	//// // SetMaxOpenConns 设置打开数据库连接的最大数量。
	//sqlDb.SetMaxOpenConns(100)
	//
	//// // SetConnMaxLifetime 设置了连接可复用的最大时间。
	//sqlDb.SetConnMaxLifetime(time.Hour)
}
