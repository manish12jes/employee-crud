package routers

import (
	"github.com/gorilla/mux"
	"github.com/manish12jes/employee-crud/controllers"
)

func ApiRoutes(apiRouters *mux.Router) *mux.Router {
	employee := apiRouters.PathPrefix("/employees").Subrouter()
	employee.HandleFunc("/", controllers.Employees).Methods("GET")
	employee.HandleFunc("/", controllers.CreateEmployee).Methods("POST")
	employee.HandleFunc("/{id}", controllers.GetEmployee).Methods("GET")
	employee.HandleFunc("/{id}", controllers.UpdateEmployee).Methods("PATCH")
	employee.HandleFunc("/{id}", controllers.DeleteEmployee).Methods("DELETE")

	return apiRouters
}
