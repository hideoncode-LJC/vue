package model

import (
	"axios/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID       int    `gorm:"primaryKey"`
	Title    string `json:"title"`
	Category string `json:"category"'`
	Time     time.Time
	State    string `json:"state"`
}

func NewArticle(id int, title, category, state string) *Article {
	return &Article{
		ID:       id,
		Title:    title,
		Category: category,
		Time:     time.Now(),
		State:    state,
	}
}

func GetDNS(mysqlConfig *config.MysqlConfig) string {
	return mysqlConfig.Username +
		":" +
		mysqlConfig.Password +
		"@(" +
		mysqlConfig.Host +
		":" +
		mysqlConfig.Port +
		")/" +
		mysqlConfig.Database +
		"?charset=" +
		mysqlConfig.Charset +
		"&parseTime=True&loc=Local"
}

func GetDatabaseConnection() *gorm.DB {
	dsn := GetDNS(config.NewMysqlConfig())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 自动迁移
	if err := db.AutoMigrate(&Article{}); err != nil {
		panic(err)
	}
	return db
}
