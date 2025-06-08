package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// addr: "user:pass@tcp(ip:port)/dbname"
func NewMysqlConnection(addr string) (*gorm.DB, error) {
	addr = fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", addr)
	return gorm.Open(mysql.Open(addr))
}
