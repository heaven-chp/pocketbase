package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/migrations"
)

func init() {
	migrations.Register(func(app core.App) error {
		settings := app.Settings()

		settings.Meta.AppName = "backend"
		settings.Logs.MaxDays = 7
		settings.Logs.MinLevel = -4

		return app.Save(settings)
	}, nil)
}
