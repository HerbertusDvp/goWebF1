package main

import (
	"goWebF1/internal/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	userSrv := user.NewService()
	userEnd := user.MakeEndpoints(userSrv)

	router.HandleFunc("/users", userEnd.Create).Methods("POST")
	router.HandleFunc("/users", userEnd.Get).Methods("GET")
	//router.HandleFunc("/users", userEnd.GetAll).Methods("GET")
	router.HandleFunc("/users", userEnd.Update).Methods("PATCH")
	router.HandleFunc("/users", userEnd.Delete).Methods("DELETE")

	srv := &http.Server{
		Handler:      router, //http.TimeoutHandler(router, time.Second*3, "e agotó el tiempo de espera"),
		Addr:         "localhost:8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
