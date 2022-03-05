package lib

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	// FIXME: 本番環境では無視したいけどやり方わからない
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("not found env file")
	}

	// ローカルでは.envから読み込んだ値, 本番ではcloud runに設定した同名のPlanetScaleへのDSNが読み込めるはず
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

	// err = db.AutoMigrate(&model.User{})
	// if err != nil {
	// 	return nil, err
	// }
	// err = db.AutoMigrate(&model.Todo{})
	// if err != nil {
	// 	return nil, err
	// }
	return db, nil
}
