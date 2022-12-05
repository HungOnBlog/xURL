package migrate

import "hungon.space/xurl/app/links"

func AppAutoMigration() {
	dbs := []DbInterface{
		&links.LinkDb{},
	}

	for _, db := range dbs {
		db.AutoMigrate()
	}
}
