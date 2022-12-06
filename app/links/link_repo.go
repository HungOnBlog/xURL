package links

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbutils "hungon.space/xurl/common/db_utils"
)

type LinkRepo struct {
}

var linkDb *gorm.DB

func init() {
	dns := dbutils.GetDbDns(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"),
	)

	linkDb, _ = gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func (l *LinkRepo) New() *LinkRepo {
	return &LinkRepo{}
}

// Implement the interface migrate.DbInterface
func (l *LinkRepo) AutoMigrate() error {

	err := linkDb.AutoMigrate(&Link{})

	if err != nil {
		return err
	}

	fmt.Println("Link table migrated connected successfully")
	return nil
}

func (l *LinkRepo) FindByID(id uint, des *Link) error {
	return linkDb.First(des, "id = ?", id).Error
}

func (l *LinkRepo) FindBySelfID(id string, des *Link) error {
	return linkDb.First(des, "link_id = ?", id).Error
}

func (l *LinkRepo) CreateOne(data *Link) error {
	result := linkDb.Create(data)
	return result.Error
}

func (l *LinkRepo) UpdateOne(data *Link) error {
	return linkDb.Save(data).Error
}

func (l *LinkRepo) DeleteOne(selfId string) error {
	return linkDb.Delete(&Link{}, "link_id = ?", selfId).Error
}

func (l *LinkRepo) LastId() (uint, error) {
	var lastId Link
	err := linkDb.Last(&lastId).Error
	if err != nil {
		return 0, err
	}

	return lastId.ID, nil
}
