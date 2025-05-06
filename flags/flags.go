package flags

import (
	"os"

	"github.com/pocketbase/pocketbase"
)

type flags struct {
	ConfigFile string
}

var _flags flags

func Parse(app *pocketbase.PocketBase) error {
	app.RootCmd.PersistentFlags().StringVar(
		&_flags.ConfigFile,
		"config-file",
		"",
		"./config.json",
	)

	return app.RootCmd.ParseFlags(os.Args[1:])
}

func Get() flags {
	return _flags
}
