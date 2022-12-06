package migrate

import "hungon.space/xurl/app/links"

func AppAutoMigration() {
	dbs := []DbInterface{
		&links.LinkRepo{},
	}

	for _, db := range dbs {
		db.AutoMigrate()
	}
}
