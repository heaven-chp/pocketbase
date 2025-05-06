package sample

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func Register(app core.App) {
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/sample/{id}", func(e *core.RequestEvent) error {
			id := e.Request.PathValue("id")

			return e.String(http.StatusOK, id)
		})

		return se.Next()
	})
}
