package company

import (
	"github.com/gorilla/mux"
)
func Route(r *mux.Router){
	c:=r.PathPrefix("/company").Subrouter().StrictSlash(true)
	c.HandleFunc("/",index).Methods("GET")
	c.HandleFunc("/register",register).Methods("POST")
	c.HandleFunc("/invite",invite).Methods("POST")
}