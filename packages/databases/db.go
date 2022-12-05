package databases

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hungon.space/xurl/app/models"
)

var db *gorm.DB

// Return a database connection string from environment variables
func dbDnsString() string {
	return "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
}

func init() {
	var err error
	db, err = gorm.Open(postgres.Open(dbDnsString()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")
	db.AutoMigrate(&models.Link{})
	fmt.Println("Database Migrated")
}

func GetDB() *gorm.DB {
	return db
}
