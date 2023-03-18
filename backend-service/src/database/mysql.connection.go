package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("root:123456@tcp(mysql:3306)/cunewbie-search-poc?charset=utf8&parseTime=True")

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
