package schemas

import "github.com/pocketbase/pocketbase/core"

type schemaInterface interface {
	Type() string
	Name() string

	Fields() core.FieldsList

	ListRule() string
	ViewRule() string
	CreateRule() string
	UpdateRule() string
	DeleteRule() string
}

func Get() []schemaInterface {
	return []schemaInterface{
		&sample{},
	}
}
