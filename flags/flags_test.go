package flags_test

import (
	"os"
	"pocketbase/flags"
	"testing"

	"github.com/pocketbase/pocketbase"
)

func TestParse(t *testing.T) {
	os.Args = []string{
		"test",
		"--config-file=./config/config.json",
	}

	app := pocketbase.New()

	if err := flags.Parse(app); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	TestParse(t)

	if flags.Get().ConfigFile != "./config/config.json" {
		t.Fatal(flags.Get().ConfigFile)
	}
}
