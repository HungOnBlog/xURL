package migrate

type DbInterface interface {
	AutoMigrate() error
}
