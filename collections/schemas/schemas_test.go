package schemas_test

import (
	"pocketbase/collections/schemas"
	"testing"
)

func TestGet(t *testing.T) {
	if len(schemas.Get()) != 1 {
		t.Fatal(len(schemas.Get()))
	}
}
