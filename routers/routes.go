package routers

import (
	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	router = ApiRoutes(router) // Register API routes
	// router = WebRoutes(router)  // Register web routes
	return router
}
