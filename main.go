package main

import (
	"go-rest-api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // Set authentication routes
    router = routes.SetAuthRoutes(router)

    // Set user routes
    router = routes.SetUserRoutes(router)

    http.ListenAndServe(":8080", router)
}
