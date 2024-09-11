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

	// Example command handling (No migration, just setup)
	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "setup":
		// Perform any CouchDB-specific setup or initialization here
		log.Println("Setup complete")
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
