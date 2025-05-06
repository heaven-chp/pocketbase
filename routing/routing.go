package routing

import (
	"pocketbase/routing/sample"

	"github.com/pocketbase/pocketbase/core"
)

func Register(app core.App) {
	sample.Register(app)
}
