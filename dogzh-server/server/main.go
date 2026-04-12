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
	os.Getenv("noss")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server is running on http://127.0.0.1:8080")
	log.Fatal(srv.ListenAndServe())
}
