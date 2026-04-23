package migration_function

import (
	"net/http"

	"github.com/atharvyadav96k/bus-safty/migration/applayer"
	"github.com/atharvyadav96k/gcp/common/res"
)

func Migration(w http.ResponseWriter, r *http.Request) {
	app := applayer.Init()
	tables := applayer.GetMigrationTables()
	if err := app.RegisterModels(tables...); err != nil {
		res.InternalServerError(w, []error{err})
		return
	}
	res.Success(w, "Successful Completed migration", nil)
}
