package main

import (
	"diary-app/handler"
	"diary-app/middleware"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/diary", handler.RegisterDiary).Methods("POST")
	r.HandleFunc("/login", handler.Login).Methods("POST")

	// Protected routes (username from JWT)
	protected := r.PathPrefix("/diary").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// These are now protected
	protected.HandleFunc("/", handler.GetDiary).Methods("GET")
	protected.HandleFunc("/", handler.DeleteDiary).Methods("DELETE")
	protected.HandleFunc("/entry", handler.AddEntry).Methods("POST")
	protected.HandleFunc("/entry/{id}", handler.UpdateEntry).Methods("PUT")
	protected.HandleFunc("/entry/{id}", handler.DeleteEntry).Methods("DELETE")
	protected.HandleFunc("/lock", handler.LockDiary).Methods("PUT")
	protected.HandleFunc("/unlock", handler.UnlockDiary).Methods("PUT")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
