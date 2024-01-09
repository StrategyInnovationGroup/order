package initializer

import (
	"fmt"
	"log"
	"order/pkg/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConfig *config.DatabaseConfig

func DBConnection() *gorm.DB {

	dbConfig = LoadDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata", dbConfig.DB_HOST, dbConfig.DB_USER, dbConfig.DB_PSWD, dbConfig.DB_NAME, dbConfig.DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf("Connection to Postgres DB failed !!!! ")
	}

	return db
}

func RunDBMigration() {

	dbConfig = LoadDBConfig()

	fmt.Println("DB config user details := ", dbConfig)

	m, err := migrate.New(
		"file://pkg/db/migrate",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&search_path=public", dbConfig.DB_USER, dbConfig.DB_PSWD, dbConfig.DB_HOST, dbConfig.DB_PORT, dbConfig.DB_NAME),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
