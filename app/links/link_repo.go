package links

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbutils "hungon.space/xurl/common/db_utils"
)

type LinkDb struct {
}

var linkRepo *gorm.DB

func init() {
	dns := dbutils.GetDbDns(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"),
	)

	linkRepo, _ = gorm.Open(postgres.Open(dns), &gorm.Config{})

}

// Implement the interface migrate.DbInterface
func (l *LinkDb) AutoMigrate() error {

	err := linkRepo.AutoMigrate(&Link{})

	if err != nil {
		return err
	}

	fmt.Println("Link table migrated connected successfully")
	return nil
}

func LastLinkId() (uint, error) {
	var link Link
	result := linkRepo.Last(&link)
	if link == (Link{}) {
		return 0, nil
	}

	return link.ID, result.Error
}

func CreateLink(link *Link) error {
	result := linkRepo.Create(link)
	return result.Error
}
