package user

import (
	"context"
	"fmt"

	"github.com/NuthChanReaksa/ap-with-golang-001/types"
	"github.com/go-kivik/kivik"
	"github.com/google/uuid"
)

type Store struct {
	db *kivik.DB
}

func NewStore(db *kivik.DB) *Store {
	return &Store{db: db}
}

func generateUniqueID() string {
	return uuid.NewString()
}

func (s *Store) CreateUser(user types.User) error {
	user.ID = generateUniqueID()
	userDoc := map[string]interface{}{
		"_id":       user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"password":  user.Password,
	}

	_, err := s.db.Put(context.Background(), user.ID, userDoc)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (s *Store) GetUserByID(id string) (*types.User, error) {
	doc := s.db.Get(context.Background(), id)
	if doc.Err != nil { // Access the Err field directly
		return nil, fmt.Errorf("user not found: %v", doc.Err)
	}

	var user types.User
	err := doc.ScanDoc(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Find(context.Background(), map[string]interface{}{
		"selector": map[string]interface{}{
			"email": email,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	var result types.User
	if rows.Next() {
		if err := rows.ScanDoc(&result); err != nil {
			return nil, err
		}
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("user not found: %v", rows.Err())
	}

	return &result, nil
}
