package models

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        int        `json:"id" binding:"-" gorm:"size:36;primary_key"`
	CreatedAt time.Time  `json:"-"  binding:"-" gorm:"index:idx_created_at"`
	UpdatedAt time.Time  `json:"-"  binding:"-"`
	DeletedAt *time.Time `json:"-"  binding:"-" gorm:"index:idx_deleted_at"`
}

var DB *gorm.DB

func InitDB(host string, port int, user string, pwd string, database string) {
	var err error
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pwd, host, port, database,
	)
	DB, err = gorm.Open("mysql", url)
	if err != nil {
		glog.Fatalf("Failed to connect database: %v", err)
		return
	}

	DB.LogMode(true) // FIXME debug

	// defer DB.Close()

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(50)
}

func Seed() {
	DB.AutoMigrate(&User{})
}

func Offset(limit, page int) int {
	if page != -1 && limit != -1 {
		return (page - 1) * limit
	}
	return -1
}
