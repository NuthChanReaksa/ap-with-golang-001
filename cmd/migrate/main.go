package main

import (
	"log"
	"os"

	"github.com/NuthChanReaksa/ap-with-golang-001/db"
	"github.com/go-kivik/kivik"
)

func main() {
	// Initialize CouchDB storage
	couchDB, err := db.NewCouchDBStorage()
	if err != nil {
		log.Fatal(err)
	}

	// Check CouchDB connection
	if err := initStorage(couchDB); err != nil {
		log.Fatal(err)
	}

	// Example command handling with migration support
	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "setup":
		// Perform any CouchDB-specific setup or initialization here
		log.Println("Setup complete")
	case "up":
		// Implement migration logic here (e.g., creating databases, inserting initial data)
		log.Println("Performing migration up")
		// Add migration logic for 'up'
		performMigrationUp(couchDB)
	default:
		log.Fatalf("Unknown command: %s", cmd)
	}
}

// initStorage checks if the connection to CouchDB is successful
func initStorage(db *kivik.DB) error {
	// For CouchDB, you might want to check if the database exists or is reachable
	err := db.Err()
	if err != nil {
		return err
	}

	log.Println("DB: Successfully connected!")
	return nil
}

// performMigrationUp defines the logic to migrate the database up
func performMigrationUp(db *kivik.DB) {
	// Example migration logic: Creating a new document or database
	log.Println("Running 'up' migration...")

	log.Println("Migration up completed successfully!")
}
