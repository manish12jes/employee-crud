package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/manish12jes/employee-crud/routers"
)

func main() {
	routes := routers.Routes()
	// routes = routers.WebRoutes()   // if other routes are present

	server := &http.Server{
		Handler:      routes,
		Addr:         "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Starting server at :", server.Addr)
	log.Fatal(server.ListenAndServe())
}
