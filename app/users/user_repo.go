package users

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbutils "hungon.space/xurl/common/db_utils"
)

type UserRepo struct{}

var userDb *gorm.DB

func init() {
	dns := dbutils.GetDbDns(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSL"),
	)

	userDb, _ = gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func (u *UserRepo) New() *UserRepo {
	return &UserRepo{}
}

// Implement the interface migrate.DbInterface
func (u *UserRepo) AutoMigrate() error {

	err := userDb.AutoMigrate(&User{})

	if err != nil {
		return err
	}

	fmt.Println("User table migrated connected successfully")
	return nil
}

func (u *UserRepo) FindByID(id uint, des *User) error {
	return userDb.First(des, "id = ?", id).Error
}

func (u *UserRepo) FindBySelfID(id string, des *User) error {
	return userDb.First(des, "user_id = ?", id).Error
}

func (u *UserRepo) CreateOne(data *User) error {
	result := userDb.Create(data)
	return result.Error
}

func (u *UserRepo) UpdateOne(data *User) error {
	result := userDb.Save(data)
	return result.Error
}

func (u *UserRepo) DeleteOne(data *User) error {
	result := userDb.Delete(data)
	return result.Error
}

func (u *UserRepo) LastId() (uint, error) {
	var user User
	result := userDb.Last(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (u *UserRepo) FindByEmail(email string, des *User) error {
	return userDb.First(des, "email = ?", email).Error
}

func (u *UserRepo) FindByApiKey(apikey string, des *User) error {
	return userDb.First(des, "api_key = ?", apikey).Error
}
