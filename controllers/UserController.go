package controllers

import (
	"encoding/json"
	"go-rest-api/models"
	"go-rest-api/utils"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user information from JWT token
    user, err := utils.ExtractUserFromToken(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Retrieve user data from the database
    db, err := utils.ConnectDB()
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    var retrievedUser models.User
    err = db.QueryRow("SELECT id, username, role FROM users WHERE id = $1", user.Username).Scan(&retrievedUser.Username, &retrievedUser.Role)
    if err != nil {
        http.Error(w, "Failed to get user data", http.StatusInternalServerError)
        return
    }

    // Return user data
    json.NewEncoder(w).Encode(retrievedUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    // Extract user information from JWT token
    user, err := utils.ExtractUserFromToken(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Parse request body to get updated user data
    var updatedUser models.User
    err = json.NewDecoder(r.Body).Decode(&updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the user is authorized to perform the update (e.g., based on role)
    if user.Role != "admin" && user.Username != updatedUser.Username {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Update user data in the database
    db, err := utils.ConnectDB()
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = db.Exec("UPDATE users SET username = $1, role = $2 WHERE id = $3", updatedUser.Username, updatedUser.Role, updatedUser.ID)
    if err != nil {
        http.Error(w, "Failed to update user data", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User updated successfully"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Extract user information from JWT token
    user, err := utils.ExtractUserFromToken(r)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Parse request body to get user ID to be deleted
    var requestUserID struct {
        UserID int `json:"user_id"`
    }
    err = json.NewDecoder(r.Body).Decode(&requestUserID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Check if the user is authorized to perform the delete operation (e.g., based on role)
    if user.Role != "admin" {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Delete user from the database
    db, err := utils.ConnectDB()
    if err != nil {
        http.Error(w, "Database connection error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM users WHERE id = $1", requestUserID.UserID)
    if err != nil {
        http.Error(w, "Failed to delete user", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User deleted successfully"))
}
