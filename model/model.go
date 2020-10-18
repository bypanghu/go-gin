package model

import (
	"code/01/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	DB *gorm.DB
)

// gorm.Model 定义
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func InitMysql()(err error)  {
	dsn := config.DataBaseRoot+":"+config.DataBasePassword+"@("+ config.DateBaseIp +":"+config.DataBasePort+")/"+config.DataBaseTable+"?charset=utf8mb4&parseTime=True&loc=Local"
	DB , err = gorm.Open("mysql",dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close()  {
	DB.Close()
}