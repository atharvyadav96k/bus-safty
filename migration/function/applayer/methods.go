package applayer

import database_models "github.com/atharvyadav96k/bus-safty-app/database/models"

func GetMigrationTables() []interface{} {
	return []interface{}{
		database_models.Org{},
		database_models.WhiteListedEmail{},

		database_models.RFID{},

		database_models.RootUser{},
		database_models.Scanner{},
		database_models.Vehicle{},

		database_models.User{},
	}
}
