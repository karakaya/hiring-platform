package company

import "github.com/gorilla/mux"
func Route(r *mux.Router){
	r.StrictSlash(true)
	c:=r.PathPrefix("/company").Subrouter()
	c.HandleFunc("/",index).Methods("GET")
	c.HandleFunc("/register",register).Methods("POST")
	c.HandleFunc("/invite",invite).Methods("POST")
}