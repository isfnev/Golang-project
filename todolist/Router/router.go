package router

import (
	service "Todo/Services"

	"github.com/gorilla/mux"
)

func Router() (R *mux.Router) {
	R = mux.NewRouter()

	R.HandleFunc("/todo_", service.CreateNote).Methods("POST")
	R.HandleFunc("/todo", service.UpdateNote).Methods("PUT")
	R.HandleFunc("/todo", service.GetAllNotes).Methods("GET")
	R.HandleFunc("/todo", service.DeleteNote).Methods("DELETE")

	return
}
