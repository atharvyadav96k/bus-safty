package controller

import (
	"github.com/atharvyadav96k/bus-safty/dogzh-server/services"
	"github.com/gorilla/mux"
)

func registerUserRoutes(orgRouter *mux.Router) {
	orgRouter.HandleFunc("/register", services.User_Register).Methods("POST")
}
