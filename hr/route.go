package hr

import "github.com/gorilla/mux"

func Route(hr *mux.Router) {
	h := hr.PathPrefix("/hr").Subrouter().StrictSlash(true)
	h.HandleFunc("/join",join).Methods("POST")
}
