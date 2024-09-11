package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user ID is not a valid CouchDB ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user/abc", nil) // "abc" is not a valid UUID or CouchDB-style ID
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/user/{userID}", handler.handleGetUser).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should handle get user by valid ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/user/12345", nil) // Assuming "12345" is a valid CouchDB ID for this test
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/user/{userID}", handler.handleGetUser).Methods(http.MethodGet)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) UpdateUser(u types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return &types.User{}, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}

// Modified to use string for CouchDB IDs
func (m *mockUserStore) GetUserByID(id string) (*types.User, error) {
	return &types.User{
		ID:        id, // Mock a user with the provided CouchDB ID
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
	}, nil
}
