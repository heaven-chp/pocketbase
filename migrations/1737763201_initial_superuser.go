package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
)

func init() {
	const email = "admin@admin.com"
	const password = "admin123"

	migrations.Register(func(app core.App) error {
		if record, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, email); record != nil {
			return nil
		} else if collection, err := app.FindCachedCollectionByNameOrId(core.CollectionNameSuperusers); err != nil {
			return err
		} else {
			record := core.NewRecord(collection)
			record.Set("email", email)
			record.Set("password", password)

			return app.Save(record)
		}
	}, nil)
}
