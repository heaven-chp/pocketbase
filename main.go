package main

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
	"k8s.io/klog/v2"
)

func init() {
	migrations.Register(func(app core.App) error {
		superusers, err := app.FindCachedCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		record := core.NewRecord(superusers)
		record.Set("email", "admin@test.com")
		record.Set("password", "admin123")

		return app.Save(record)
	}, func(app core.App) error {
		return nil
	})
}

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		klog.ErrorS(err, "")
	}
}
