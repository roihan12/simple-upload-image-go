package database

import (
	"fmt"
	"log"
	"time"

	"github.com/roihan12/task-5-pbi-btpns-roihan-sori-nasution/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_USERNAME string = helpers.GetConfig("POSTGRES_USER")
	DB_PASSWORD string = helpers.GetConfig("POSTGRES_PASSWORD")
	DB_NAME     string = helpers.GetConfig("POSTGRES_DB")
	DB_HOST     string = helpers.GetConfig("POSTGRES_HOST")
	DB_PORT     string = helpers.GetConfig("POSTGRES_PORT")
)

func DBConnection() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		DB_HOST,
		DB_USERNAME,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Return
	return db
}
