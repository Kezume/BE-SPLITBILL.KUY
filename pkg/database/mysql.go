package database

import (
	"fmt"

	"github.com/Kezume/BE-SPLITBILL.KUY/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "splitbill.kuy")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&model.User{},
		&model.Expense{},
		&model.Friendship{},
		&model.GroupMember{},
		&model.Groups{},
		&model.Transaction{},
		&model.ExpenseSplit{},
	)

	DB = db
}
