package config_test

import (
	"pocketbase/config"
	"testing"
)

func TestRead(t *testing.T) {
	if err := config.Read("./config.json"); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	TestRead(t)
}
