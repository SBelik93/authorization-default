package models

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var dbGlobal *gorm.DB

func NewDB(env *Environment) (sqlDB *sql.DB, err error) {
	fmt.Println("Start connecting to DB: ", env.ConnectionDbString)
	dbGlobal, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  env.ConnectionDbString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return
	}
	// logs
	dbGlobal = dbGlobal.Debug()
	sqlDB, err = dbGlobal.DB()
	if err != nil {
		return
	}
	sqlDB.SetConnMaxLifetime(time.Duration(1) * time.Hour)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)
	fmt.Println("DB connected: ", env.ConnectionDbString)
	return
}

func GetDB() *gorm.DB {
	return dbGlobal
}
