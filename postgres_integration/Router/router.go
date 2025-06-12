package router

import (
	"github.com/gorilla/mux"
	middleware "github.com/isfnev/postgres_integration/Middleware"
)

func Router() (r *mux.Router) {
	r = mux.NewRouter()

	r.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET")
	r.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST")
	r.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET")
	r.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT")
	r.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE")

	return
}
