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

	log.Println("Starting database migration...")

	// Verify database connection
	if app.Neon == nil {
		log.Println("ERROR: Neon service is nil")
		res.BadRequest(w, []error{fmt.Errorf("database service not initialized")})
		return
	}

	db := app.Neon.GetDB()
	if db == nil {
		log.Println("ERROR: Database instance is nil")
		res.BadRequest(w, []error{fmt.Errorf("database connection failed")})
		return
	}

	// Test database connection
	if err := db.Exec("SELECT 1").Error; err != nil {
		log.Printf("ERROR: Database connection test failed: %v\n", err)
		res.BadRequest(w, []error{fmt.Errorf("failed to connect to database: %w", err)})
		return
	}
	log.Println("✓ Database connection established")

	tables := applayer.GetMigrationTables()
	errs := make([]error, 0)
	migratedTables := make([]string, 0)
	failedTables := make([]string, 0)

	// Migrate tables sequentially
	for i, table := range tables {
		tableName := fmt.Sprintf("%T", table)
		log.Printf("[%d/%d] Migrating table: %s\n", i+1, len(tables), tableName)

		// AutoMigrate with explicit error checking
		result := db.AutoMigrate(table)
		if result.Error != nil {
			log.Printf("ERROR migrating %s: %v\n", tableName, result.Error)
			errs = append(errs, fmt.Errorf("failed to migrate %s: %w", tableName, result.Error))
			failedTables = append(failedTables, tableName)
			continue
		}

		// Verify table was created by checking if it exists
		if !db.Migrator().HasTable(table) {
			errMsg := fmt.Sprintf("table %s was not created", tableName)
			log.Printf("ERROR: %s\n", errMsg)
			errs = append(errs, fmt.Errorf(errMsg))
			failedTables = append(failedTables, tableName)
			continue
		}

		migratedTables = append(migratedTables, tableName)
		log.Printf("✓ Successfully migrated: %s\n", tableName)
	}

	// Check if any errors occurred during migration
	if len(errs) > 0 {
		log.Printf("Migration completed with %d error(s), %d failed\n", len(errs), len(failedTables))
		res.BadRequest(w, errs)
		return
	}

	log.Printf("✓ All tables migrated successfully (%d tables)\n", len(migratedTables))
	res.Success(w, "All tables migrated successfully", map[string]interface{}{
		"total_tables":  len(migratedTables),
		"migrated":      migratedTables,
		"failed_tables": failedTables,
	})
}
