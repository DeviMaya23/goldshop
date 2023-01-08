package db

import (
	"devimaya/goldshop/buybackstorage/internal/config"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLDB struct {
	db *gorm.DB
}

var (
	GoldShopDB *gorm.DB
)

func init() {
	var err error
	GoldShopDB, err = NewPostgreSQLDB()
	if err != nil {
		fmt.Println(err)
	}
}

func NewPostgreSQLDB() (*gorm.DB, error) {

	// TODO : config db
	conf := config.GetConfig()
	// TODO : config db
	dsn := "host=" + conf.DBHost + " user=" + conf.DBUsername + " password=" + conf.DBPassword + " dbname=" + conf.DBName + " port=" + conf.DBPort + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	postgre, _ := db.DB()
	postgre.SetMaxIdleConns(10)
	postgre.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// Db get db instance of gorm
func (m *PostgreSQLDB) Db() interface{} {
	return m.db
}
