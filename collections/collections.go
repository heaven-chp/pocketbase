package collections

import (
	"database/sql"
	"pocketbase/collections/schemas"
	"pocketbase/config"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func Cron(app core.App) {
	job := func(jobId, name string, days int, now time.Time) {
		app.Logger().Info("cron start", "jobId", jobId)
		defer app.Logger().Info("cron end", "jobId", jobId)

		fn := func(txApp core.App) error {
			created := now.AddDate(0, 0, -days).UTC().Format("2006-01-02 15:04:05")

			if records, err := txApp.FindRecordsByFilter(name, "created < {:created}", "created", 0, 0, dbx.Params{"created": created}); err != nil {
				return err
			} else {
				app.Logger().Debug("cron in progress", "jobId", jobId, "created", created, "len(records)", len(records))

				for _, record := range records {
					if err := txApp.Delete(record); err != nil {
						return err
					}
				}
			}

			return nil
		}

		if err := app.RunInTransaction(fn); err != nil {
			app.Logger().Error("app.RunInTransaction error", "jobId", jobId, "error", err)
		}
	}

	for _, retention := range config.Get().Collections.Retention {
		jobId := "collections.retention." + retention.Name
		cronExpr := retention.Expression
		run := func() {
			job(jobId, retention.Name, retention.Days, time.Now())
		}

		if err := app.Cron().Add(jobId, cronExpr, run); err != nil {
			app.Logger().Error("app.Cron().Add error", "jobId", jobId, "error", err)
		}
	}
}

func Upsert(app core.App) {
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		for _, schema := range schemas.Get() {
			collection, err := app.FindCollectionByNameOrId(schema.Name())
			if err != nil && err != sql.ErrNoRows {
				return err
			} else if collection == nil {
				collection = core.NewCollection(schema.Type(), schema.Name())
			}

			collection.Fields = schema.Fields()

			collection.ListRule = types.Pointer(schema.ListRule())
			collection.ViewRule = types.Pointer(schema.ViewRule())
			collection.CreateRule = types.Pointer(schema.CreateRule())
			collection.UpdateRule = types.Pointer(schema.UpdateRule())
			collection.DeleteRule = types.Pointer(schema.DeleteRule())

			if err := e.App.Save(collection); err != nil {
				return err
			}
		}

		return e.Next()
	})
}
