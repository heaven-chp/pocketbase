package superuser

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
)

func Create(email, password string) {
	migrations.Register(func(app core.App) error {
		if collection, err := app.FindCachedCollectionByNameOrId(core.CollectionNameSuperusers); err != nil {
			return err
		} else {
			record := core.NewRecord(collection)
			record.Set("email", email)
			record.Set("password", password)

			return app.Save(record)
		}
	}, func(app core.App) error {
		if record, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, email); record == nil {
			return nil
		} else {
			return app.Delete(record)
		}
	}, "superuser.Create-"+email)
}
