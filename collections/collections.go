package collections

import (
	"database/sql"
	"pocketbase/collections/schemas"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func Upsert(app core.App) {
	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		for _, schema := range schemas.Get() {
			collection, err := app.FindCollectionByNameOrId(schema.Name())
			if err != nil && err != sql.ErrNoRows {
				return err
			} else if collection == nil {
				collection = core.NewCollection(schema.Type(), schema.Name())
			}

			collection.Fields = schema.Fields()

			collection.ListRule = types.Pointer(schema.ListRule())
			collection.ViewRule = types.Pointer(schema.ViewRule())
			collection.CreateRule = types.Pointer(schema.CreateRule())
			collection.UpdateRule = types.Pointer(schema.UpdateRule())
			collection.DeleteRule = types.Pointer(schema.DeleteRule())

			if err := e.App.Save(collection); err != nil {
				return err
			}
		}

		return e.Next()
	})
}
