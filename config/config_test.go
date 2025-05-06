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

	if len(config.Get().Collections.Retention) != 1 {
		t.Fatal(len(config.Get().Collections.Retention))
	} else if config.Get().Collections.Retention[0].Name != "sample" {
		t.Fatal(config.Get().Collections.Retention[0].Name)
	} else if config.Get().Collections.Retention[0].Days != 10 {
		t.Fatal(config.Get().Collections.Retention[0].Days)
	} else if config.Get().Collections.Retention[0].Expression != "@daily" {
		t.Fatal(config.Get().Collections.Retention[0].Expression)
	}
}
