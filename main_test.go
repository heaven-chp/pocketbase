package main

import (
	"os"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMain(t *testing.T) {
	testDataDir := "./test_pb_data_" + uuid.New().String()
	defer os.RemoveAll(testDataDir)

	os.Args = []string{"test", "serve",
		"--config-file=" + "./config/config.json",
		"--dir=" + testDataDir,
		"--http=" + ":8090",
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		main()
	}()
	time.Sleep(1 * time.Second)

	if err := syscall.Kill(os.Getpid(), syscall.SIGTERM); err != nil {
		t.Fatal(err)
	}

	wg.Wait()
}
