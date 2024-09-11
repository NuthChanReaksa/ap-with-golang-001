package api

import (
	"log"
	"net/http"

	"github.com/NuthChanReaksa/ap-with-golang-001/services/cart"
	"github.com/NuthChanReaksa/ap-with-golang-001/services/order"
	"github.com/NuthChanReaksa/ap-with-golang-001/services/product"
	"github.com/NuthChanReaksa/ap-with-golang-001/services/user"
	"github.com/go-kivik/kivik"
	"github.com/gorilla/mux"
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
