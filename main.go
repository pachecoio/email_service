package main

import (
	"github.com/gorilla/mux"
	_interface "github.com/pachecoio/email_service/interface"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", _interface.SendEmailHandler).Methods("POST")
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		panic(err)
	}
}
