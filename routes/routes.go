package routes

import (
	"net/http"

	"go-rest-api/controllers"
	"go-rest-api/utils"

	"github.com/gorilla/mux"
)
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract user information from JWT token
        user, err := utils.ExtractUserFromToken(r)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Check if the user is an admin
        if user.Role != "admin" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // Proceed to the next middleware or handler
        next.ServeHTTP(w, r)
    })
}

func SetAuthRoutes(router *mux.Router) *mux.Router {
    router.HandleFunc("/register", controllers.Register).Methods("POST")
    router.HandleFunc("/login", controllers.Login).Methods("POST")
    return router
}
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.Use(AuthMiddleware)
    router.HandleFunc("/user", controllers.GetUser).Methods("GET")
    router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
    router.HandleFunc("/user", controllers.DeleteUser).Methods("DELETE")
    return router
}