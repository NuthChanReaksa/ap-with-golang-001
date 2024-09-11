package db

import (
	"context"                       // Import context package
	_ "github.com/go-kivik/couchdb" // Import the CouchDB driver
	"github.com/go-kivik/kivik"
	"github.com/sikozonpc/ecom/configs"
)

func NewCouchDBStorage() (*kivik.DB, error) {
	cfg := configs.Envs
	client, err := kivik.New("couch", cfg.CouchDBAddress)
	if err != nil {
		return nil, err
	}

	// Create a context for the DB method
	ctx := context.Background()

	// Use the context with the DB method
	db := client.DB(ctx, cfg.CouchDBName)
	if err := db.Err(); err != nil {
		return nil, err
	}

	return db, nil
}
