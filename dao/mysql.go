package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

// 初始化数据库
func InitMySQL() (err error) {
	dsn := "root:root1234@(127.0.0.1:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = DB.DB().Ping()
	return err
}

func Close() {
	DB.Close()
}
