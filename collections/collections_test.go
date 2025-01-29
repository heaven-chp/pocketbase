package collections_test

import (
	"net/http"
	"os"
	"pocketbase/collections"
	"pocketbase/config"
	"testing"

	"github.com/google/uuid"
	"github.com/pocketbase/pocketbase/tests"
)

var testDataDir = "./test_pb_data_" + uuid.New().String()

func getTestApp(t testing.TB) *tests.TestApp {
	testApp, err := tests.NewTestApp(testDataDir)
	if err != nil {
		t.Fatal(err)
	}

	collections.Upsert(testApp)
	collections.Cron(testApp)

	return testApp
}

func TestMain(m *testing.M) {
	setup := func() {
		if err := config.Read("../config/config.json"); err != nil {
			panic(err)
		} else if err := os.Mkdir(testDataDir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	teardown := func() {
		if err := os.RemoveAll(testDataDir); err != nil {
			panic(err)
		}
	}

	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestCron(t *testing.T) {
}

func TestUpsert(t *testing.T) {
	apiScenarios := []tests.ApiScenario{
		{
			Name:            "get sample collection records.",
			Method:          http.MethodGet,
			URL:             "/api/collections/sample/records",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`{"items":[],"page":1,"perPage":30,"totalItems":0,"totalPages":0}`},
			TestAppFactory:  getTestApp,
		},
	}

	for _, apiScenario := range apiScenarios {
		apiScenario.Test(t)
	}
}
