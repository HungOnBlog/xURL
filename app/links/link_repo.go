package links

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbutils "hungon.space/xurl/common/db_utils"
)

type LinkDb struct {
}

// Implement the interface migrate.DbInterface
func (l *LinkDb) AutoMigrate() error {
	dns := dbutils.GetDbDns(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"),
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&Link{})
	return nil
}
