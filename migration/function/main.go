package migration_function

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atharvyadav96k/bus-safty/migration/applayer"
	"github.com/atharvyadav96k/gcp/common/res"
)

func Migration(w http.ResponseWriter, r *http.Request) {
	app := applayer.Init()
	defer app.Close()

	tables := applayer.GetMigrationTables()
	errs := make([]error, 0)
	migratedTables := make([]string, 0)

	for i, table := range tables {
		tableName := fmt.Sprintf("%T", table)
		log.Printf("[%d/%d] Migrating table: %s\n", i+1, len(tables), tableName)

		if err := app.Neon.GetDB().AutoMigrate(table); err != nil {
			log.Printf("ERROR migrating %s: %v\n", tableName, err)
			errs = append(errs, fmt.Errorf("failed to migrate %s: %w", tableName, err))
			continue
		}

		migratedTables = append(migratedTables, tableName)
		log.Printf("✓ Successfully migrated: %s\n", tableName)
	}

	if len(errs) > 0 {
		log.Printf("Migration completed with %d error(s)\n", len(errs))
		res.BadRequest(w, errs)
		return
	}

	log.Printf("✓ All tables migrated successfully (%d tables)\n", len(migratedTables))
	res.Success(w, "All tables migrated successfully", map[string]interface{}{
		"total_tables": len(migratedTables),
		"tables":       migratedTables,
	})
}
