package lib

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/do-it-if-i-can/server/graph/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("loading failed var from env: %v", err)
	}

	dsn := os.Getenv("MYSQL_DSN")
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

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
