package db

import (
	"log"

	"github.com/hamza72x/blog-in-your-email/helper"
	"github.com/hamza72x/blog-in-your-email/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
)

func Conn() *gorm.DB {
	var err error
	if conn == nil {
		conn, err = gorm.Open(sqlite.Open(helper.GetDBFile()), &gorm.Config{})
		if err != nil {
			log.Printf("Error opening database connection: %v", err)
			panic(err)
		}
		conn.AutoMigrate(&model.Post{})
	}
	return conn
}

func IsFirstRun(blogTitle string) bool {
	var count int64
	Conn().Table("posts").Where("blog_title = ?", blogTitle).Count(&count)
	return count == 0
}
