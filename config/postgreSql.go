package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	appConfig := ViperConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=require TimeZone=Asia/Jakarta",
		appConfig.HostDB,
		appConfig.UsernameDB,
		appConfig.PasswordDB,
		appConfig.DatabaseName,
		appConfig.PortDB,
	)

	// CONNECT TO POSTGRESQL
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Default().Println("Database Connected Succesfully")

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(15)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}
