package migration_function

import (
	"net/http"

	"github.com/atharvyadav96k/bus-safty/migration/applayer"
	"github.com/atharvyadav96k/gcp/common/res"
)

func Migration(w http.ResponseWriter, r *http.Request) {
	app := applayer.Init()
	tables := applayer.GetMigrationTables()
	errs := make([]error, 0)
	for _, table := range tables {
		if err := app.Neon.GetDB().AutoMigrate(table); err != nil {
			errs = append(errs, err)
			return
		}
	}
	res.Success(w, "Successful Completed migration", nil)
}
