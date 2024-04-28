package api

import (
	"gof/internal/models/departments"
	"gof/internal/models/districts"
	"gof/internal/models/ranks"
	"gof/internal/models/staffs"

	"gofr.dev/pkg/gofr"
)

func RegisterCRUD(app *gofr.App) {

	for _, model := range []any{
		&departments.Department{},
		&districts.District{},
		&ranks.Rank{},
		&staffs.Staff{},
	} {
		if err := app.AddRESTHandlers(model); err != nil {
			panic(err)
		}
	}

}
