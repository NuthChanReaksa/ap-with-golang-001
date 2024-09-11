package api

import (
	"log"
	"net/http"

	"github.com/go-kivik/kivik"
	"github.com/gorilla/mux"
	"github.com/sikozonpc/ecom/services/cart"
	"github.com/sikozonpc/ecom/services/order"
	"github.com/sikozonpc/ecom/services/product"
	"github.com/sikozonpc/ecom/services/user"
)

type APIServer struct {
	addr string
	db   *kivik.DB
}

func NewAPIServer(addr string, db *kivik.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize stores
	userStore := user.NewStore(s.db)
	productStore := product.NewStore(s.db)
	orderStore := order.NewStore(s.db)

	// Initialize handlers
	userHandler := user.NewHandler(userStore)
	productHandler := product.NewHandler(productStore, userStore)
	cartHandler := cart.NewHandler(productStore, orderStore, userStore)

	// Register routes
	userHandler.RegisterRoutes(subrouter)
	productHandler.RegisterRoutes(subrouter)
	cartHandler.RegisterRoutes(subrouter)

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Listening on", s.addr)
	err := http.ListenAndServe(s.addr, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return err
}
