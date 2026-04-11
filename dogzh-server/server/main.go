package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/common/response"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api/health", HealthCheckHandler).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server is running on http://127.0.0.1:8080")
	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the Bus Safety API")
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := response.Response{
		Message: "System is up and running",
		Status:  http.StatusOK,
	}
	response.SendResponse(w)
}
