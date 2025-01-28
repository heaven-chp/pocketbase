package main

import (
	"pocketbase/collections"
	"pocketbase/config"
	"pocketbase/flags"
	"pocketbase/superuser"

	"github.com/pocketbase/pocketbase"
)

func init() {
	superuser.Create("admin@admin.com", "admin123")
}

func main() {
	app := pocketbase.New()

	if err := flags.Parse(app); err != nil {
		app.Logger().Error("flags.Parse error", "error", err)
		return
	} else if err := config.Read(flags.Get().ConfigFile); err != nil {
		app.Logger().Error("config.Read error", "file", flags.Get().ConfigFile, "error", err)
		return
	} else {
		app.Logger().Info("config", "", config.Get())
	}

	collections.Upsert(app)

	if err := app.Start(); err != nil {
		app.Logger().Error("app.Start error", "error", err)
	}
}
