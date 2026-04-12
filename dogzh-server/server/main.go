package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/atharvyadav96k/bus-safty/dogzh-server/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	controller.RegisterApiRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
