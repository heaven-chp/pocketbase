package flag

import (
	"os"

	"github.com/pocketbase/pocketbase"
)

type flag struct {
	ConfigFile string
}

var _flag flag

func Parse(app *pocketbase.PocketBase) error {
	app.RootCmd.PersistentFlags().StringVar(
		&_flag.ConfigFile,
		"config-file",
		"",
		"./config.json",
	)

	return app.RootCmd.ParseFlags(os.Args[1:])
}

func Get() flag {
	return _flag
}
