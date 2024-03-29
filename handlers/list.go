package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nabilaps1/API-POSTGRESQL/models"
)

// retorna uma lista de todos

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
