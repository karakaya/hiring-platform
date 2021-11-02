package employee

import "github.com/gorilla/mux"

func Route(r *mux.Router){
	e:=r.PathPrefix("/employee").Subrouter().StrictSlash(true)
	e.HandleFunc("/apply",apply).Methods("POST")
}
