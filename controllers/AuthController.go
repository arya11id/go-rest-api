package controllers

import (
	"encoding/json"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    user.Password = string(hashedPassword)

    // Save user to database
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Retrieve user from database by username
    // Compare hashed password from the database with the provided password
    // if err := bcrypt.CompareHashAndPassword([]byte(hashedPasswordFromDB), []byte(user.Password)); err != nil {
    //     http.Error(w, "Invalid username or password", http.StatusUnauthorized)
    //     return
    // }

    // Generate JWT token
    token, err := utils.GenerateToken(user.Username, user.Role)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Return token to client
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
