package database

import (
	"fmt"
	utils "tech_tubbies/middleware/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	DBErr  error
)

func ConnectPostgres() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		utils.GetEnv("POSTGRES_HOST"), utils.GetEnv("POSTGRES_USERNAME"), utils.GetEnv("POSTGRES_PASSWORD"), utils.GetEnv("DATABASE_NAME"), utils.GetEnv("POSTGRES_PORT"), utils.GetEnv("POSTGRES_SSL"), utils.GetEnv("POSTGRES_TIMEZONE"))
	DBConn, DBErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func CreateTable(table any) error {
	if err := DBConn.AutoMigrate(table); err != nil {
		return err
	}
	return nil
}
