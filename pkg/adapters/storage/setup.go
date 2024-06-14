package storage

import (
	"fmt"
	"service/config"
	"service/pkg/adapters/storage/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlGormConnection(dbConfig config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()

	migrator.AutoMigrate(&entities.User{})
}
