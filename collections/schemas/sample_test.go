package schemas

import (
	"encoding/json"
	"testing"

	"github.com/pocketbase/pocketbase/core"
)

func TestSample(t *testing.T) {
	sample := sample{}

	if sample.Type() != core.CollectionTypeBase {
		t.Fatal(sample.Type())
	} else if sample.Name() != "sample" {
		t.Fatal(sample.Name())
	} else if result, err := json.Marshal(sample.Fields()); err != nil {
		t.Fatal(err)
	} else if string(result) != `[{"autogeneratePattern":"","hidden":false,"id":"","max":0,"min":0,"name":"field_01","pattern":"","presentable":false,"primaryKey":false,"required":false,"system":false,"type":"text"},{"hidden":false,"id":"","name":"created","onCreate":true,"onUpdate":false,"presentable":false,"system":false,"type":"autodate"},{"hidden":false,"id":"","name":"updated","onCreate":true,"onUpdate":true,"presentable":false,"system":false,"type":"autodate"}]` {
		t.Fatal(string(result))
	} else if sample.ListRule() != "" {
		t.Fatal(sample.ListRule())
	} else if sample.ViewRule() != "" {
		t.Fatal(sample.ViewRule())
	} else if sample.CreateRule() != "" {
		t.Fatal(sample.CreateRule())
	} else if sample.UpdateRule() != "" {
		t.Fatal(sample.UpdateRule())
	} else if sample.DeleteRule() != "" {
		t.Fatal(sample.DeleteRule())
	}
}
