package migration_function

import (
	"net/http"

	"github.com/atharvyadav96k/bus-safty/migration/applayer"
)

func Migration(w http.ResponseWriter, r *http.Request) {
	app := applayer.Init()
	tables := applayer.GetMigrationTables()
	app.RegisterModels(tables...)
}
