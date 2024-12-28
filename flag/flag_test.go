package flag_test

import (
	"os"
	"pocketbase/flag"
	"testing"

	"github.com/pocketbase/pocketbase"
)

func TestParse(t *testing.T) {
	os.Args = []string{
		"test",
		"--config-file=./config/config.json",
	}

	app := pocketbase.New()

	if err := flag.Parse(app); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	TestParse(t)

	if flag.Get().ConfigFile != "./config/config.json" {
		t.Fatal(flag.Get().ConfigFile)
	}
}
