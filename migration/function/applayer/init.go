package applayer

import (
	"log"

	"github.com/atharvyadav96k/gcp/app"
)

func Init() *app.App {
	log.Println("Initializing app...")
	
	// Create a new app instance with all dependencies
	appInstance := new(app.App).Init()
	if appInstance == nil {
		log.Panic("Failed to initialize app: returned nil")
	}
	
	log.Println("✓ App initialized")
	log.Println("Initializing environment variables...")
	
	// Initialize environment variables
	appInstance.InitEnvironmentVariables()
	log.Println("✓ Environment variables initialized")
	
	log.Println("Initializing Neon database...")
	
	// Initialize Neon database connection
	if err := appInstance.InitNeon(); err != nil {
		log.Printf("PANIC: Failed to initialize Neon: %v\n", err)
		panic(err)
	}
	
	log.Println("✓ Neon database initialized")
	
	return appInstance
}
