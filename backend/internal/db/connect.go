package db

import (
	"fmt"

	"github.com/xuying7/kill-leaf/internal/config"
	"github.com/xuying7/kill-leaf/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	host := config.EnvVar("DB_HOST", "localhost")
	port := config.EnvVar("DB_PORT", "3306")
	user := config.EnvVar("DB_USER", "root")
	password := config.EnvVar("DB_PASSWORD", "")
	dbName := config.EnvVar("DB_NAME", "backend")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = dbConn

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	return nil
}
