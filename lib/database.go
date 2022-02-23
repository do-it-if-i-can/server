package lib

import (
	"database/sql"
	"os"

	"github.com/do-it-if-i-can/server/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("MYSQL_DSN")
	sqlDB, _ := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
