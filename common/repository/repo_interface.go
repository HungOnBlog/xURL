package repository

type RepoInterface interface {
	FindByID(id uint, des *interface{}) error
	FindBySelfID(id string, des *interface{}) error
	CreateOne(data interface{}) error
	UpdateOne(data interface{}) error
	DeleteOne(selfId string) error
	LastId() (uint, error)
}
