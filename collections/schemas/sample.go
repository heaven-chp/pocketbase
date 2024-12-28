package schemas

import "github.com/pocketbase/pocketbase/core"

type sample struct {
}

func (this *sample) Type() string {
	return core.CollectionTypeBase
}

func (this *sample) Name() string {
	return "sample"
}

func (this *sample) Fields() core.FieldsList {
	return core.FieldsList{
		&core.TextField{Name: "field_01"},

		&core.AutodateField{
			Name:     "created",
			OnCreate: true,
		},
		&core.AutodateField{
			Name:     "updated",
			OnCreate: true,
			OnUpdate: true,
		},
	}
}

func (this *sample) ListRule() string {
	return ""
}

func (this *sample) ViewRule() string {
	return ""
}

func (this *sample) CreateRule() string {
	return ""
}

func (this *sample) UpdateRule() string {
	return ""
}

func (this *sample) DeleteRule() string {
	return ""
}
