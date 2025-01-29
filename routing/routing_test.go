package routing_test

import (
	"os"
	"pocketbase/routing"
	"testing"

	"github.com/google/uuid"
	"github.com/pocketbase/pocketbase/tests"
)

var testDataDir = "./test_pb_data_" + uuid.New().String()

func TestRegister(t *testing.T) {
	if err := os.Mkdir(testDataDir, os.ModePerm); err != nil {
		panic(err)
	}
	defer func() {
		if err := os.RemoveAll(testDataDir); err != nil {
			t.Fatal(err)
		}
	}()

	if testApp, err := tests.NewTestApp(testDataDir); err != nil {
		t.Fatal(err)
	} else {
		defer testApp.Cleanup()

		routing.Register(testApp)
	}
}
