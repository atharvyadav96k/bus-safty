package controller

import "github.com/gorilla/mux"

func RegisterApiRoutes(mainRouter *mux.Router) {
	apiV1 := mainRouter.PathPrefix("/api/v1").Subrouter()

	orgRouter := apiV1.PathPrefix("/org").Subrouter()
	registerOrgRoutes(orgRouter)

	userRouter := apiV1.PathPrefix("/user").Subrouter()
	registerUserRoutes(userRouter)
}
