package controller

import (
	"github.com/atharvyadav96k/bus-safty/dogzh-server/services"
	"github.com/gorilla/mux"
)

func registerOrgRoutes(orgRouter *mux.Router) {
	orgRouter.HandleFunc("/register", services.Org_Create).Methods("POST")
}
