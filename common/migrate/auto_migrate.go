package migrate

import (
	"hungon.space/xurl/app/links"
	"hungon.space/xurl/app/users"
)

func AppAutoMigration() {
	dbs := []DbInterface{
		&links.LinkRepo{},
		&users.UserRepo{},
	}

	for _, db := range dbs {
		db.AutoMigrate()
	}
}
