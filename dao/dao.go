package dao

import (
	"ginBlog/config"
	"ginBlog/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *models.User)
}
type manager struct {
	db *gorm.DB
}

func (m manager) AddUser(user *models.User) {
	m.db.Create(user)
}

var Mgr Manager

func init() {
	dsn := config.DBConfig
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to init db:", err)
	}
	Mgr = &manager{db: db}

	merr := db.AutoMigrate(&models.User{})
	if merr != nil {
		return
	}

}
