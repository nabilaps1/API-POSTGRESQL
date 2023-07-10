package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/nabilaps1/API-POSTGRESQL/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo // parse request

	err := json.NewDecoder(r.Body).Decode(&todo) // faz o payload de uma request json. faz o decode do payload dentro do ponteiro todo
	if err != nil {
		log.Println("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   true, // a melhor pratica seria retornar o status-code da resposta
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID:: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp) // faz o encode do map para um json de resposta da aplicacao
}
