package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

func ConnectDB() *sql.DB {
	connectionString := fmt.Sprintf("%v:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("error cant connect to mysql : ", err.Error())
	} else {
		log.Println("success connection to database mysql")
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetConnMaxIdleTime(40 * time.Minute)

	return db
}
