package sample_test

import (
	"net/http"
	"os"
	"pocketbase/config"
	"pocketbase/routing/sample"
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

	sample.Register(testApp)

	return testApp
}

func TestMain(m *testing.M) {
	setup := func() {
		if err := config.Read("../../config/config.json"); err != nil {
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

func TestRegister(t *testing.T) {
	apiScenarios := []tests.ApiScenario{
		{
			Name:            "sample api test.",
			Method:          http.MethodGet,
			URL:             "/sample/id-01",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"id-01"},
			TestAppFactory:  getTestApp,
		},
	}

	for _, apiScenario := range apiScenarios {
		apiScenario.Test(t)
	}
}
