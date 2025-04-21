package handlers

import (
	"enconding/json"
	"net/http"

	
	"github.com/mabittar/go_api/internal/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []models.UserModel{
		{ID: 1, Name: "John Doe", Email: "john@test.com"},
	}

	w.Headers().Set("Content-Type": "application/json")


	json.NewEncoder(w).Encode(users)
}

